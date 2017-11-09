package plugin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Plugin *gin.Context

var drivers = map[string]InitFunc{}

type InitFunc func(r gin.IRouter) (Plugin, error)

func Register(name string, f InitFunc) {
	drivers[name] = f
}

func New(name string, r gin.IRouter) (Plugin, error) {
	if d, ok := drivers[name]; !ok {
		return nil, fmt.Errorf("'%s' is not found in bot drivers", name)
	} else {
		return d(r)
	}
}
