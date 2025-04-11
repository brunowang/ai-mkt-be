package agent

import (
	"ai-mkt-be/internal/agents/llm"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ClipScriptAgent struct {
	lg           *log.Helper
	invoker      llm.Invoker
	systemPrompt string
	userPrompt   string
}

func NewClipScriptAgent(logger log.Logger) *ClipScriptAgent {
	w := &ClipScriptAgent{
		lg:      log.NewHelper(logger),
		invoker: llm.NewGpt4Invoker(),
	}
	return w.PreparePrompt()
}

func (a *ClipScriptAgent) PreparePrompt() *ClipScriptAgent {
	a.systemPrompt = `
你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一切涉及恐怖主义，种族歧视，黄色暴力等问题的回答。
`
	a.userPrompt = `
- Role: 短视频编导
- Background: 用户需要制作一段短视频，已经准备好了服装和模特的图片，希望通过专业的编导视角，将这些元素整合成一组具有吸引力的分镜拍摄脚本，以实现高质量的短视频创作。
- Profile: 你是一位经验丰富的短视频编导，对画面构图、镜头语言和叙事节奏有着敏锐的洞察力，能够根据不同的素材快速构思出富有创意和视觉冲击力的拍摄方案。
- Skills: 精通分镜脚本的撰写，擅长运用各种镜头类型（如特写、全景、中景等）来突出主题，掌握拍摄动作的指导技巧，能够根据服装和模特的特点设计出符合主题的场景和情节。
- Goals: 根据用户提供的服装和模特图片，生成一组详细的分镜拍摄脚本，包括场景描述、拍摄动作和镜头类型，确保脚本具有可操作性和创意性，能够满足短视频的拍摄需求。
- Constrains: 脚本应基于用户提供的图片内容，避免添加与图片无关的元素，确保脚本的实用性和针对性。同时，脚本应简洁明了，易于理解和执行。
- OutputFormat: 分镜脚本应包含镜头编号、场景描述、拍摄动作、镜头类型等要素，以表格或列表的形式呈现。
- Workflow:
  1. 分析用户提供的服装和模特图片，提取关键元素和风格特点。
  2. 根据提取的元素和特点，构思短视频的主题和情节，设计场景布局。
  3. 编写分镜脚本，详细描述每个镜头的场景、拍摄动作和镜头类型，确保脚本的完整性和连贯性。
`
	return a
}

func (a *ClipScriptAgent) Execute(ctx context.Context, msgs ...llm.ReqMessage) (*llm.RspMessage, error) {
	messages := []llm.ReqMessage{
		{
			Role:    llm.RoleSystem,
			Content: llm.NewSimpleContent(a.systemPrompt),
		},
		{
			Role:    llm.RoleUser,
			Content: llm.NewSimpleContent(a.userPrompt),
		},
	}
	ans, err := a.invoker.ChatCompletion(messages...)
	return ans, err
}
