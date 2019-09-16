package response

type Response struct {
	RunTime float64     `json:"run_time"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
