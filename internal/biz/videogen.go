package biz

import "ai-mkt-be/internal/aigc"

type VideoGen struct {
	kling *aigc.KlingSDK
}

func NewVideoGen(kling *aigc.KlingSDK) *VideoGen {
	return &VideoGen{kling: kling}
}
