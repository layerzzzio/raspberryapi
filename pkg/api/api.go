package api

import (
	"github.com/raspibuddy/rpi/pkg/api/actions/appaction"
	aal "github.com/raspibuddy/rpi/pkg/api/actions/appaction/logging"
	aas "github.com/raspibuddy/rpi/pkg/api/actions/appaction/platform/sys"
	aat "github.com/raspibuddy/rpi/pkg/api/actions/appaction/transport"
	"github.com/raspibuddy/rpi/pkg/api/actions/appinstall"
	ail "github.com/raspibuddy/rpi/pkg/api/actions/appinstall/logging"
	ais "github.com/raspibuddy/rpi/pkg/api/actions/appinstall/platform/sys"
	ait "github.com/raspibuddy/rpi/pkg/api/actions/appinstall/transport"
	"github.com/raspibuddy/rpi/pkg/api/actions/configure"
	acl "github.com/raspibuddy/rpi/pkg/api/actions/configure/logging"
	acs "github.com/raspibuddy/rpi/pkg/api/actions/configure/platform/sys"
	act "github.com/raspibuddy/rpi/pkg/api/actions/configure/transport"
	"github.com/raspibuddy/rpi/pkg/api/actions/destroy"
	adl "github.com/raspibuddy/rpi/pkg/api/actions/destroy/logging"
	ads "github.com/raspibuddy/rpi/pkg/api/actions/destroy/platform/sys"
	adt "github.com/raspibuddy/rpi/pkg/api/actions/destroy/transport"
	"github.com/raspibuddy/rpi/pkg/api/actions/general"
	agl "github.com/raspibuddy/rpi/pkg/api/actions/general/logging"
	ags "github.com/raspibuddy/rpi/pkg/api/actions/general/platform/sys"
	agt "github.com/raspibuddy/rpi/pkg/api/actions/general/transport"
	"github.com/raspibuddy/rpi/pkg/api/admin/deployment"
	del "github.com/raspibuddy/rpi/pkg/api/admin/deployment/logging"
	des "github.com/raspibuddy/rpi/pkg/api/admin/deployment/platform/sys"
	det "github.com/raspibuddy/rpi/pkg/api/admin/deployment/transport"
	"github.com/raspibuddy/rpi/pkg/api/infos/appconfig"
	iacl "github.com/raspibuddy/rpi/pkg/api/infos/appconfig/logging"
	iacs "github.com/raspibuddy/rpi/pkg/api/infos/appconfig/platform/sys"
	iact "github.com/raspibuddy/rpi/pkg/api/infos/appconfig/transport"
	"github.com/raspibuddy/rpi/pkg/api/infos/appstatus"
	iasl "github.com/raspibuddy/rpi/pkg/api/infos/appstatus/logging"
	iass "github.com/raspibuddy/rpi/pkg/api/infos/appstatus/platform/sys"
	iast "github.com/raspibuddy/rpi/pkg/api/infos/appstatus/transport"
	"github.com/raspibuddy/rpi/pkg/api/infos/boot"
	ibol "github.com/raspibuddy/rpi/pkg/api/infos/boot/logging"
	ibos "github.com/raspibuddy/rpi/pkg/api/infos/boot/platform/sys"
	ibot "github.com/raspibuddy/rpi/pkg/api/infos/boot/transport"
	"github.com/raspibuddy/rpi/pkg/api/infos/configfile"
	icol "github.com/raspibuddy/rpi/pkg/api/infos/configfile/logging"
	icos "github.com/raspibuddy/rpi/pkg/api/infos/configfile/platform/sys"
	icot "github.com/raspibuddy/rpi/pkg/api/infos/configfile/transport"
	"github.com/raspibuddy/rpi/pkg/api/infos/display"
	idil "github.com/raspibuddy/rpi/pkg/api/infos/display/logging"
	idis "github.com/raspibuddy/rpi/pkg/api/infos/display/platform/sys"
	idit "github.com/raspibuddy/rpi/pkg/api/infos/display/transport"
	"github.com/raspibuddy/rpi/pkg/api/infos/humanuser"
	ihul "github.com/raspibuddy/rpi/pkg/api/infos/humanuser/logging"
	ihus "github.com/raspibuddy/rpi/pkg/api/infos/humanuser/platform/sys"
	ihut "github.com/raspibuddy/rpi/pkg/api/infos/humanuser/transport"
	"github.com/raspibuddy/rpi/pkg/api/infos/port"
	ptl "github.com/raspibuddy/rpi/pkg/api/infos/port/logging"
	pts "github.com/raspibuddy/rpi/pkg/api/infos/port/platform/sys"
	ptt "github.com/raspibuddy/rpi/pkg/api/infos/port/transport"
	"github.com/raspibuddy/rpi/pkg/api/infos/rpinterface"
	iinl "github.com/raspibuddy/rpi/pkg/api/infos/rpinterface/logging"
	iins "github.com/raspibuddy/rpi/pkg/api/infos/rpinterface/platform/sys"
	iint "github.com/raspibuddy/rpi/pkg/api/infos/rpinterface/transport"
	"github.com/raspibuddy/rpi/pkg/api/infos/software"
	isol "github.com/raspibuddy/rpi/pkg/api/infos/software/logging"
	isos "github.com/raspibuddy/rpi/pkg/api/infos/software/platform/sys"
	isot "github.com/raspibuddy/rpi/pkg/api/infos/software/transport"
	"github.com/raspibuddy/rpi/pkg/api/infos/version"
	vel "github.com/raspibuddy/rpi/pkg/api/infos/version/logging"
	ves "github.com/raspibuddy/rpi/pkg/api/infos/version/platform/sys"
	vet "github.com/raspibuddy/rpi/pkg/api/infos/version/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/cpu"
	cl "github.com/raspibuddy/rpi/pkg/api/metrics/cpu/logging"
	cs "github.com/raspibuddy/rpi/pkg/api/metrics/cpu/platform/sys"
	ct "github.com/raspibuddy/rpi/pkg/api/metrics/cpu/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/disk"
	dl "github.com/raspibuddy/rpi/pkg/api/metrics/disk/logging"
	ds "github.com/raspibuddy/rpi/pkg/api/metrics/disk/platform/sys"
	dt "github.com/raspibuddy/rpi/pkg/api/metrics/disk/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/filestructure"
	fsl "github.com/raspibuddy/rpi/pkg/api/metrics/filestructure/logging"
	fss "github.com/raspibuddy/rpi/pkg/api/metrics/filestructure/platform/sys"
	fst "github.com/raspibuddy/rpi/pkg/api/metrics/filestructure/transport"
	"github.com/raspibuddy/rpi/pkg/api/metrics/host"
	hl "github.com/raspibuddy/rpi/pkg/api/metrics/host/logging"
	hs "github.com/raspibuddy/rpi/pkg/api/metrics/host/platform/sys"
	ht "github.com/raspibuddy/rpi/pkg/api/metrics/host/transport"
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
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/config"
	"github.com/raspibuddy/rpi/pkg/utl/infos"
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
	a := actions.New()
	i := infos.New()

	// metrics
	ct.NewHTTP(cl.New(cpu.New(cs.CPU{}, m), log).Service, v1)
	vt.NewHTTP(vl.New(vcore.New(vs.VCore{}, m), log).Service, v1)
	mt.NewHTTP(ml.New(mem.New(ms.Mem{}, m), log).Service, v1)
	dt.NewHTTP(dl.New(disk.New(ds.Disk{}, m), log).Service, v1)
	lt.NewHTTP(ll.New(load.New(ls.Load{}, m), log).Service, v1)
	pt.NewHTTP(pl.New(process.New(ps.Process{}, m), log).Service, v1)
	ht.NewHTTP(hl.New(host.New(hs.Host{}, m), log).Service, v1)
	ut.NewHTTP(ul.New(user.New(us.User{}, m), log).Service, v1)
	nt.NewHTTP(nl.New(net.New(ns.Net{}, m), log).Service, v1)
	fst.NewHTTP(fsl.New(filestructure.New(fss.FileStructure{}, m), log).Service, v1)

	// actions
	adt.NewHTTP(adl.New(destroy.New(ads.Destroy{}, a), log).Service, v1)
	agt.NewHTTP(agl.New(general.New(ags.General{}, a), log).Service, v1)
	act.NewHTTP(acl.New(configure.New(acs.Configure{}, a, i), log).Service, v1)
	ait.NewHTTP(ail.New(appinstall.New(ais.Install{}, a, i), log).Service, v1)
	aat.NewHTTP(aal.New(appaction.New(aas.AppAction{}, a, i), log).Service, v1)

	// infos
	ihut.NewHTTP(ihul.New(humanuser.New(ihus.HumanUser{}, i), log).Service, v1)
	ibot.NewHTTP(ibol.New(boot.New(ibos.Boot{}, i), log).Service, v1)
	idit.NewHTTP(idil.New(display.New(idis.Display{}, i), log).Service, v1)
	icot.NewHTTP(icol.New(configfile.New(icos.ConfigFile{}, i), log).Service, v1)
	iint.NewHTTP(iinl.New(rpinterface.New(iins.RpInterface{}, i), log).Service, v1)
	isot.NewHTTP(isol.New(software.New(isos.Software{}, i), log).Service, v1)
	iact.NewHTTP(iacl.New(appconfig.New(iacs.AppConfigVPNWithOvpn{}, i), log).Service, v1)
	iast.NewHTTP(iasl.New(appstatus.New(iass.AppStatus{}, i), log).Service, v1)
	ptt.NewHTTP(ptl.New(port.New(pts.Port{}, i), log).Service, v1)

	// admin
	vet.NewHTTP(vel.New(version.New(ves.Version{}, i), log).Service, v1)
	det.NewHTTP(del.New(deployment.New(des.Deployment{}, a), log).Service, v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
