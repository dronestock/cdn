package main

type domain struct {
	// 域名
	Name string `json:"name"`
	// 提供商
	Provider provider `default:"tencentyun" json:"provider"`
	// 应用用户名
	Id string `json:"id" validate:"required"`
	// 应用用户名
	Key string `json:"key" validate:"required"`
	// 协议
	Protocol string `default:"https" json:"protocol" validate:"oneof=https http"`
}
