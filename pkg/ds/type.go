package ds

import (
	"embed"

	"github.com/hashicorp/go-hclog"

	"github.com/gizmo-platform/gizmo/pkg/firmware"
)

// DriverStation binds all methods related to the driver station task,
// which is a complicated component consisting of service supervision,
// configuration, and all the normal components that make up a field
// server.
type DriverStation struct {
	l hclog.Logger

	cfg firmware.Config

	svc *runit
}

// Option enables variadic configuration of the driver's station
// components.
type Option func(*DriverStation)

// A ConfigureStep performa various changes to the system to configure
// it.
type ConfigureStep func() error

//go:embed tpl/*.tpl
var efs embed.FS
