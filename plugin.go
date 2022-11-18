package main

import (
	"fmt"
	"strings"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 提供商
	Provider provider `default:"${PROVIDER=tencentyun}" validate:"required,oneof=tencentyun aliyun"`
	// 应用用户名
	Id string `default:"${ID}" validate:"required_if=Provider tencentyun"`
	// 应用用户名
	Key string `default:"${KEY}" validate:"required_if=Provider tencentyun"`
	// 区域
	Regin string `default:"${REGIN}"`
	// 域名
	Domain string `default:"${DOMAIN}" validate:"required_without=Domains"`
	// 域名列表
	Domains []string `default:"${DOMAINS}"`
	// 协议
	Protocol string `default:"${PROTOCOL=https}" validate:"oneof=https http"`
	// 地址列表
	Uris []string `default:"${URIS}" validate:"required"`

	domains     []string
	urls        []*string
	directories []*string
}

func newPlugin() drone.Plugin {
	return &plugin{
		domains:     make([]string, 0, 1),
		urls:        make([]*string, 0, 1),
		directories: make([]*string, 0, 1),
	}
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(p.refresh, drone.Name("刷新")),
	}
}

func (p *plugin) Setup() (unset bool, err error) {
	if "" != p.Domain {
		p.domains = append(p.domains, p.Domain)
	}
	p.domains = append(p.domains, p.Domains...)

	for _, uri := range p.Uris {
		if "/" == uri {
			uri = ""
		}
		if strings.HasPrefix(uri, "/") {
			uri = uri[1:]
		}

		for _, domain := range p.domains {
			url := fmt.Sprintf("%s://%s/%s", p.Protocol, domain, uri)
			if strings.HasSuffix(uri, "*") {
				p.directories = append(p.directories, &url)
			} else {
				p.urls = append(p.urls, &url)
			}
		}
	}

	return
}

func (p *plugin) Fields() gox.Fields {
	return gox.Fields{
		field.String("provider", string(p.Provider)),
		field.Strings("uris", p.Uris...),
	}
}
