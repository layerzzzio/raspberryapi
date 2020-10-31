package api

import (
	"github.com/raspibuddy/rpi/pkg/api/metrics/cpu"
	cl "github.com/raspibuddy/rpi/pkg/api/metrics/cpu/logging"
	cs "github.com/raspibuddy/rpi/pkg/api/metrics/cpu/platform/sys"
	ct "github.com/raspibuddy/rpi/pkg/api/metrics/cpu/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/disk"
	dl "github.com/raspibuddy/rpi/pkg/api/metrics/disk/logging"
	ds "github.com/raspibuddy/rpi/pkg/api/metrics/disk/platform/sys"
	dt "github.com/raspibuddy/rpi/pkg/api/metrics/disk/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/host"
	hl "github.com/raspibuddy/rpi/pkg/api/metrics/host/logging"
	hs "github.com/raspibuddy/rpi/pkg/api/metrics/host/platform/sys"
	ht "github.com/raspibuddy/rpi/pkg/api/metrics/host/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/largestfile"
	lfl "github.com/raspibuddy/rpi/pkg/api/metrics/largestfile/logging"
	lfs "github.com/raspibuddy/rpi/pkg/api/metrics/largestfile/platform/sys"
	lft "github.com/raspibuddy/rpi/pkg/api/metrics/largestfile/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/load"
	ll "github.com/raspibuddy/rpi/pkg/api/metrics/load/logging"
	ls "github.com/raspibuddy/rpi/pkg/api/metrics/load/platform/sys"
	lt "github.com/raspibuddy/rpi/pkg/api/metrics/load/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/mem"
	ml "github.com/raspibuddy/rpi/pkg/api/metrics/mem/logging"
	ms "github.com/raspibuddy/rpi/pkg/api/metrics/mem/platform/sys"
	mt "github.com/raspibuddy/rpi/pkg/api/metrics/mem/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/net"
	nl "github.com/raspibuddy/rpi/pkg/api/metrics/net/logging"
	ns "github.com/raspibuddy/rpi/pkg/api/metrics/net/platform/sys"
	nt "github.com/raspibuddy/rpi/pkg/api/metrics/net/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/process"
	pl "github.com/raspibuddy/rpi/pkg/api/metrics/process/logging"
	ps "github.com/raspibuddy/rpi/pkg/api/metrics/process/platform/sys"
	pt "github.com/raspibuddy/rpi/pkg/api/metrics/process/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/user"
	ul "github.com/raspibuddy/rpi/pkg/api/metrics/user/logging"
	us "github.com/raspibuddy/rpi/pkg/api/metrics/user/platform/sys"
	ut "github.com/raspibuddy/rpi/pkg/api/metrics/user/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/vcore"
	vl "github.com/raspibuddy/rpi/pkg/api/metrics/vcore/logging"
	vs "github.com/raspibuddy/rpi/pkg/api/metrics/vcore/platform/sys"
	vt "github.com/raspibuddy/rpi/pkg/api/metrics/vcore/transport"
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
	lft.NewHTTP(lfl.New(largestfile.New(lfs.LargestFile{}, m), log).Service, v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}