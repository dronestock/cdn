package internal

import (
	"github.com/dronestock/cdn/internal/internal/config"
	"github.com/dronestock/cdn/internal/internal/step"
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base
	config.Secret
	config.Refresh `default:"${REFRESH}" json:"refresh,omitempty"`
}

func New() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(step.NewRefresh(&p.Secret, &p.Refresh, p.Logger)).Name("刷新预热").Build(),
	}
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("secret", p.Secret),
		field.New("refresh", p.Refresh),
	}
}
