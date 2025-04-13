package aigc

import (
	v1 "ai-mkt-be/api/filmclip/v1"
	"ai-mkt-be/internal/lib"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/brunowang/gframe/gfs3"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	KlingStatusQueueing   = "submitted"
	KlingStatusProcessing = "processing"
	KlingStatusSuccess    = "succeed"
	KlingStatusFail       = "failed"
)

type KlingSDK struct {
	lg        *log.Helper
	httpc     *resty.Client
	s3mgr     *gfs3.S3Mgr
	endpoint  string
	apiKey    string
	apiSecret string
}

func NewKlingSDK(logger log.Logger, s3mgr *gfs3.S3Mgr) *KlingSDK {
	return &KlingSDK{
		lg:        log.NewHelper(logger),
		httpc:     resty.New().SetTimeout(30 * time.Second),
		s3mgr:     s3mgr,
		endpoint:  "https://api.klingai.com",
		apiKey:    os.Getenv("KLING_API_KEY"),
		apiSecret: os.Getenv("KLING_API_SECRET"),
	}
}

func (s *KlingSDK) InvokeVideoGeneration(ctx context.Context, params InvokeVideoGenerationParams) (res InvokeVideoGenerationResult, traceID string, err error) {
	nowt := time.Now().UTC()
	defer func() {
		res.TaskID = "multi-" + res.TaskID
		var taskID = string(res.TaskID)
		s.lg.WithContext(ctx).Infof("KlingSDK.InvokeVideoGeneration costTime: %v, taskID: %s, thirdTaskID: %s, prompt: %s",
			time.Since(nowt), params.TaskID, taskID, params.Prompt)
		if err != nil {
			s.lg.WithContext(ctx).Errorf("KlingSDK.InvokeVideoGeneration err: %v, taskID: %s", err, params.TaskID)
		}
	}()

	url := fmt.Sprintf("%s/%s", s.endpoint, "v1/videos/image2video")
	payload := map[string]any{
		"model_name": "kling-v1-6",
		"mode":       "std",
		"duration":   "5",
		"image":      params.FirstFrameImage,
		"prompt":     params.Prompt,
		"cfg_scale":  0.5,
	}
	if params.NegativePrompt != "" {
		payload["negative_prompt"] = params.NegativePrompt
	}
	if params.CfgScale != 0 {
		payload["cfg_scale"] = params.CfgScale
	}
	if params.AspectRatio != "" {
		payload["aspect_ratio"] = params.AspectRatio
	}
	if len(params.EffectImageList) > 0 {
		url = fmt.Sprintf("%s/%s", s.endpoint, "v1/videos/multi-image2video")
		delete(payload, "image")
		payload["image_list"] = append([]string{params.FirstFrameImage}, params.EffectImageList...)
	}
	token, err := s.encodeJWTToken(s.apiKey, s.apiSecret)
	if err != nil {
		return res, "", err
	}
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
		"Content-Type":  "application/json",
	}
	var result invokeKlingResult
	s.lg.WithContext(ctx).Infof("KlingSDK invoke create begin, url: %s, payload: %+v, task_id: %s", url, payload, params.TaskID)
	rsp, err := s.httpc.R().SetHeaders(headers).SetBody(payload).Post(url)
	if err != nil {
		return result.Data, "", err
	} else if err := json.Unmarshal(rsp.Body(), &result); err != nil {
		return result.Data, "", err
	}
	traceID = result.RequestID
	if result.Code != 0 {
		if result.Code == 1303 { // 触发限流
			s.lg.WithContext(ctx).Infof("KlingSDK.InvokeVideoGeneration: %s", result.Message)
			return result.Data, traceID, errors.InternalServer(v1.ErrorReason_REQUEST_TOO_FREQUENT.String(), "request too frequent")
		}
		return result.Data, traceID, NewErrReason("KlingSDK.InvokeVideoGeneration failed",
			result.Code, result.Message).WithTrace(traceID)
	} else if result.Data.TaskID == "" {
		return result.Data, traceID, NewErrReason("KlingSDK.InvokeVideoGeneration got empty taskID",
			result.Code, result.Message).WithTrace(traceID)
	}
	return result.Data, traceID, nil
}

