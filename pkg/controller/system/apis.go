/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/8
*/

package system

import (
	"context"
	"github.com/jinzhu/copier"
	"go-easy-admin/internal/model/system"
	"go-easy-admin/pkg/global"

	reqSystem "go-easy-admin/internal/model/request/system"
)

type SysApis interface {
	Create(req *reqSystem.CreateAPIsReq) error
	Delete(id int) error
	Update(id int, req *reqSystem.UpdateAPIsReq) error
	List(apiGroup string) (error, interface{})
	Get(id int) (error, *system.APIs)
	GetApiGroup() (error, []string)
}
type sysApis struct {
	tips string
	ctx  context.Context
}

func NewSysApis(ctx context.Context) SysApis {
	return &sysApis{ctx: ctx, tips: "路由"}
}

func (sa *sysApis) Create(req *reqSystem.CreateAPIsReq) error {
	api := new(system.APIs)
	if err := copier.Copy(&api, req); err != nil {
		return global.OtherErr(err, "转换路由数据失败")
	}
	api.CreateBy = sa.ctx.Value("username").(string)
	if err := global.GORM.WithContext(sa.ctx).Create(api).Error; err != nil {
		return global.CreateErr(sa.tips, err)
	}
	return nil
}

func (sa *sysApis) Delete(id int) error {
	err, api := sa.Get(id)
	if err != nil {
		return err
	}
	if err = global.GORM.WithContext(sa.ctx).Model(&api).Association("Menus").Clear(); err != nil {
		return global.OtherErr(err, "删除关联菜单失败")
	}
	if err = global.GORM.WithContext(sa.ctx).Delete(&api).Error; err != nil {
		return global.DeleteErr(sa.tips, err)
	}
	return nil
}

func (sa *sysApis) Update(id int, req *reqSystem.UpdateAPIsReq) error {
	api := new(system.APIs)
	if err := copier.Copy(&api, req); err != nil {
		return global.OtherErr(err, "转换路由数据失败")
	}
	api.CreateBy = sa.ctx.Value("username").(string)
	if err := global.GORM.WithContext(sa.ctx).Model(&system.APIs{}).Where("id = ?", id).Updates(api).Error; err != nil {
		return global.UpdateErr(sa.tips, err)
	}
	return nil

}

func (sa *sysApis) List(apiGroup string) (error, interface{}) {
	var resApis []system.APIs
	if err := global.GORM.WithContext(sa.ctx).Model(&system.APIs{}).Where("api_group LIKE ? ", "%"+apiGroup+"%").
		Find(&resApis).Error; err != nil {
		return global.GetErr(sa.tips, err), nil
	}
	return nil, &resApis
}

func (sa *sysApis) Get(id int) (error, *system.APIs) {
	api := new(system.APIs)
	if err := global.GORM.WithContext(sa.ctx).Model(&system.APIs{}).Where("id = ?", id).Preload("Menus").First(&api).Error; err != nil {
		return global.GetErr(sa.tips, err), nil
	}
	return nil, api
}

func (sa *sysApis) GetApiGroup() (error, []string) {
	var apiGroups []string
	if err := global.GORM.WithContext(sa.ctx).Model(&system.APIs{}).Distinct().Pluck("api_group", &apiGroups).Error; err != nil {
		return global.GetErr(sa.tips, err), nil
	}
	return nil, apiGroups
}
