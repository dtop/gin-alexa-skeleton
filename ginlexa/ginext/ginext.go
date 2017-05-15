package ginext

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zalando/gin-glog"
	"net/http"
)

/*****************************************************
            Dependency Injection Container
******************************************************/

// Di is the actual dependency container object
type Di struct {
	obj  map[string]interface{}
	lazy map[string]func() interface{}
}

// NewDi creates a new dependency container object
func NewDi() *Di {

	return &Di{
		obj:  make(map[string]interface{}),
		lazy: make(map[string]func() interface{}),
	}
}

// Register registers a dependency for further use
func (di Di) Register(key string, obj interface{}) {
	di.obj[key] = obj
}

// RegisterLazy allows to register a func which will be executen on first use and made available then
func (di Di) RegisterLazy(key string, cb func() interface{}) {
	di.lazy[key] = cb
}

// Get returns the previously registered dependencies
func (di Di) Get(key string) interface{} {

	if val, ok := di.lazy[key]; ok {
		x := val()

		if x == nil {
			panic("lazy service was nil")
			return nil
		}

		di.Register(key, x)
		delete(di.lazy, key)
	}

	if val, ok := di.obj[key]; ok {
		return val
	}

	return nil
}

// GetObj assignes the dependency to a given object
func (di Di) GetObj(key string, obj interface{}) {

	val := di.Get(key)
	obj = val
}

// Attach sets the di container to the gin context to make it available in endpoint handlerfuncs
func (di *Di) Attach() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Set("di", di)
		c.Next()
	}
}

/*****************************************************
                     App Container
******************************************************/

// App struct is a wrapper object around the actual gin obj
type App struct {
	gin *gin.Engine
	di  *Di
}

// NewApp creates a new app obj
func NewApp(di *Di) *App {
	return &App{di: di}
}

// Di returns the dependency injection container
func (a *App) Di() *Di {
	return a.di
}

// Bootstrap prepares the gin app for running
func (a *App) Bootstrap(setupRoutes func(router *gin.Engine)) {

	a.gin = gin.New()
	a.gin.Use(a.di.Attach())
	a.gin.Use(ginglog.Logger(120))
	a.gin.Use(gin.Recovery())

	setupRoutes(a.gin)
}

// Run starts the actual server
func (a *App) Run(port int) {

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: a.gin,
	}

	server.ListenAndServe()
}

/*****************************************************
                   Var Injector
******************************************************/

// VarInject injects dependencies into the gin context on declaration of routes
func VarInject(name string, obj interface{}) gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Set(name, obj)
		c.Next()
	}
}