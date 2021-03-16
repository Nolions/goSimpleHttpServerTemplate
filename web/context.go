package web

import (
	"simpleHttpServer/store"
)

type Context struct {
	// DB操作對象
	store store.Context
}

var context *Context

func Setup(db store.Context) {
	context = &Context{
		store: db,
	}
}

// 新增會員
func CreateMember(name string) error {
	return context.store.CreateMember(&store.Member{
		Name: name,
	})
}
