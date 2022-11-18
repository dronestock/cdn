package main

import (
	"github.com/goexl/gox/field"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func (p *plugin) tencentyun() (err error) {
	credential := common.NewCredential(p.Id, p.Key)
	_profile := profile.NewClientProfile()

	if client, ce := cdn.NewClient(credential, p.Regin, _profile); nil != err {
		err = ce
	} else if ue := p.tencentyunUrls(client); nil != ue {
		err = ue
	} else if pe := p.tencentyunPaths(client); nil != pe {
		err = pe
	}

	return
}

func (p *plugin) tencentyunUrls(client *cdn.Client) (err error) {
	if 0 == len(p.directories) {
		return
	}

	urls := cdn.NewPurgeUrlsCacheRequest()
	urls.Urls = p.urls
	if rsp, pue := client.PurgeUrlsCache(urls); nil != pue {
		err = pue
	} else {
		request := field.Stringp("request", rsp.Response.RequestId)
		task := field.Stringp("task", rsp.Response.TaskId)
		p.Debug("刷新地址成功", request, task)
	}

	return
}

func (p *plugin) tencentyunPaths(client *cdn.Client) (err error) {
	if 0 == len(p.directories) {
		return
	}

	paths := cdn.NewPurgePathCacheRequest()
	paths.Paths = p.directories
	if rsp, ppe := client.PurgePathCache(paths); nil != ppe {
		err = ppe
	} else {
		request := field.Stringp("request", rsp.Response.RequestId)
		task := field.Stringp("task", rsp.Response.TaskId)
		p.Debug("刷新目录成功", request, task)
	}

	return
}
