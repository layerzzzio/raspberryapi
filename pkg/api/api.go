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
	"github.com/raspibuddy/rpi/pkg/api/host"
	hl "github.com/raspibuddy/rpi/pkg/api/host/logging"
	hs "github.com/raspibuddy/rpi/pkg/api/host/platform/sys"
	ht "github.com/raspibuddy/rpi/pkg/api/host/transport"
	"github.com/raspibuddy/rpi/pkg/api/load"
	ll "github.com/raspibuddy/rpi/pkg/api/load/logging"
	ls "github.com/raspibuddy/rpi/pkg/api/load/platform/sys"
	lt "github.com/raspibuddy/rpi/pkg/api/load/transport"
	"github.com/raspibuddy/rpi/pkg/api/mem"
	ml "github.com/raspibuddy/rpi/pkg/api/mem/logging"
	ms "github.com/raspibuddy/rpi/pkg/api/mem/platform/sys"
	mt "github.com/raspibuddy/rpi/pkg/api/mem/transport"
	"github.com/raspibuddy/rpi/pkg/api/net"
	nl "github.com/raspibuddy/rpi/pkg/api/net/logging"
	ns "github.com/raspibuddy/rpi/pkg/api/net/platform/sys"
	nt "github.com/raspibuddy/rpi/pkg/api/net/transport"
	"github.com/raspibuddy/rpi/pkg/api/process"
	pl "github.com/raspibuddy/rpi/pkg/api/process/logging"
	ps "github.com/raspibuddy/rpi/pkg/api/process/platform/sys"
	pt "github.com/raspibuddy/rpi/pkg/api/process/transport"
	"github.com/raspibuddy/rpi/pkg/api/user"
	ul "github.com/raspibuddy/rpi/pkg/api/user/logging"
	us "github.com/raspibuddy/rpi/pkg/api/user/platform/sys"
	ut "github.com/raspibuddy/rpi/pkg/api/user/transport"
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
	m := metrics.New(metrics.Service{})

	ct.NewHTTP(cl.New(cpu.New(cs.CPU{}, m), log).Service, v1)
	vt.NewHTTP(vl.New(vcore.New(vs.VCore{}, m), log).Service, v1)
	mt.NewHTTP(ml.New(mem.New(ms.Mem{}, m), log).Service, v1)
	dt.NewHTTP(dl.New(disk.New(ds.Disk{}, m), log).Service, v1)
	lt.NewHTTP(ll.New(load.New(ls.Load{}, m), log).Service, v1)
	pt.NewHTTP(pl.New(process.New(ps.Process{}, m), log).Service, v1)
	ht.NewHTTP(hl.New(host.New(hs.Host{}, m), log).Service, v1)
	ut.NewHTTP(ul.New(user.New(us.User{}, m), log).Service, v1)
	nt.NewHTTP(nl.New(net.New(ns.Net{}, m), log).Service, v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
