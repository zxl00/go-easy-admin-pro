/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/2
*/

package system

type CreateUserReq struct {
	Username string `json:"userName" binding:"required"`            // 用户登录名
	Password string `json:"password" aes:"true" binding:"required"` // 用户登录密码
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Roles    []int  `json:"roles" binding:"required"`
}

type UpdateUserReq struct {
	Username string `json:"userName"`            // 用户登录名
	Password string `json:"password" aes:"true"` // 用户登录密码
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Roles    []int  `json:"roles"`
}
