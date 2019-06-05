package controller

import (
	"context"
)

type ControllerInterface interface {
	Init(ct *context.Context, cn string) // 初始化上下文和子类名称
	Prepare()                            // 开始执行之前的一些处理

	Finish() // 执行完成之后的处理

	RouteMaps()
}
