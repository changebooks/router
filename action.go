package router

import "github.com/changebooks/kernel"

type Action interface {
	GetName() string                   // action's name，全小写，"_"拼接
	Before(holder *Holder) error       // if err != nil，跳过Run()，执行After()
	Run(holder *Holder) *kernel.Result // 主流程
	After(holder *Holder)              // 清理、释放资源
	AllowMethods() []string            // 允许的请求方式，全部大写，GET、POST、PUT、DELETE、OPTIONS、...中的一种或多种
}
