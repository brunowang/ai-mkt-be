package service

import (
	v1 "ai-mkt-be/api/filmclip/v1"
	"ai-mkt-be/internal/agents/llm"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	"os"
	"path/filepath"

	"ai-mkt-be/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/uuid"
)

// FilmclipService is a filmclip service.
type FilmclipService struct {
	v1.UnimplementedFilmclipServer
	lg *log.Helper

	agentGraph *biz.AgentGraph
	uc         *biz.GreeterUsecase
}

// NewFilmclipService new a filmclip service.
func NewFilmclipService(logger log.Logger, agentGraph *biz.AgentGraph, uc *biz.GreeterUsecase) *FilmclipService {
	return &FilmclipService{lg: log.NewHelper(logger), agentGraph: agentGraph, uc: uc}
}

func (s *FilmclipService) UploadImage(ctx context.Context, in *v1.UploadImageRequest) (*v1.UploadImageReply, error) {
	if in.Base64 == "" {
		return nil, errors.New(400, "IMAGE_EMPTY", "image data cannot be empty")
	}

	// 解码base64数据
	data, err := base64.StdEncoding.DecodeString(in.Base64)
	if err != nil {
		return nil, errors.New(400, "BASE64_DECODE_ERROR", err.Error())
	}

	// 确定图片格式
	contentType := http.DetectContentType(data)
	var ext string
	switch contentType {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	default:
		return nil, errors.New(400, "UNSUPPORTED_IMAGE_TYPE", "unsupported image type: "+contentType)
	}

	// 创建图片存储目录
	imageDir := "./uploads/images"
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		return nil, errors.New(500, "STORAGE_ERROR", "failed to create storage directory: "+err.Error())
	}

	// 生成唯一文件名
	fileName := fmt.Sprintf("%s_%s%s", in.Name, uuid.New().String(), ext)
	filePath := filepath.Join(imageDir, fileName)

	// 写入文件
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return nil, errors.New(500, "FILE_WRITE_ERROR", "failed to write file: "+err.Error())
	}

	// 生成file://协议的URL
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, errors.New(500, "PATH_ERROR", "failed to get absolute path: "+err.Error())
	}
	url := fmt.Sprintf("file://%s", absPath)

	// 记录日志
	s.lg.WithContext(ctx).Infof("Image uploaded: %s, saved to: %s", in.Name, filePath)

	return &v1.UploadImageReply{Url: url}, nil
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
				llm.NewTextContent("参考这两张图片，第一张是服装图，第二张是模特图。\n" + req.Prompt),
				llm.NewImageContent(req.ClothingImage),
				llm.NewImageContent(req.ModelImage),
			},
		},
	}
	ans, err := ag.Execute(ctx, msgs...)
	if err != nil {
		return nil, err
	}
	return &v1.GenClipScriptReply{Content: ans.Content}, nil
}