func (s *KlingSDK) QueryVideoGeneration(ctx context.Context, taskID string) (res *QueryVideoGenerationResult, traceID string, err error) {
	nowt := time.Now().UTC()
	defer func() {
		s.lg.WithContext(ctx).Infof("KlingSDK.QueryVideoGeneration costTime: %v, thirdTaskID: %s", time.Since(nowt), taskID)
		if err != nil {
			s.lg.WithContext(ctx).Errorf("KlingSDK.QueryVideoGeneration err: %v, taskID: %s", err, taskID)
		}
	}()

	url := fmt.Sprintf("%s/%s/%s", s.endpoint, "v1/videos/image2video", taskID)
	if strings.HasPrefix(taskID, "multi-") {
		taskID = strings.TrimPrefix(taskID, "multi-")
		url = fmt.Sprintf("%s/%s/%s", s.endpoint, "v1/videos/multi-image2video", taskID)
	}
	var result fetchKlingResult
	token, err := s.encodeJWTToken(s.apiKey, s.apiSecret)
	if err != nil {
		return res, "", err
	}
	if _, err := s.httpc.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetResult(&result).Get(url); err != nil {
		return nil, "", err
	}
	traceID = result.RequestID
	if result.Code != 0 {
		return nil, traceID, NewErrReason("QueryVideoGeneration failed", result.Code,
			fmt.Sprintf("%s, task_id: %s", result.Message, taskID)).WithTrace(traceID)
	} else if result.Data.TaskStatusMsg == "Failure to pass the risk control system" {
		return nil, traceID, NewErrReason("QueryVideoGeneration failed", 1301,
			fmt.Sprintf("%s, task_id: %s", result.Data.TaskStatusMsg, taskID)).WithTrace(traceID)
	} else if result.Data.TaskStatusMsg == "The resolution or aspect ratio of the input image is invalid. Please ensure it meets the required specifications." {
		return nil, traceID, NewErrReason("QueryVideoGeneration failed", 1201,
			fmt.Sprintf("%s, task_id: %s", result.Data.TaskStatusMsg, taskID)).WithTrace(traceID)
	}
	res = &QueryVideoGenerationResult{
		TaskID:    result.Data.TaskID,
		Status:    result.Data.TaskStatus,
		StatusMsg: result.Data.TaskStatusMsg,
		VideoInfo: VideoInfo{FileID: string(result.Data.TaskID)},
	}
	return res, traceID, nil
}

func (s *KlingSDK) ResaveResultVideo(ctx context.Context, downloadURL string) (videoURL string, costTime time.Duration, err error) {
	nowt := time.Now().UTC()
	rsp, err := s.httpc.R().Get(downloadURL)
	if err != nil {
		return "", time.Since(nowt), err
	}
	urlParsed, err := url.Parse(downloadURL)
	if err != nil {
		return "", time.Since(nowt), err
	}
	s3bucket := os.Getenv("S3_BUCKET")
	s3dir := os.Getenv("S3_IMAGE_DIR")
	s3key := fmt.Sprintf("%s/%s_%s", s3dir, lib.GenUniqueID(), filepath.Base(urlParsed.Path))
	s3url, err := s.s3mgr.Upload(ctx, s3bucket, s3key, bytes.NewReader(rsp.Body()), nil)
	if err != nil {
		return "", time.Since(nowt), err
	}
	return s3url, time.Since(nowt), nil
}

func (s *KlingSDK) encodeJWTToken(ak, sk string) (string, error) {
	// 定义 JWT 的 Header 和 Payload
	claims := jwt.MapClaims{
		"iss": ak,                                      // 发行人
		"exp": time.Now().Add(30 * time.Minute).Unix(), // 过期时间 (当前时间 + 30 分钟)
		"nbf": time.Now().Add(-5 * time.Second).Unix(), // 不早于时间 (当前时间 - 5 秒)
	}

	// 创建一个新的 JWT Token，使用 HS256 签名方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用 Secret Key (sk) 进行签名
	tokenString, err := token.SignedString([]byte(sk))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
