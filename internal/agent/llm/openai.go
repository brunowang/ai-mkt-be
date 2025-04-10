package llm

import (
	"strings"
)

const (
	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
	RoleTool      = "tool"
)

type OpenAIReq struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Tools       []ToolDef `json:"tools,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Stream      bool      `json:"stream"`
	Temperature float64   `json:"temperature"`
}

type OpenAIRsp struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Index        int     `json:"index"`
	Delta        Message `json:"delta"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
	Name       string     `json:"name,omitempty"`
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

type History []Message

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
		sb.WriteString(v.Content)
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
			return h[i].Content
		}
	}
	return ""
}
