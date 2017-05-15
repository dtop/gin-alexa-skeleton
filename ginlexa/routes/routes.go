package routes

import (
	"fmt"
	"github.com/dtop/gin-alexa"
	"github.com/dtop/gin-alexa-skeleton/ginlexa/ginext"
	ginlexaoauth "github.com/dtop/gin-alexa-skeleton/ginlexa/oauth"
	"github.com/gin-gonic/gin"
)

// New returns a func that is executed to setup the routes
func New(apps map[string]*ginalexa.EchoApplication) func(r *gin.Engine) {

	return func(r *gin.Engine) {

		rtr := router{r: r}
		rtr.setupOAuth(apps)

		for name, app := range apps {
			rtr.setupApp(name, app)
		}
	}
}

type router struct {
	r *gin.Engine
}

func (r router) setupOAuth(apps map[string]*ginalexa.EchoApplication) {

	oauth := r.r.Group("/oauth")
	{
		oauth.GET("/authorize", ginext.VarInject("apps", apps), ginlexaoauth.Authorize)
		oauth.POST("/token", ginext.VarInject("apps", apps), ginlexaoauth.Token)
	}
}

func (r router) setupApp(name string, app *ginalexa.EchoApplication) {

	var endpoint string
	if endpoint = app.Config.GetString("Endpoint"); endpoint == "" {
		endpoint = name
	}

	r.r.POST(fmt.Sprintf("/v1/%v", endpoint), ginalexa.EchoMiddlewareAutomatic(app))
}
