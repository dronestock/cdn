package provider

import (
	"context"

	"github.com/dronestock/cdn/internal/internal/config"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type Tencent struct {
	secret *config.Secret
	logger log.Logger

	client *cdn.Client
}

func NewTencent(secret *config.Secret, logger log.Logger) *Tencent {
	return &Tencent{
		secret: secret,
		logger: logger,
	}
}

func (t *Tencent) Init() (err error) {
	credential := common.NewCredential(t.secret.Id, t.secret.Key)
	_profile := profile.NewClientProfile()
	t.client, err = cdn.NewClient(credential, "ap-chengdu", _profile)

	return
}

func (t *Tencent) Refresh(ctx *context.Context, config *config.Refresh) (err error) {
	if pe := t.path(ctx, config); nil != pe {
		err = pe
	}

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
	if _, err = t.client.PurgePathCacheWithContext(*ctx, req); nil != err {
		t.logger.Warn("刷新预热目录出错", field.New("paths", paths), field.Error(err))
	}

	return
}
