package step

import (
	"context"

	"github.com/dronestock/cdn/internal/internal/config"
	"github.com/dronestock/cdn/internal/internal/step/internal/core"
	"github.com/dronestock/cdn/internal/internal/step/internal/provider"
	"github.com/goexl/log"
)

type Refresh struct {
	secret *config.Secret
	config *config.Refresh
	logger log.Logger
}

func NewRefresh(secret *config.Secret, config *config.Refresh, logger log.Logger) *Refresh {
	return &Refresh{
		secret: secret,
		config: config,
		logger: logger,
	}
}

func (r *Refresh) Runnable() bool {
	return "" != r.config.Url || 0 != len(r.config.Urls) || "" != r.config.Path || 0 != len(r.config.Paths)
}

func (r *Refresh) Run(ctx *context.Context) (err error) {
	var refresher core.Refresher
	switch r.config.Provider {
	case "tencent":
		refresher = provider.NewTencent(r.secret, r.logger)
	default:
		refresher = provider.NewTencent(r.secret, r.logger)
	}
	if ie := refresher.Init(); nil != ie {
		err = ie
	} else if re := refresher.Refresh(ctx, r.config); nil != re {
		err = re
	}

	return
}
