package biz

import (
	v1 "ai-mkt-be/api/filmclip/v1"
	"ai-mkt-be/internal/agents/agent"
	"ai-mkt-be/internal/agents/llm"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrAgentNotFound = errors.NotFound(v1.ErrorReason_AGENT_NOT_FOUND.String(), "agent not found")
)

type Agent interface {
	Execute(ctx context.Context, msgs ...llm.ReqMessage) (*llm.RspMessage, error)
}

type AgentGraph struct {
	agents map[v1.Intent]Agent
}

func NewAgentGraph(logger log.Logger) *AgentGraph {
	graph := map[v1.Intent]Agent{
		v1.Intent_GenClipScript: agent.NewClipScriptAgent(logger),
	}
	return &AgentGraph{agents: graph}
}

func (a *AgentGraph) GetAgent(intent v1.Intent) (Agent, error) {
	ret, ok := a.agents[intent]
	if !ok {
		return nil, ErrAgentNotFound
	}
	return ret, nil
}
