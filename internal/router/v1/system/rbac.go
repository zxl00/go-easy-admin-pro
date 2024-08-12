/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/9
*/

package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "go-easy-admin/api/system"
)

func RBAC(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/create", apiSystem.NewSysRBAC().Create)
		r.POST("/delete/:id", apiSystem.NewSysRBAC().Delete)
		r.POST("/update/:id", apiSystem.NewSysRBAC().Update)
		r.GET("/list", apiSystem.NewSysRBAC().List)
	}
	return r
}
