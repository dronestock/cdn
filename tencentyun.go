package main

import (
	"github.com/goexl/gox/field"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func (p *plugin) tencentyun(domain *domain) (err error) {
	credential := common.NewCredential(domain.Id, domain.Key)
	_profile := profile.NewClientProfile()

	if client, ce := cdn.NewClient(credential, p.Regin, _profile); nil != err {
		err = ce
	} else if ue := p.tencentyunUrls(domain, client); nil != ue {
		err = ue
	} else if pe := p.tencentyunPaths(domain, client); nil != pe {
		err = pe
	}

	return
}

func (p *plugin) tencentyunUrls(domain *domain, client *cdn.Client) (err error) {
	if 0 == len(p.urls[domain.Name]) {
		return
	}

	req := cdn.NewPurgeUrlsCacheRequest()
	req.Urls = p.urls[domain.Name]
	if rsp, pue := client.PurgeUrlsCache(req); nil != pue {
		err = pue
	} else {
		rf := field.New("request.id", rsp.Response.RequestId)
		tf := field.New("task.id", rsp.Response.TaskId)
		p.Debug("刷新地址成功", rf, tf)
	}

	return
}

func (p *plugin) tencentyunPaths(domain *domain, client *cdn.Client) (err error) {
	if 0 == len(p.directories[domain.Name]) {
		return
	}

	req := cdn.NewPurgePathCacheRequest()
	req.Paths = p.directories[domain.Name]
	if rsp, ppe := client.PurgePathCache(req); nil != ppe {
		err = ppe
	} else {
		rf := field.New("request.id", rsp.Response.RequestId)
		tf := field.New("task.id", rsp.Response.TaskId)
		p.Debug("刷新目录成功", rf, tf)
	}

	return
}
