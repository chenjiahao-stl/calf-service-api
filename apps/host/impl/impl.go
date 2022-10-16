package impl

import (
	"calf-restful-api/apps/host"
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger"
)

func (i *IHostServiceImpl) CreateHost(ctx context.Context, host *host.Host) (*host.Host, error) {
	// 直接打印日志
	i.l.Named("Creaet").Debug("create host")
	i.l.Info("create host")
	// // 带Format的日志打印, fmt.Sprintf()
	i.l.Debugf("create host %s", host.Name)
	// // 携带额外meta数据, 常用于Trace系统
	i.l.With(logger.NewAny("request-id", "req01")).Debug("create host with meta kv")

	// 校验数据合法性
	if err := host.Validate(); err != nil {
		fmt.Errorf("create host validate is error: %s", err.Error())
		return nil, err
	}

	//默认值填充
	host.InjectDefault()

	if err := i.dao.save(ctx, host); err != nil {
		return nil, err
	}
	return host, nil
}

func (i *IHostServiceImpl) QueryHost(ctx context.Context, request *host.QueryHostRequest) (*host.HostSet, error) {

	i.dao.query(ctx, request)
	return nil, nil
}
