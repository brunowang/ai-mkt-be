package aigc

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	KlingBaseResp struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	}

	InvokeVideoGenerationParams struct {
		TaskID          string   `json:"task_id"`
		PromptKey       string   `json:"prompt_key"`
		Prompt          string   `json:"prompt"`
		NegativePrompt  string   `json:"negative_prompt"`
		FirstFrameImage string   `json:"first_frame_image"`
		EffectImageList []string `json:"effect_image_list"`
		CfgScale        float32  `json:"cfg_scale"`
		AspectRatio     string   `json:"aspect_ratio"`
	}

	InvokeVideoGenerationResult struct {
		TaskID TaskID `json:"task_id"`
	}

	invokeKlingResult struct {
		KlingBaseResp `json:",inline"`

		Data InvokeVideoGenerationResult `json:"data"`
	}
)

type (
	VideoInfo struct {
		FileID      string `json:"file_id"`
		VideoUrl    string `json:"video_url"`
		VideoWidth  int    `json:"video_width"`
		VideoHeight int    `json:"video_height"`
	}

	QueryVideoGenerationResult struct {
		TaskID    TaskID `json:"task_id"`
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
		VideoInfo `json:",inline"`
	}
)

type (
	FetchVideoGenerationResult struct {
		ResaveCost  time.Duration `json:"resave_cost"`
		DownloadURL string        `json:"download_url"`
	}

	fetchKlingResult struct {
		KlingBaseResp `json:",inline"`

		Data fetchKlingResultData `json:"data"`
	}

	fetchKlingResultData struct {
		TaskID        TaskID `json:"task_id"`
		TaskStatus    string `json:"task_status"`
		TaskStatusMsg string `json:"task_status_msg"`
		TaskResult    struct {
			Videos []fetchKlingResultDataVideo `json:"videos"`
		} `json:"task_result"`
	}

	fetchKlingResultDataVideo struct {
		Id       string `json:"id"`
		Url      string `json:"url"`
		Duration string `json:"duration"`
	}
)

type TaskID string

func (t *TaskID) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*t = TaskID(str)
		return nil
	}

	var num json.Number
	if err := json.Unmarshal(data, &num); err == nil {
		*t = TaskID(num.String())
		return nil
	}

	return fmt.Errorf("invalid task_id format")
}
