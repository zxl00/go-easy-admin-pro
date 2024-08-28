/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/9
*/

package system

import (
	"context"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go-easy-admin/pkg/global"
)

// casbin 权限

type SysRBAC interface {
	Create(rules [][]string) error
	Delete(id int) bool
	Update(id int, rule []string) error
	List(v0, v1, v2, v3 string) (error, interface{})
}
type sysRbac struct {
	tips string
	ctx  context.Context
}

func NewSysRBAC(ctx context.Context) SysRBAC {
	return &sysRbac{ctx: ctx, tips: "权限"}
}

// 添加权限

func (sr *sysRbac) Create(rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
			V3:    rules[i][3],
			V4:    sr.ctx.Value("username").(string),
		})
	}
	// 这里需要写数据库
	if err := global.GORM.WithContext(sr.ctx).Create(&casbinRules).Error; err != nil {
		return err
	}
	freshRBAC()
	return nil
}

func (sr *sysRbac) Delete(id int) bool {
	// 根据ID查找权限
	var casbinRule gormadapter.CasbinRule
	global.GORM.Model(&casbinRule).WithContext(sr.ctx).Where("id = ?", id).First(&casbinRule)
	success, _ := global.CasbinCacheEnforcer.RemovePolicy(casbinRule.V0, casbinRule.V1, casbinRule.V2, casbinRule.V3, casbinRule.V4)
	return success
}

func (sr *sysRbac) Update(id int, rule []string) error {
	var casbinRule = gormadapter.CasbinRule{
		Ptype: "p",
		V0:    rule[0],
		V1:    rule[1],
		V2:    rule[2],
		V3:    rule[3],
	}
	if err := global.GORM.Model(&casbinRule).WithContext(sr.ctx).Where("id = ?", id).Updates(casbinRule).Error; err != nil {
		return err
	}
	freshRBAC()
	return nil
}

func (sr *sysRbac) List(v0, v1, v2, v3 string) (error, interface{}) {
	var resCasbinRules []gormadapter.CasbinRule
	if err := global.GORM.Model(&gormadapter.CasbinRule{}).WithContext(sr.ctx).Where("v0 LIKE ? and v1 LIKE ? and v2 LIKE ? and v3 LIKE ?",
		"%"+v0+"%", "%"+v1+"%", "%"+v2+"%", "%"+v3+"%").Find(&resCasbinRules).Error; err != nil {
		return err, nil
	}
	return nil, &resCasbinRules
}

func freshRBAC() {
	_ = global.CasbinCacheEnforcer.LoadPolicy()
}
