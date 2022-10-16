package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	implApps = map[string]ImplService{} //service注册
	ginApps  = map[string]GinService{}  //controller注册
)

func RegistryImpl(svc ImplService) {
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}
	implApps[svc.Name()] = svc
}

func GetServiceImpl(appName string) ImplService {
	if service, ok := implApps[appName]; ok {
		return service
	}
	return nil
}

func InitImpl() {
	for _, v := range implApps {
		v.Config()
	}
}

func RegistryGin(gin GinService) {
	if _, ok := ginApps[gin.Name()]; ok {
		panic(fmt.Sprintf("gin %s has registried", gin.Name()))
	}
	ginApps[gin.Name()] = gin
}

func LoadGinApps() (names []string) {
	for _, ginService := range ginApps {
		names = append(names, ginService.Name())
	}
	return names
}

func InitGin(r gin.IRouter) {
	// 先初始化好所有对象
	for _, v := range ginApps {
		v.Config()
	}

	// 完成Http Handler的注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

type ImplService interface {
	Config()
	Name() string
}

type GinService interface {
	Name() string
	Config()
	Registry(r gin.IRouter)
}
