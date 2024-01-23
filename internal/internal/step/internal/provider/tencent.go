package provider

import (
	"context"

	"github.com/dronestock/cdn/internal/internal/config"
	"github.com/goexl/gox/field"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type Tencent struct {
	secret *config.Secret
	client *cdn.Client
}

func (t *Tencent) Init() (err error) {
	credential := common.NewCredential(t.secret.Id, t.secret.Key)
	_profile := profile.NewClientProfile()
	t.client, err = cdn.NewClient(credential, "ap-chengdu", _profile)

	return
}

func (t *Tencent) Refresh(ctx *context.Context, config *config.Refresh) (err error) {
	return
}
func (t *Tencent) path(ctx *context.Context, config *config.Refresh) (err error) {
	paths := make([]*string, 0, len(config.Paths))
	if "" != config.Path {
		paths = append(paths, &config.Path)
	}
	for _, path := range config.Paths {
		cloned := path
		paths = append(paths, &cloned)
	}

	req := cdn.NewPurgePathCacheRequest()
	req.Paths = paths
	req.FlushType = &config.Type
	if _, err = cdn.PurgePathCacheWithContext(*ctx, req); nil != err {
		t.logger.Warn("刷新预热目录出错", field.New("paths", paths), field.Error(err))
	}

	return
}
