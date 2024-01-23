package config

type Secret struct {
	// 授权，类型于用户名
	Id string `default:"${ID}" validate:"required" json:"id,omitempty"`
	// 授权，类型于密码
	Key string `default:"${KEY}" validate:"required" json:"key,omitempty"`
}
