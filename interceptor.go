package router

type Interceptor interface {
	Before(holder *Holder) error // if err != nil，跳过Before()和After()之间代码，直接执行After()
	After(holder *Holder)        // 清理、释放资源
}
