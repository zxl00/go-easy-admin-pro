/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/9
*/

package system

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/pkg/controller/system"
	"go-easy-admin/pkg/global"
	"strconv"
)

type SysRBAC interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
}
type sysRbac struct {
}

func NewSysRBAC() SysRBAC {
	return &sysRbac{}
}

func (sr *sysRbac) Create(ctx *gin.Context) {
	body := new(struct {
		Rules [][]string `json:"rules"`
	})
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysRBAC(ctx).Create(body.Rules); err != nil {
		global.ReturnContext(ctx).Failed("创建失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("创建成功", nil)
}

func (sr *sysRbac) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if ok := system.NewSysRBAC(ctx).Delete(id); !ok {
		global.ReturnContext(ctx).Failed("删除失败", nil)
		return
	}
	global.ReturnContext(ctx).Successful("删除成功", nil)
}

func (sr *sysRbac) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	body := new(struct {
		Rule []string `json:"rule"`
	})
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysRBAC(ctx).Update(id, body.Rule); err != nil {
		global.ReturnContext(ctx).Failed("更新失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("更新成功", nil)
}

func (sr *sysRbac) List(ctx *gin.Context) {
	params := new(struct {
		RoleID string ` form:"role_id"`
		Path   string `form:"path"`
		Method string `form:"method"`
		Desc   string `form:"desc"`
	})
	if err := ctx.ShouldBindQuery(&params); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err, data := system.NewSysRBAC(ctx).List(params.RoleID, params.Path, params.Method, params.Desc); err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
	} else {
		global.ReturnContext(ctx).Successful("查询成功", data)
	}
}
