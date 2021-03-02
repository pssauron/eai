package context

import (
	"github.com/pssauron/eai/conf"
	"github.com/pssauron/eai/router"
	"github.com/pssauron/eai/task"
)

//EnvContext 集成环境单例
var EnvContext *EAIContext

type EAIContext struct {
	//ScriptPath 配置文件路径
	ScriptPath string
	//MidDB 中间库地址
	MidDB *conf.DBConfig

	//Components 组件
	Components map[string]*IComponent

	//Routers 路由
	Routers []*router.EAIRouter

	Tasks []*task.EAIPlan
}
