package response

import "time"

type res struct {
	Status         string      `json:"status"`
	Message        string      `json:"message,omitempty"`
	Description    string      `json:"description,omitempty" `
	ServerResponse string      `json:"server_response_at"`
	Result         interface{} `json:"result"`
}

func ResponseSuccess(result interface{}) interface{} {
	res := res{
		Status:         "200",
		Message:        "success",
		ServerResponse: time.Now().Format("2006-01-02 03:04:05"),
	}
	res.Result = result

	return res
}
