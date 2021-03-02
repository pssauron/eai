package context

type IComponent interface {
	//GetSysParam 获取初始化参数
	GetSysParam() map[string]interface{}

	Exec(param string) Result
}
