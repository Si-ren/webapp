package forms

type JobCreateForm struct {
	Key    string `form:"key"`
	Remark string `form:"remark"`
	Node   int    `form:"node"`
}

type JobModifyForm struct {
	ID     int    `form:"id"`
	Key    string `form:"key"`
	Remark string `form:"remark"`
	Node   int    `form:"node"`
}

type TargetModifyForm struct {
	ID     int    `form:"id"`
	Name   string `form:"name"`
	Remark string `form:"remark"`
	Job    int    `form:"job"`
}
type TargetCreateForm struct {
	ID     int    `form:"id"`
	Name   string `form:"name"`
	Remark string `form:"remark"`
	Job    int    `form:"job"`
}

type NodeRegisterForm struct {
	UUID     string `json:"uuid"`
	Hostname string `json:"hostname"`
	Addr     string `json:"addr"`
}
