package tool

type ClipScript struct {
	Content string `json:"content"`
}

type ClipScriptWriter struct {
	systemPrompt string
}

func (w *ClipScriptWriter) PreparePrompt() *ClipScriptWriter {
	w.systemPrompt = `
你是一名资深的短视频编导，擅长服装穿搭方向的拍摄脚本书写。
`
	return w
}

func (w *ClipScriptWriter) WriteClipScript() []ClipScript {
	return nil
}
