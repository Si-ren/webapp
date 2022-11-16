package impl

import (
	resource "cmdb/pkg/resource/api/v1"
	"context"
)

const (
	sqlInsertResource = `INSERT INTO resource (
		id,resource_type,vendor,region,zone,create_at,expire_at,category,type,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type,describe_hash,resource_hash,secret_id,domain,
		namespace,env,usage_mode
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	// 定义的用于变更Informtion属性
	sqlUpdateResource = `UPDATE resource SET
		expire_at=?,category=?,type=?,name=?,description=?,
		status=?,update_at=?,sync_at=?,sync_accout=?,
		public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?,
		secret_id=?,namespace=?,env=?,usage_mode=?
	WHERE id = ?`
	sqlDeleteResource = `DELETE FROM resource WHERE id = ?;`
	// SELECT r.* FROM resource r LEFT JOIN resource_tag t ON r.id=t.resource_id WHERE t.t_key='xx', t.t_value='xxx';
	sqlQueryResource = `SELECT r.* FROM resource r %s JOIN resource_tag t ON r.id = t.resource_id`
	// 	-- resourceA   t1=v1  t2=v2
	// -- resourceA  t1=v1
	// -- resourceA  t2=v2
	// -- 使用DISTINCT对字段去重
	// -- 用于分页时使用
	sqlCountResource = `SELECT COUNT(DISTINCT r.id) FROM resource r %s JOIN resource_tag t ON r.id = t.resource_id`

	sqlQueryResourceTag  = `SELECT t_key,t_value,description,resource_id,weight,type FROM resource_tag`
	sqlDeleteResourceTag = `
		DELETE 
		FROM
			resource_tag 
		WHERE
			resource_id =? 
			AND t_key =? 
			AND t_value =?;
	`
	sqlInsertOrUpdateResourceTag = `
		INSERT INTO resource_tag ( type, t_key, t_value, description, resource_id, weight, create_at)
		VALUES
			( ?,?,?,?,?,?,? ) 
			ON DUPLICATE KEY UPDATE description =
		IF
			( type != 1,?, description ),
			weight =
		IF
			( type != 1,?, weight );
	`
)

func (s *service) create(ctx context.Context, r resource.Resource) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(r.Base).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(r.Information).Error; err != nil {
		tx.Rollback()
		return err
	}
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	return tx.Commit().Error
}
func (s *service) updateTag(ctx context.Context, r *resource.UpdateTagRequest) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, v := range r.Tags {
		tx.Model(&resource.Tag{}).Updates(resource.Tag{
			ResourceId: v.ResourceId,
			Type:       v.Type,
			Key:        v.Key,
			Value:      v.Value,
			Describe:   v.Describe,
			Weight:     v.Weight,
			IsCost:     v.IsCost,
			Hidden:     v.Hidden,
			Meta:       v.Meta,
		})
	}
	tx.Commit()
	return nil
}
