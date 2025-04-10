package service

import (
	"context"

	v1 "ai-mkt-be/api/helloworld/v1"
	"ai-mkt-be/internal/biz"
)

// FilmclipService is a filmclip service.
type FilmclipService struct {
	v1.UnimplementedFilmclipServer

	uc *biz.GreeterUsecase
}

// NewFilmclipService new a filmclip service.
func NewFilmclipService(uc *biz.GreeterUsecase) *FilmclipService {
	return &FilmclipService{uc: uc}
}

func (s *FilmclipService) UploadImage(ctx context.Context, in *v1.UploadImageRequest) (*v1.UploadImageReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.UploadImageReply{Url: "Hello " + g.Hello}, nil
}

func (s *FilmclipService) GenClipScript(context.Context, *v1.GenClipScriptRequest) (*v1.GenClipScriptReply, error) {
	return &v1.GenClipScriptReply{}, nil
}
