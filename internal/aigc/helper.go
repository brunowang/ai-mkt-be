package aigc

import "encoding/json"

type ErrReason struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Trace string `json:"trace,omitempty"`
	Type  string `json:"type"`
}

func NewErrReason(typ string, code int, msg string) *ErrReason {
	return &ErrReason{Type: typ, Code: code, Msg: msg}
}

func (e *ErrReason) JsonStr() string {
	if e == nil {
		return "{}"
	}
	bs, _ := json.Marshal(e)
	return string(bs)
}

func (e *ErrReason) GetTrace() string {
	if e == nil {
		return ""
	}
	return e.Trace
}

func (e *ErrReason) WithTrace(trace string) *ErrReason {
	if e == nil {
		return &ErrReason{Trace: trace}
	}
	e.Trace = trace
	return e
}

func (e *ErrReason) WithoutTrace() *ErrReason {
	if e == nil {
		return &ErrReason{}
	}
	cp := *e
	cp.Trace = ""
	return &cp
}

func (e *ErrReason) Error() string {
	return e.JsonStr()
}
