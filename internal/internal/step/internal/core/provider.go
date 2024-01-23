package core

import (
	"context"

	"github.com/dronestock/cdn/internal/internal/config"
)

type Provider interface {
	Init() error

	Refresh(*context.Context, *config.Refresh) error
}
