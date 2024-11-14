package wenxin

import (
	"github.com/anjude/bcore/pkg/httputil"
	"github.com/anjude/gtools/infrastructure/response"
	"time"
)

var httpClient = httputil.NewHTTPClient("https://api.beanflow.top", time.Minute*2, "")

type AiResp struct {
	response.StandardHttpResp
	Data struct {
		Answer string `json:"answer"`
	}
}

func GetWenXinReply(question, userId string) (string, error) {
	resp := &AiResp{}
	err := httpClient.Request("POST", "/api/so/common/ai/reply", nil, map[string]interface{}{
		"question": question,
		"user_id":  userId,
	}, resp)
	if err != nil {
		return "", err
	}
	return resp.Data.Answer, nil
}
