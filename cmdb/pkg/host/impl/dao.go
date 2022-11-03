package impl

import (
	"cmdb/pkg/host"
	"context"
)

func (s *service) create(ctx context.Context, h *host.Host) error {
	//var (
	//	stmt *sql.Stmt
	//	err  error
	//)
	// 开启一个事物
	// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
	// 关于事物级别可以参考文章: https://zhuanlan.zhihu.com/p/117476959
	// wiki: https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels
	//tx, err := s.db.BeginTx(ctx, nil)
	//if err != nil {
	//	return err
	//}

	// 执行结果提交或者回滚事务
	// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
	//defer func() {
	//	if err != nil {
	//		tx.Rollback()
	//		return
	//	}
	//}()

	// 避免SQL注入, 请使用Prepare
	//stmt, err = tx.Prepare(insertResourceSQL)
	//if err != nil {
	//	return err
	//}
	//defer stmt.Close()
	//
	//// 生成描写信息的Hash
	//if err := h.GenHash(); err != nil {
	//	return err
	//}
	//
	//// vendor  h.Version.String()
	//_, err = stmt.Exec(
	//	h.Id, h.Vendor, h.Region, h.Zone, h.CreateAt, h.ExpireAt, h.Category, h.Type, h.InstanceId,
	//	h.Name, h.Description, h.Status, h.UpdateAt, h.SyncAt, h.SyncAccount, h.PublicIP,
	//	h.PrivateIP, h.PayType, h.DescribeHash, h.ResourceHash,
	//)
	//if err != nil {
	//	return err
	//}
	//
	//// 避免SQL注入, 请使用Prepare
	//stmt, err = tx.Prepare(insertHostSQL)
	//if err != nil {
	//	return err
	//}
	//defer stmt.Close()
	//
	//_, err = stmt.Exec(
	//	h.ResourceId, h.CPU, h.Memory, h.GPUAmount, h.GPUSpec, h.OSType, h.OSName,
	//	h.SerialNumber, h.ImageID, h.InternetMaxBandwidthOut,
	//	h.InternetMaxBandwidthIn, h.KeyPairName, h.SecurityGroups,
	//)
	//if err != nil {
	//	return err
	//}
	//
	//return tx.Commit()

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(h.Resource).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(h.Base).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(h.Describe).Error; err != nil {
		tx.Rollback()
		return err
	}
	if tx.Error != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}

