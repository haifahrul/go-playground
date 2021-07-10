package modules

import (
	_ "github.com/haifahrul/go-playground/internal/modules/example"
	"github.com/haifahrul/go-playground/internal/router"
)

func init() {
	router.MiddlewareInit()
}
