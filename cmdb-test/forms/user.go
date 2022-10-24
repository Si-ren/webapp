package forms

//要和html中提交的数据的key保持一致,html提交上来的key都为小写
type UserModifyForm struct {
	ID       int    `form:"id"`
	Name     string `form:"name"`
	Password string `form:"password"`
}
