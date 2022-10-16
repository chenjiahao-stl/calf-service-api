package api

import (
	"calf-restful-api/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

// 控制层controller
func (h *Handler) createHost(c *gin.Context) {

	ins := host.NewHost()
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		h.l.Errorf("controller host bind error: %s", err.Error())
		return
	}
	ins, err := h.svc.CreateHost(c, ins)
	if err != nil {
		response.Failed(c.Writer, err)
		h.l.Errorf("CreateHost error: %s", err.Error())
		return
	}

	response.Success(c.Writer, ins)
}

func (h *Handler) queryHost(ctx *gin.Context) {

}
