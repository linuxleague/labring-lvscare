package create

import (
	"github.com/fanux/lvscare/service"
	"github.com/wonderivan/logger"
)

//VsAndRsCreate is
func VsAndRsCreate(vs string, rs []string, lvs service.Lvser) {
	//ip, _ := utils.SplitServer(vs)
	if lvs == nil {
		lvs = service.BuildLvscare()
	}
	var errs []string
	err := lvs.CreateVirtualServer(vs, true)
	//virtual server is exists
	if err != nil {
		//can't return
		errs = append(errs, err.Error())
	}
	for _, r := range rs {
		err = lvs.CreateRealServer(r, true)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) != 0 {
		logger.Error("VsAndRsCreate error:", errs)
	}

}
