package models

//后端与页面交互的数据模型
type ObjLoginuser struct {
	Uid      int
	Username string
	Now      int
	Ip       string
	Sign     string
}
