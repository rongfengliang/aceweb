package main

import (
	"github.com/plimble/ace"
	"github.com/rongfengliang/aceweb/core"
)

func main() {
	a := ace.New()
	a.GET("/:name", func(c *ace.C) {
		name := c.Params.ByName("name")
		responseinfo := core.Response{
			StateCode: 200,
			Message:   "ok",
			Data:      map[string]interface{}{"hello": name, "age": 3333},
		}
		c.JSON(200, responseinfo)
	})
	a.Run(":8080")
}
