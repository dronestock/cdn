package core

import (
	"context"

	"github.com/dronestock/cdn/internal/internal/config"
)

type Refresher interface {
	Initiate

	Refresh(*context.Context, *config.Refresh) error
}
