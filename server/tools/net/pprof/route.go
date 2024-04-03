package pprof

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func (p *PprofService) Route(route *gin.Engine) {
	pprof.Register(route, "pprof")
}
