package ReturnMessage

import "encoding/json"

type ApifoxModal struct {
	AccessToken string `json:"access_token"` // 获取到的凭证，最长为512字节
	Errcode     int64  `json:"errcode"`      // 出错返回码，为0表示成功，非0表示调用失败
	Errmsg      string `json:"errmsg"`       // 返回码提示语
	ExpiresIn   int64  `json:"expires_in"`   // 凭证的有效时间（秒）
}

func (a ApifoxModal) Marshal(data []byte) ([]byte, error) {
	err := json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(a)
	return marshal, err
}
