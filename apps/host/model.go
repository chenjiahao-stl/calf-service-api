package host

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var (
	validate = validator.New()
)

type HostSet struct {
	total int
	hosts []*Host
}

func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

func (h *Host) Validate() error {
	return validate.Struct(h)
}

func (h *Host) InjectDefault() {
	if h.CreateAt == 0 {
		h.CreateAt = time.Now().UnixMilli()
	}
}

type Host struct {
	// 资源公共属性部分
	*Resource
	// 资源独有属性部分
	*Describe
}

type Vendor int

const (
	// 枚举的默认值
	PRIVATE_IDC Vendor = iota
	// 阿里云
	ALIYUN
	// 腾讯云
	TXYUN
)

type Resource struct {
	Id          string
	Vendor      Vendor
	Region      string
	CreateAt    int64
	ExpireAt    int64
	Type        string
	Name        string
	Description string
	Status      string
	Tags        map[string]string
	UpdateAt    int64
	SyncAt      int64
	Account     string
	PublicIP    string
	PrivateIP   string
}

type Describe struct {
	CPU          int    // 核数
	Memory       int    // 内存
	GPUAmount    int    // GPU数量
	GPUSpec      string // GPU类型
	OSType       string // 操作系统类型，分为Windows和Linux
	OSName       string // 操作系统名称
	SerialNumber string // 序列号
}

type QueryHostRequest struct {
}
