/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/6
*/

package system

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/go-ldap/ldap/v3"
	"go-easy-admin/internal/model/system"
	"go-easy-admin/pkg/global"
	"gorm.io/gorm"
)

type SysLdap interface {
	Create(req *system.Ldap) error
	List() (error, *system.Ldap)
	Get() (error, *system.Ldap)
	Ping(req *system.Ldap) error
}
type sysLdap struct {
	tips string
	ctx  context.Context
}

func NewSysLdap(ctx context.Context) SysLdap {
	return &sysLdap{ctx: ctx, tips: "Ldap"}
}

// 创建或更新

func (sl *sysLdap) Create(req *system.Ldap) error {
	// 先判断是否存在 ，且只能存在一条，多条 以第一条为准
	ldap := new(system.Ldap)
	if err := global.GORM.WithContext(sl.ctx).First(&ldap).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建
			req.CreateBy = sl.ctx.Value("username").(string)
			if err = global.GORM.WithContext(sl.ctx).Create(&req).Error; err != nil {
				return global.CreateErr(sl.tips, err)
			}
			return nil
		}
		// 更新
	}
	if err := global.GORM.WithContext(sl.ctx).Model(&ldap).Updates(req).Error; err != nil {
		return global.UpdateErr(sl.tips, err)
	}
	return nil
}

// 获取

func (sl *sysLdap) List() (error, *system.Ldap) {
	ldap := new(system.Ldap)
	if err := global.GORM.WithContext(sl.ctx).First(&ldap).Error; err != nil {
		return global.GetErr(sl.tips, err), nil
	}
	ldap.Password = ""
	return nil, ldap
}

func (sl *sysLdap) Get() (error, *system.Ldap) {
	ldap := new(system.Ldap)
	if err := global.GORM.WithContext(sl.ctx).Where("status = ?", 1).First(&ldap).Error; err != nil {
		return global.GetErr(sl.tips, err), nil
	}
	return nil, ldap
}

func (sl *sysLdap) Ping(req *system.Ldap) error {
	var (
		ld  *ldap.Conn
		err error
	)
	if req.SSL == 1 {
		ld, err = ldap.DialTLS("tcp", req.Address, &tls.Config{InsecureSkipVerify: true})
	} else {
		ld, err = ldap.Dial("tcp", req.Address)
	}
	if err != nil {
		return err
	}
	defer ld.Close()
	if ld != nil {
		if err = ld.Bind(req.DN, req.Password); err != nil {
			return global.OtherErr(errors.New("连接失败" + err.Error()))
		}
	}
	return nil
}
