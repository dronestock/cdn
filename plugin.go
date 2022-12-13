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

	// 区域
	Regin string `default:"${REGIN}"`
	// 域名
	Domain *domain `default:"${DOMAIN}" validate:"required_without=Domains"`
	// 域名列表
	Domains []*domain `default:"${DOMAINS}"`
	// 地址
	Url string `default:"${URL=/}" validate:"required_without=Urls"`
	// 地址列表
	Urls []string `default:"${URLS}" validate:"required"`

	domains     []*domain
	urls        map[string][]*string
	directories map[string][]*string
}

func newPlugin() drone.Plugin {
	return &plugin{
		domains:     make([]*domain, 0, 1),
		urls:        make(map[string][]*string),
		directories: make(map[string][]*string),
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
	if nil != p.Domain {
		p.domains = append(p.domains, p.Domain)
	}
	p.domains = append(p.domains, p.Domains...)

	for _, uri := range p.Urls {
		uri = strings.TrimPrefix(uri, "/")
		for _, domain := range p.domains {
			url := fmt.Sprintf("%s://%s/%s", domain.Protocol, domain.Name, uri)
			if _, ok := p.directories[domain.Name]; !ok {
				p.directories[domain.Name] = make([]*string, 0)
			}
			if _, ok := p.urls[domain.Name]; !ok {
				p.urls[domain.Name] = make([]*string, 0)
			}

			if strings.HasSuffix(uri, "*") {
				p.directories[domain.Name] = append(p.directories[domain.Name], &url)
			} else {
				p.urls[domain.Name] = append(p.urls[domain.Name], &url)
			}
		}
	}

	return
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("domains", p.domains),
		field.New("uris", p.Urls),
	}
}
