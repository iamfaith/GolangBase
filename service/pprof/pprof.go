package pprof

import (
	"net/http"
	libpprof "net/http/pprof"

	"github.com/astaxie/beego"
)

type cpuProfHandler string

func (h cpuProfHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	libpprof.Profile(w, r)
}

func Monitor() {
	adminCtl := adminController{}
	pprof := beego.NewNamespace("/_adm",
		beego.NSNamespace("/pprof",
			beego.NSBefore(adminCtl.Auth),
			beego.NSHandler("/heap", libpprof.Handler("heap")),
			beego.NSHandler("/goroutine", libpprof.Handler("goroutine")),
			beego.NSHandler("/cpu", new(cpuProfHandler)),
		),
	)
	beego.AddNamespace(pprof)
}
