package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tiagohora/test-app/src/infra/dummy"
	"github.com/tiagohora/test-app/src/subject/xrate/xratedelivery"
)

var (
	router = gin.New()
	routerFunctions = map[string]func(r string, h ...gin.HandlerFunc) gin.IRoutes{
		"GET": router.GET,
	}
)

type HFunc func(c context.Context, r dummy.ResponseWriter)

func main()  {

	mapRoute("GET","/ping", xratedelivery.Factory().Handle)

	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Server error %s", err)
	}
}

// TODO: must be able to receive n middlewares
func mapRoute(verb string, route string, handle HFunc) {
	routerFunctions[verb](route, func (c *gin.Context) {
		handle(context.Context(c), c.Writer)
	})
}


