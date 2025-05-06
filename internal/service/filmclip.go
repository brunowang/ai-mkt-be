package service

import (
	v1 "ai-mkt-be/api/filmclip/v1"
	"ai-mkt-be/internal/agents/llm"
	"ai-mkt-be/internal/biz"
	"ai-mkt-be/internal/lib"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/brunowang/gframe/gfs3"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"os"
)

// FilmclipService is a filmclip service.
type FilmclipService struct {
	v1.UnimplementedFilmclipServer
	lg    *log.Helper
	s3mgr *gfs3.S3Mgr

	agentGraph *biz.AgentGraph
	videoGen   *biz.VideoGen
	planUC     *biz.PlanUsecase
}

// NewFilmclipService new a filmclip service.
func NewFilmclipService(logger log.Logger, s3mgr *gfs3.S3Mgr, agentGraph *biz.AgentGraph,
	videoGen *biz.VideoGen, planUC *biz.PlanUsecase) *FilmclipService {
	return &FilmclipService{
		lg:         log.NewHelper(logger),
		s3mgr:      s3mgr,
		agentGraph: agentGraph,
		videoGen:   videoGen,
		planUC:     planUC,
	}
}

func (s *FilmclipService) CreatePlan(ctx context.Context, req *v1.CreatePlanRequest) (*v1.CreatePlanReply, error) {
	planID := lib.GenUniqueID()
	if err := s.planUC.CreatePlan(ctx, &biz.Plan{
		PlanID:   planID,
		PlanName: req.Name,
		Step:     1,
	}); err != nil {
		return nil, errors.New(500, "CREATE_PLAN_ERROR", err.Error())
	}
	return &v1.CreatePlanReply{
		PlanId: planID,
	}, nil
}

func (s *FilmclipService) ListPlan(ctx context.Context, req *v1.ListPlanRequest) (*v1.ListPlanReply, error) {
	list, err := s.planUC.ListPlan(ctx, req.UserId)
	if err != nil {
		return nil, errors.New(500, "LIST_PLAN_ERROR", err.Error())
	}
	ret := new(v1.ListPlanReply)
	for _, plan := range list {
		ret.List = append(ret.List, &v1.PlanInfo{
			PlanId: plan.PlanID,
			Name:   plan.PlanName,
			Step:   int32(plan.Step),
		})
	}
	return ret, nil
}

func (s *FilmclipService) UploadImage(ctx context.Context, req *v1.UploadImageRequest) (*v1.UploadImageReply, error) {
	data, err := base64.StdEncoding.DecodeString(req.Base64)
	if err != nil {
		return nil, errors.New(400, "BASE64_DECODE_ERROR", err.Error())
	}

	s3bucket := os.Getenv("S3_BUCKET")
	s3dir := os.Getenv("S3_IMAGE_DIR")
	s3key := fmt.Sprintf("%s/%s_%s", s3dir, lib.GenUniqueID(), req.Name)
	s3url, err := s.s3mgr.Upload(ctx, s3bucket, s3key, bytes.NewReader(data))
	if err != nil {
		return nil, errors.New(500, "UPLOAD_S3_ERROR", err.Error())
	}

	s.lg.WithContext(ctx).Infof("Image uploaded: %s, saved to: %s", req.Name, s3url)

	plan, err := s.planUC.QueryPlan(ctx, req.PlanId)
	if err != nil {
		return nil, errors.New(500, "QUERY_PLAN_ERROR", err.Error())
	}
	if len(plan.Images) == 0 {
		plan.Images = make(map[string]string)
	}
	plan.Images[req.Type.String()] = s3url

	if err := s.planUC.UpdatePlan(ctx, req.PlanId,
		&biz.Plan{
			Images: plan.Images,
		},
	); err != nil {
		return nil, errors.New(500, "UPDATE_PLAN_ERROR", err.Error())
	}

	return &v1.UploadImageReply{Url: s3url}, nil
}

func (s *FilmclipService) GenClipScript(ctx context.Context, req *v1.GenClipScriptRequest) (*v1.GenClipScriptReply, error) {
	ag, err := s.agentGraph.GetAgent(v1.Intent_GenClipScript)
	if err != nil {
		return nil, err
	}
	msgs := []llm.ReqMessage{
		{
			Role: llm.RoleUser,
			Content: llm.MultiContent{
				llm.NewTextContent("参考这张角色图片。主要关注人物形象和服装穿搭\n" + req.Prompt),
				llm.NewImageContent(req.ActorImage),
			},
		},
	}
	ans, err := ag.Execute(ctx, msgs...)
	if err != nil {
		return nil, err
	}
	if err := s.planUC.UpdatePlan(ctx, req.PlanId,
		&biz.Plan{
			ClipScript: ans.Content,
		},
	); err != nil {
		return nil, errors.New(500, "UPDATE_PLAN_ERROR", err.Error())
	}
	scenes, err := parseSceneScript(ans.Content)
	if err != nil {
		return nil, err
	}
	return &v1.GenClipScriptReply{Scenes: scenes}, nil
}

func parseSceneScript(content string) ([]*v1.SceneScript, error) {
	js, err := llm.ExtractJSONFromText(content)
	if err != nil {
		return nil, err
	}
	arr := make([]map[string]any, 0)
	if err := json.Unmarshal([]byte(js), &arr); err != nil {
		return nil, err
	}
	scenes := make([]*v1.SceneScript, 0, len(arr))
	for _, ele := range arr {
		var scene v1.SceneScript
		sequence, _ := ele["镜头编号"].(float64)
		scene.Sequence = cast.ToString(sequence)
		scene.Description, _ = ele["场景描述"].(string)
		scene.Actions, _ = ele["拍摄动作"].(string)
		scene.ShotType, _ = ele["镜头类型"].(string)
		scenes = append(scenes, &scene)
	}
	return scenes, nil
}
