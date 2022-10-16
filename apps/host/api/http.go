package api

import (
	"calf-restful-api/apps"
	"calf-restful-api/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var handler = &Handler{}

type Handler struct {
	//注入service对象进行调用
	svc host.Service
	l   logger.Logger
}

func (h *Handler) Name() string {
	return host.APPNAME
}

func (h *Handler) Config() {
	h.svc = apps.GetServiceImpl(host.APPNAME).(host.Service)
	h.l = zap.L().Named("Host Handler")
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/host", h.createHost)
	r.GET("/host", h.queryHost)
}

func init() {
	apps.RegistryGin(handler)
}
