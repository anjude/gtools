package response

type StandardHttpResp struct {
	ErrCode int32  `json:"err_code"`
	Msg     string `json:"msg"`
	Detail  string `json:"detail"`
}
