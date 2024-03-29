package handler

import (
	"github.com/andeya/faygo"
)

type Body struct {
	Json map[string]interface{} `param:"<in:body>"`
}

func (b *Body) Serve(ctx *faygo.Context) error {
	return ctx.JSON(200, b.Json, true)
}
