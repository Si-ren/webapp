package host

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"
)

const (
	PrivateIDC Vendor = iota
	Tencent
	AliYun
	HuaWei
)

type Vendor int

func NewHost() *Host {
	return &Host{
		&Base{},
		&Resource{},
		&Describe{},
	}
}

type Host struct {
	*Base
	*Resource
	*Describe
}

func (h *Host) Put(req *UpdateHostData) {
	h.Resource = req.Resource
	h.Describe = req.Describe
	h.Resource.UpdateAt = time.Now() // time, 13 时间戳
	h.GenHash()
}

func (h *Host) Patch(req *UpdateHostData) error {
	err := ObjectPatch(h.Resource, req.Resource)
	if err != nil {
		return err
	}

	err = ObjectPatch(h.Describe, req.Describe)
	if err != nil {
		return err
	}

	h.Resource.UpdateAt = time.Now()
	h.GenHash()
	return nil
}

// patch JSON {a: 1, b： 2}， {b:20}  ===> {a:1, b:20}
func ObjectPatch(old, new interface{}) error {
	// {b: 20}
	newByte, err := json.Marshal(new)
	if err != nil {
		return err
	}
	// {a:1, b:2}
	// {a:1, b: 20}
	return json.Unmarshal(newByte, old)
}

func (h *Host) GenHash() error {
	hash := sha1.New()

	b, err := json.Marshal(h.Resource)
	if err != nil {
		return err
	}
	hash.Write(b)
	h.Base.ResourceHash = fmt.Sprintf("%x", hash.Sum(nil))

	b, err = json.Marshal(h.Describe)
	if err != nil {
		return err
	}
	hash.Reset()
	hash.Write(b)
	h.Base.DescribeHash = fmt.Sprintf("%x", hash.Sum(nil))
	return nil
}

type Base struct {
	BaseId       int       `json:"base_id" gorm:"primaryKey;comment:BaseId;not null"`               // BaseId
	InstanceId   string    `json:"instance_id" gorm:"comment:实例ID;not null"`                        // 实例ID
	SyncAt       int       `json:"sync_at,omitempty" gorm:"autoCreateTime;comment:同步时间"`            // 同步时间
	Vendor       Vendor    `json:"vendor" gorm:"comment:厂商;not null"`                               // 厂商
	Region       string    `json:"region" gorm:"comment:地域;not null"`                               // 地域
	Zone         string    `json:"zone" gorm:"comment:区域;not null"`                                 // 区域
	CreateAt     time.Time `json:"create_at,omitempty" gorm:"autoCreateTime;comment:创建时间;not null"` // 创建时间
	ResourceHash string    `json:"resource_hash,omitempty" gorm:"comment:基础数据Hash;not null"`        // 基础数据Hash
	DescribeHash string    `json:"describe_hash,omitempty" gorm:"comment:描述数据Hash;not null"`        // 描述数据Hash
}

// TableName 会将 User 的表名重写为 `profiles`
func (Base) TableName() string {
	return "base"
}

type Resource struct {
	ResourceId  string    `json:"resource_id,omitempty" gorm:"primaryKey;comment:ResourceId;not null"` // ResourceId
	ExpireAt    time.Time `json:"expire_at,omitempty" gorm:"comment:过期时间"`                             // 过期时间
	Category    string    `json:"category,omitempty" gorm:"comment:种类;not null"`                       // 种类
	Type        string    `json:"type,omitempty" gorm:"comment:规格"`                                    // 规格
	Name        string    `json:"name,omitempty" gorm:"comment:名称"`                                    // 名称
	Description string    `json:"description,omitempty" gorm:"comment:描述"`                             // 描述
	Status      string    `json:"status,omitempty" gorm:"comment:服务商中的状态;not null"`                    // 服务商中的状态
	Tags        string    `json:"tags,omitempty" gorm:"comment:标签"`                                    // 标签
	UpdateAt    time.Time `json:"update_at,omitempty" gorm:"autoUpdateTime;comment:更新时间"`              // 更新时间
	SyncAccount string    `json:"sync_accout,omitempty" gorm:"comment:同步的账号"`                          // 同步的账号
	PublicIP    string    `json:"public_ip,omitempty" gorm:"comment:公网IP"`                             // 公网IP
	PrivateIP   string    `json:"private_ip,omitempty" gorm:"comment:内网IP"`                            // 内网IP
	PayType     string    `json:"pay_type,omitempty" gorm:"comment:实例付费方式;not null"`                   // 实例付费方式
}

// TableName 会将 User 的表名重写为 `profiles`
func (Resource) TableName() string {
	return "resource"
}

type Describe struct {
	DescribeId              string `json:"describe_id,omitempty" gorm:"comment:关联DescribeId;size:255"`             // 关联DescribeId
	CPU                     int    `json:"cpu,omitempty" gorm:"comment:核数;not null;type:tinyint;size:2"`           // 核数
	Memory                  int    `json:"memory,omitempty" gorm:"comment:内存;not null"`                            // 内存
	GPUAmount               int    `json:"gpu_amount,omitempty" gorm:"comment:GPU数量;type:tinyint"`                 // GPU数量
	GPUSpec                 string `json:"gpu_spec,omitempty" gorm:"comment:GPU类型"`                                // GPU类型
	OSType                  string `json:"os_type,omitempty" gorm:"comment:操作系统类型，分为Windows和Linux;size:255"`       // 操作系统类型，分为Windows和Linux
	OSName                  string `json:"os_name,omitempty" gorm:"comment:操作系统名称"`                                // 操作系统名称
	SerialNumber            string `json:"serial_number,omitempty" gorm:"comment:序列号"`                             // 序列号
	ImageID                 string `json:"image_id,omitempty" gorm:"comment:镜像ID"`                                 // 镜像ID
	InternetMaxBandwidthOut int    `json:"internet_max_bandwidth_out,omitempty" gorm:"comment:公网出带宽最大值，单位为 Mbps;"` // 公网出带宽最大值，单位为 Mbps
	InternetMaxBandwidthIn  int    `json:"internet_max_bandwidth_in,omitempty" gorm:"comment:公网入带宽最大值，单位为 Mbps"`   // 公网入带宽最大值，单位为 Mbps
	KeyPairName             string `json:"key_pair_name,omitempty" gorm:"comment:秘钥对名称"`                           // 秘钥对名称
	SecurityGroups          string `json:"security_groups,omitempty" gorm:"comment:安全组 采用逗号分隔"`                    // 安全组  采用逗号分隔
}

// TableName 会将 User 的表名重写为 `profiles`
func (Describe) TableName() string {
	return "describe"
}
func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}

type HostSet struct {
	Items []*Host `json:"items"`
	Total int     `json:"total"`
}

func (s *HostSet) Add(item *Host) {
	s.Items = append(s.Items, item)
}
