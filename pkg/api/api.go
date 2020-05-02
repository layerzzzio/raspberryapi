package api

import (
	"github.com/raspibuddy/rpi/pkg/api/cpu"
	cl "github.com/raspibuddy/rpi/pkg/api/cpu/logging"
	cs "github.com/raspibuddy/rpi/pkg/api/cpu/platform/sys"
	ct "github.com/raspibuddy/rpi/pkg/api/cpu/transport"
	"github.com/raspibuddy/rpi/pkg/api/vcore"
	vl "github.com/raspibuddy/rpi/pkg/api/vcore/logging"
	vs "github.com/raspibuddy/rpi/pkg/api/vcore/platform/sys"
	vt "github.com/raspibuddy/rpi/pkg/api/vcore/transport"
	"github.com/raspibuddy/rpi/pkg/utl/config"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/raspibuddy/rpi/pkg/utl/zlog"
)

// Start starts the API service.
func Start(cfg *config.Configuration) error {
	e := server.New()

	log := zlog.New()

	v1 := e.Group("/v1")

	ct.NewHTTP(cl.New(cpu.New(cs.CPU{}), log).Service, v1)
	vt.NewHTTP(vl.New(vcore.New(vs.VCore{}), log).Service, v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
