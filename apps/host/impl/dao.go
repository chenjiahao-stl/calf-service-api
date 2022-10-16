package impl

import (
	"calf-restful-api/apps/host"
	"context"
	"fmt"
)

// 把Host对象保存到数据内, 数据的一致性   context.Context 上下文信息->防止取消的行为出现能够及时感知到
func (i *IHostDao) save(ctx context.Context, ins *host.Host) error {

	var err error

	tx, err := i.db.BeginTx(ctx, nil) //开启事务
	if err != nil {
		return fmt.Errorf("start tx error :%s", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Errorf("tx rollback error:%s", err.Error())
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Errorf("tx Commit error:%s", err.Error())
			}
		}
	}()

	// 插入Resource数据
	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()

	_, err = rstmt.ExecContext(ctx,
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return err
	}
	// 插入Describe 数据
	dstmt, err := tx.PrepareContext(ctx, InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()

	_, err = dstmt.ExecContext(ctx,
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return err
	}

	return nil
}

func (i *IHostDao) update(ctx context.Context, ins *host.Host) error {

	return nil
}

func (i *IHostDao) query(ctx context.Context, request *host.QueryHostRequest) (*host.HostSet, error) {

	return nil, nil
}
