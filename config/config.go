package config

type Client struct {
	Corpid     string `json:"corpid"`     // 企业ID
	CorPsecret string `json:"corpsecret"` // 应用ID
}
