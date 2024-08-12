/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/6
*/

package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "go-easy-admin/api/system"
)

func Ldap(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/create", apiSystem.NewSysLdap().Create)
		r.GET("/list", apiSystem.NewSysLdap().List)
		r.POST("/ping", apiSystem.NewSysLdap().Ping)
	}
	return r
}
