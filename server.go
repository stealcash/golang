package main

import (
	"golang/model"
	"golang/service"

	"github.com/kataras/iris/v12"
)

var m []model.Book

func main() {
	app := iris.New()

	booksAPI := app.Party("/books")
	// GET: http://localhost:8080/books
	booksAPI.Get("/", list)
	// POST: http://localhost:8080/books
	booksAPI.Post("/", create)

	app.Listen(":8080")
}

func list(ctx iris.Context) {

	ctx.JSON(m)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

func create(ctx iris.Context) {
	service.Create(ctx)
}
