package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	"os"
	"path/filepath"

	v1 "ai-mkt-be/api/helloworld/v1"
	"ai-mkt-be/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/uuid"
)

// FilmclipService is a filmclip service.
type FilmclipService struct {
	v1.UnimplementedFilmclipServer
	logger *log.Helper

	uc *biz.GreeterUsecase
}

// NewFilmclipService new a filmclip service.
func NewFilmclipService(logger log.Logger, uc *biz.GreeterUsecase) *FilmclipService {
	return &FilmclipService{logger: log.NewHelper(logger), uc: uc}
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
	s.logger.WithContext(ctx).Infof("Image uploaded: %s, saved to: %s", in.Name, filePath)

	return &v1.UploadImageReply{Url: url}, nil
}

func (s *FilmclipService) GenClipScript(context.Context, *v1.GenClipScriptRequest) (*v1.GenClipScriptReply, error) {
	return &v1.GenClipScriptReply{}, nil
}
