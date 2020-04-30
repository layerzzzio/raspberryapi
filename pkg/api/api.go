package api

import (
	"github.com/raspibuddy/rpi/pkg/api/cpu"
	cl "github.com/raspibuddy/rpi/pkg/api/cpu/logging"
	"github.com/raspibuddy/rpi/pkg/api/cpu/platform/system"
	ct "github.com/raspibuddy/rpi/pkg/api/cpu/transport"
	"github.com/raspibuddy/rpi/pkg/api/vcore"
	vl "github.com/raspibuddy/rpi/pkg/api/vcore/logging"
	vt "github.com/raspibuddy/rpi/pkg/api/vcore/transport"
	"github.com/raspibuddy/rpi/utl/config"
	"github.com/raspibuddy/rpi/utl/server"
	"github.com/raspibuddy/rpi/utl/zlog"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	e := server.New()

	log := zlog.New()

	v1 := e.Group("/v1")

	ct.NewHTTP(cl.New(cpu.New(system.CPU{}), log).Service, v1)
	vt.NewHTTP(vl.New(vcore.New(), log).Service, v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
