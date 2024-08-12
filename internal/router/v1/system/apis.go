/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/8
*/

package system

import (
	"github.com/gin-gonic/gin"
	apiSystem "go-easy-admin/api/system"
)

/*
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetApiGroup(ctx *gin.Context)
*/

func Apis(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/create", apiSystem.NewSysApis().Create)
		r.POST("/delete/:id", apiSystem.NewSysApis().Delete)
		r.POST("/update/:id", apiSystem.NewSysApis().Update)
		r.GET("/list", apiSystem.NewSysApis().List)
		r.GET("/get/:id", apiSystem.NewSysApis().Get)
		r.GET("/get/group", apiSystem.NewSysApis().GetApiGroup)
	}
	return r
}