func (s *service) delete(ctx context.Context, req *host.DeleteHostRequest) error {
	//var (
	//	stmt *sql.Stmt
	//	err  error
	//)
	//
	//// 开启一个事物
	//// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
	//// 关于事物级别可以参考文章: https://zhuanlan.zhihu.com/p/117476959
	//tx, err := s.db.BeginTx(ctx, nil)
	//if err != nil {
	//	return err
	//}
	//
	//// 执行结果提交或者回滚事务
	//// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	//// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
	//defer func() {
	//	if err != nil {
	//		tx.Rollback()
	//		return
	//	}
	//}()
	//
	//stmt, err = tx.Prepare(deleteHostSQL)
	//if err != nil {
	//	return err
	//}
	//defer stmt.Close()
	//
	//_, err = stmt.Exec(req.Id)
	//if err != nil {
	//	return err
	//}
	//
	//stmt, err = s.db.Prepare(deleteResourceSQL)
	//if err != nil {
	//	return err
	//}
	//defer stmt.Close()
	//
	//_, err = stmt.Exec(req.Id)
	//if err != nil {
	//	return err
	//}
	//
	//return tx.Commit()
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("id = ?", req.Id).Delete(&host.Resource{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (s *service) query(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	//var (
	//	stmt *sql.Stmt
	//	err  error
	//)
	//
	//// 开启一个事物
	//// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
	//// 关于事物级别可以参考文章: https://zhuanlan.zhihu.com/p/117476959
	//tx, err := s.db.BeginTx(ctx, nil)
	//if err != nil {
	//	return err
	//}
	//
	//// 执行结果提交或者回滚事务
	//// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	//// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
	//defer func() {
	//	if err != nil {
	//		tx.Rollback()
	//	}
	//}()
	//
	//stmt, err = tx.Prepare(queryHostSQL)
	//if err != nil {
	//	return err
	//}
	//defer stmt.Close()
	//
	//if req.Keywords != "" {
	//	stmt.Query()
	//}
	//res, err = stmt.Query()
	//if err != nil {
	//	return err
	//}
	//
	//return tx.Commit()

	// if err := s.db.Where("id>? limit ?", req.PageSize*req.PageNumber, req.PageSize).Find(&host.Base{}).Error; err != nil {
	// 	return nil, err
	// }

	hSet := host.NewHostSet()
	hSet.Items = make([]*host.Host, 0)
	sliceNum := 0
	rowsBase, err := s.db.Raw(queryBaseSQL, req.PageNumber*req.PageSize, req.PageSize).Rows()
	defer rowsBase.Close()
	if err != nil {
		return nil, err
	}
	for rowsBase.Next() {
		hb := host.NewHost()
		rowsBase.Scan(
			&hb.BaseId,
			&hb.InstanceId,
			&hb.SyncAt,
			&hb.Vendor,
			&hb.Region,
			&hb.Zone,
			&hb.CreateAt,
			&hb.ResourceHash,
			&hb.DescribeHash)

		hSet.Add(hb)
		sliceNum++
	}
	sliceNum = 0
	rowsDescribe, err := s.db.Raw(queryDescribeSQL, req.PageNumber*req.PageSize, req.PageSize).Rows()
	defer rowsDescribe.Close()
	if err != nil {
		return nil, err
	}
	for rowsDescribe.Next() {
		rowsBase.Scan(hSet.Items[sliceNum].Describe.DescribeId,
			&hSet.Items[sliceNum].Describe.CPU,
			&hSet.Items[sliceNum].Describe.Memory,
			&hSet.Items[sliceNum].Describe.GPUAmount,
			&hSet.Items[sliceNum].Describe.GPUSpec,
			&hSet.Items[sliceNum].Describe.OSType,
			&hSet.Items[sliceNum].Describe.OSName,
			&hSet.Items[sliceNum].Describe.SerialNumber,
			&hSet.Items[sliceNum].Describe.ImageID,
			&hSet.Items[sliceNum].Describe.InternetMaxBandwidthIn,
			&hSet.Items[sliceNum].InternetMaxBandwidthOut,
			&hSet.Items[sliceNum].KeyPairName,
			&hSet.Items[sliceNum].SecurityGroups)
		sliceNum++
	}
	sliceNum = 0
	rowsResource, err := s.db.Raw(queryResourceSQL, req.PageNumber*req.PageSize, req.PageSize).Rows()
	defer rowsResource.Close()
	if err != nil {
		return nil, err
	}
	for rowsResource.Next() {
		rowsBase.Scan(hSet.Items[sliceNum].Resource.ResourceId,
			&hSet.Items[sliceNum].Resource.ExpireAt,
			&hSet.Items[sliceNum].Resource.Category,
			&hSet.Items[sliceNum].Resource.Type,
			&hSet.Items[sliceNum].Resource.Name,
			&hSet.Items[sliceNum].Resource.Description,
			&hSet.Items[sliceNum].Resource.Status,
			&hSet.Items[sliceNum].Resource.Tags,
			&hSet.Items[sliceNum].Resource.UpdateAt,
			&hSet.Items[sliceNum].Resource.SyncAccount,
			&hSet.Items[sliceNum].Resource.PublicIP,
			&hSet.Items[sliceNum].Resource.PrivateIP,
			&hSet.Items[sliceNum].Resource.PayType)
		sliceNum++
	}
	hSet.Total = sliceNum
	return hSet, nil
}
