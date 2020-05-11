package api

import (
	"github.com/raspibuddy/rpi/pkg/api/cpu"
	cl "github.com/raspibuddy/rpi/pkg/api/cpu/logging"
	cs "github.com/raspibuddy/rpi/pkg/api/cpu/platform/sys"
	ct "github.com/raspibuddy/rpi/pkg/api/cpu/transport"
	"github.com/raspibuddy/rpi/pkg/api/disk"
	dl "github.com/raspibuddy/rpi/pkg/api/disk/logging"
	ds "github.com/raspibuddy/rpi/pkg/api/disk/platform/sys"
	dt "github.com/raspibuddy/rpi/pkg/api/disk/transport"
	"github.com/raspibuddy/rpi/pkg/api/load"
	ll "github.com/raspibuddy/rpi/pkg/api/load/logging"
	ls "github.com/raspibuddy/rpi/pkg/api/load/platform/sys"
	lt "github.com/raspibuddy/rpi/pkg/api/load/transport"
	"github.com/raspibuddy/rpi/pkg/api/mem"
	ml "github.com/raspibuddy/rpi/pkg/api/mem/logging"
	ms "github.com/raspibuddy/rpi/pkg/api/mem/platform/sys"
	mt "github.com/raspibuddy/rpi/pkg/api/mem/transport"
	"github.com/raspibuddy/rpi/pkg/api/vcore"
	vl "github.com/raspibuddy/rpi/pkg/api/vcore/logging"
	vs "github.com/raspibuddy/rpi/pkg/api/vcore/platform/sys"
	vt "github.com/raspibuddy/rpi/pkg/api/vcore/transport"
	"github.com/raspibuddy/rpi/pkg/utl/config"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/server"
	"github.com/raspibuddy/rpi/pkg/utl/zlog"
)

// Start starts the API service.
func Start(cfg *config.Configuration) error {
	e := server.New()
	log := zlog.New()
	v1 := e.Group("/v1")
	m := metrics.Service{}

	ct.NewHTTP(cl.New(cpu.New(cs.CPU{}, m), log).Service, v1)
	vt.NewHTTP(vl.New(vcore.New(vs.VCore{}, m), log).Service, v1)
	mt.NewHTTP(ml.New(mem.New(ms.Mem{}, m), log).Service, v1)
	dt.NewHTTP(dl.New(disk.New(ds.Disk{}, m), log).Service, v1)
	lt.NewHTTP(ll.New(load.New(ls.Load{}, m), log).Service, v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
