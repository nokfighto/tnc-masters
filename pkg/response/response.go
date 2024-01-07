package response

type x struct {
	Status      string      `json:"status"`
	Message     string      `json:"message,omitempty"`
	Description string      `json:"description,omitempty" `
	Result      interface{} `json:"result"`
}

func ResponseSuccess(result interface{}) interface{} {
	res := x{
		Status: "ok",
	}
	res.Result = result

	return res
}
