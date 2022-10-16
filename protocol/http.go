package protocol

import (
	"calf-restful-api/apps"
	"calf-restful-api/conf"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
	"time"
)

func NewHttpService() *HttpService {
	r := gin.Default()

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HttpAdress(),
		Handler:           r,
	}
	return &HttpService{
		server: server,
		r:      r,
		l:      zap.L().Named("Http Service"),
	}
}

type HttpService struct {
	server *http.Server
	r      gin.IRouter
	l      logger.Logger
}

func (s *HttpService) Start() error {
	// 加载Handler, 把所有的模块的Handler注册给了Gin Router
	apps.InitGin(s.r)

	// 已加载App的日志信息
	apps := apps.LoadGinApps()
	s.l.Infof("loaded gin apps :%v", apps)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service stoped success")
			return nil
		}
		return fmt.Errorf("start gin server error:%s", err.Error())
	}
	return nil
}

func (s *HttpService) Stop() {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		fmt.Println("shutdown http server error:%s", err.Error())
	}
}
