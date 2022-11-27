package service

import (
	"fmt"
	"golang/model"

	"github.com/kataras/iris"
)

type Service interface {
	Create(ctx iris.Context)
}

func Create(ctx iris.Context) {
	var b model.Book
	ctx.ReadJSON(&b)
	fmt.Print(b)
	// m = append(m, b)
	ctx.JSON("successfull insertion")
}
