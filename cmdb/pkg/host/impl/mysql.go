package impl

import (
	"cmdb/conf"
	"cmdb/pkg"
	"cmdb/pkg/host"
	"context"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	insertResourceSQL = `INSERT INTO resource (
		id,vendor,region,zone,create_at,expire_at,category,type,instance_id,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type,describe_hash,resource_hash
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	insertHostSQL = `INSERT INTO host (
		resource_id,cpu,memory,gpu_amount,gpu_spec,os_type,os_name,
		serial_number,image_id,internet_max_bandwidth_out,
		internet_max_bandwidth_in,key_pair_name,security_groups
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`
	updateResourceSQL = `UPDATE resource SET 
		expire_at=?,category=?,type=?,name=?,description=?,
		status=?,update_at=?,sync_at=?,sync_accout=?,
		public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?
	WHERE id = ?`
	updateHostSQL = `UPDATE host SET 
		cpu=?,memory=?,gpu_amount=?,gpu_spec=?,os_type=?,os_name=?,
		image_id=?,internet_max_bandwidth_out=?,
		internet_max_bandwidth_in=?,key_pair_name=?,security_groups=?
	WHERE resource_id = ?`

	queryHostSQL      = `SELECT * FROM resource as r LEFT JOIN host h ON r.id=h.resource_id`
	deleteHostSQL     = `DELETE FROM host WHERE resource_id = ?;`
	deleteResourceSQL = `DELETE FROM resource WHERE id = ?;`
)

var (
	// HostService 服务实例
	HostService              = &service{}
	_           host.Service = (*service)(nil)
)

type service struct {
	db  *gorm.DB
	log *logrus.Logger
}

func (s *service) Name() string {
	return host.AppName
}

func (s *service) Config() error {
	db, err := conf.Configure().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = conf.Log
	s.db = db
	return nil
}

func (s *service) CreateHost(ctx context.Context, h *host.Host) (
	*host.Host, error) {
	h.Base.InstanceId = xid.New().String()
	h.ResourceId = h.Base.InstanceId
	h.DescribeId = h.InstanceId
	if err := s.create(ctx, h); err != nil {
		return nil, err
	}

	return h, nil
}

func (s *service) QueryHost(ctx context.Context, req *host.QueryHostRequest) (
	*host.HostSet, error) {

	return nil, nil
}

func (s *service) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (
	*host.Host, error) {

	return nil, nil
}

func (s *service) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (
	*host.Host, error) {

	return nil, nil
}

func (s *service) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (
	*host.Host, error) {
	ins, err := s.DescribeHost(ctx, host.NewDescribeHostRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	if err := s.delete(ctx, req); err != nil {
		return nil, err
	}

	return ins, nil
}

func init() {
	pkg.ImplRegister(HostService)
}
