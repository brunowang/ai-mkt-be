package llm

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"regexp"
	"strings"
)

const (
	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
	RoleTool      = "tool"
)

type OpenAIReq struct {
	Model       string       `json:"model"`
	Messages    []ReqMessage `json:"messages"`
	Tools       []ToolDef    `json:"tools,omitempty"`
	MaxTokens   int          `json:"max_tokens,omitempty"`
	Stream      bool         `json:"stream"`
	Temperature float64      `json:"temperature"`
}

type OpenAIRsp struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Index        int        `json:"index"`
	Delta        RspMessage `json:"delta"`
	Message      RspMessage `json:"message"`
	FinishReason string     `json:"finish_reason"`
}

type ReqMessage struct {
	Role    string  `json:"role"`
	Content Content `json:"content"`
}

type RspMessage struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
	Name       string     `json:"name,omitempty"`
}

type Content interface {
	String() string
}

type SimpleContent string

func NewSimpleContent(text string) SimpleContent {
	return SimpleContent(text)
}

func (c SimpleContent) String() string {
	return string(c)
}

type MultiContent []SubContent

func (c MultiContent) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

type SubContent interface {
	GetType() string
	String() string
}

type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewTextContent(text string) TextContent {
	return TextContent{Type: "text", Text: text}
}

func (c TextContent) GetType() string {
	return c.Type
}

func (c TextContent) String() string {
	return c.Text
}

type ImageContent struct {
	Type     string `json:"type"`
	ImageUrl struct {
		Url    string `json:"url"`
		Detail string `json:"detail"`
	} `json:"image_url"`
}

func NewImageContent(imageUrl string) ImageContent {
	return ImageContent{
		Type: "image_url",
		ImageUrl: struct {
			Url    string `json:"url"`
			Detail string `json:"detail"`
		}{
			Url:    imageUrl,
			Detail: "low",
		},
	}
}

func (c ImageContent) GetType() string {
	return c.Type
}

func (c ImageContent) String() string {
	return c.ImageUrl.Url
}

type ToolDef struct {
	Type    string   `json:"type"`
	FuncDef *FuncDef `json:"function,omitempty"`
}

type FuncDef struct {
	Name   string         `json:"name"`
	Desc   string         `json:"description"`
	Params map[string]any `json:"parameters,omitempty"`
}

type ToolCall struct {
	ID       string    `json:"id"`
	Type     string    `json:"type"`
	FuncCall *FuncCall `json:"function,omitempty"`
}

type FuncCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type History []ReqMessage

func (h History) String() string {
	if len(h) == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("['")
	for i, v := range h {
		if i > 0 {
			sb.WriteString("','")
		}
		sb.WriteString(v.Content.String())
	}
	sb.WriteString("']")
	return sb.String()
}

func (h History) LastQ() string {
	if len(h) == 0 {
		return ""
	}
	for i := len(h) - 1; i >= 0; i-- {
		if h[i].Role == RoleUser {
			return h[i].Content.String()
		}
	}
	return ""
}

// ExtractJSONFromText 使用正则表达式从文本中提取JSON部分
func ExtractJSONFromText(text string) (string, error) {
	// 匹配```json\n和\n```之间的内容
	re := regexp.MustCompile("```json\\s*\\n([\\s\\S]*?)\\n\\s*```")
	matches := re.FindStringSubmatch(text)

	if len(matches) < 2 {
		return "", errors.New(500, "JSON_DECODE_ERROR", "json part not found")
	}

	// 返回第一个捕获组的内容（即JSON部分）
	return matches[1], nil
}
