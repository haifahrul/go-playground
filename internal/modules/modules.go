package modules

import (
	_ "github.com/haifahrul/go-playground/internal/modules/test"
	"github.com/haifahrul/go-playground/internal/router"
)

func init() {
	router.MiddlewareInit()
}
