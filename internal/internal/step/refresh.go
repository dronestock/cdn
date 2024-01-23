package step

import (
	"context"

	"github.com/dronestock/cdn/internal/internal/config"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
)

type Refresh struct {
	config *config.Refresh
	cdn    *cdn.Client
	logger log.Logger
}

func NewRefresh(config *config.Refresh, cdn *cdn.Client, logger log.Logger) *Refresh {
	return &Refresh{
		config: config,
		cdn:    cdn,
		logger: logger,
	}
}

func (r *Refresh) Runnable() bool {
	return "" != r.config.Url || 0 != len(r.config.Urls) || "" != r.config.Path || 0 != len(r.config.Paths)
}

func (r *Refresh) Run(ctx *context.Context) (err error) {
	if pe := r.path(ctx); nil != pe {
		err = pe
	}

	return
}

func (r *Refresh) path(ctx *context.Context) (err error) {
	paths := make([]*string, 0, len(r.config.Paths))
	if "" != r.config.Path {
		paths = append(paths, &r.config.Path)
	}
	for _, path := range r.config.Paths {
		cloned := path
		paths = append(paths, &cloned)
	}

	req := cdn.NewPurgePathCacheRequest()
	req.Paths = paths
	req.FlushType = &r.config.Type
	if _, err = r.cdn.PurgePathCacheWithContext(*ctx, req); nil != err {
		r.logger.Warn("刷新预热目录出错", field.New("paths", paths), field.Error(err))
	}

	return
}
