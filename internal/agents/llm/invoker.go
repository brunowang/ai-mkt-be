package llm

import (
	v1 "ai-mkt-be/api/filmclip/v1"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-resty/resty/v2"
	"net/http"
	"os"
)

type Invoker interface {
	ChatCompletion(messages ...ReqMessage) (*RspMessage, error)
	Temperature(t float64) Invoker
}

type Gpt4Invoker struct {
	Invoker
}

func NewGpt4Invoker(tools ...ToolDef) *Gpt4Invoker {
	return &Gpt4Invoker{
		Invoker: NewOpenAIInvoker().
			ApiKey(os.Getenv("OPENAI_API_KEY")).
			ApiBase("https://openrouter.ai/api/v1").
			ModelName("openai/gpt-4o").
			WithTools(tools...),
	}
}

type Claude3Invoker struct {
	Invoker
}

func NewClaude3Invoker() *Claude3Invoker {
	return &Claude3Invoker{
		Invoker: NewOpenAIInvoker().
			ApiKey(os.Getenv("OPENAI_API_KEY")).
			ApiBase("https://openrouter.ai/api/v1").
			ModelName("anthropic/claude-3-5-sonnet"),
	}
}

type OpenAIInvoker struct {
	cli     *resty.Client
	apiKey  string
	apiBase string

	modelName   string
	temperature float64
	tools       []ToolDef
}

func NewOpenAIInvoker() *OpenAIInvoker {
	return &OpenAIInvoker{
		cli:         resty.New(),
		temperature: 0.01,
	}
}

func (i *OpenAIInvoker) ApiKey(k string) *OpenAIInvoker {
	if k != "" {
		i.apiKey = k
	}
	return i
}

func (i *OpenAIInvoker) ApiBase(url string) *OpenAIInvoker {
	if url != "" {
		i.apiBase = url
	}
	return i
}

func (i *OpenAIInvoker) ModelName(name string) *OpenAIInvoker {
	if name != "" {
		i.modelName = name
	}
	return i
}

func (i *OpenAIInvoker) Temperature(t float64) Invoker {
	if t > 0.01 {
		i.temperature = t
	}
	return i
}

func (i *OpenAIInvoker) WithTools(tools ...ToolDef) Invoker {
	if len(tools) > 0 {
		i.tools = tools
	}
	return i
}

func (i *OpenAIInvoker) ChatCompletion(messages ...ReqMessage) (*RspMessage, error) {
	body := &OpenAIReq{
		Model:       i.modelName,
		Messages:    messages,
		Tools:       i.tools,
		MaxTokens:   4096,
		Stream:      false,
		Temperature: i.temperature,
	}
	rsp, err := i.cli.R().
		SetHeader("Authorization", "Bearer "+i.apiKey).
		SetBody(body).
		Post(i.apiBase + "/chat/completions")
	if err != nil {
		return nil, err
	} else if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", rsp.StatusCode(), string(rsp.Body()))
	}
	var openAIRsp OpenAIRsp
	if err := json.Unmarshal(rsp.Body(), &openAIRsp); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w, body: %s", err, string(rsp.Body()))
	}
	if len(openAIRsp.Choices) == 0 {
		return nil, errors.InternalServer(v1.ErrorReason_EMPTY_OUTPUT.String(), "got empty output from llm")
	}
	return &openAIRsp.Choices[0].Message, nil
}
