package ginlexa

import (
	"github.com/dtop/gin-alexa"
	"github.com/dtop/gin-alexa-skeleton/ginlexa/abstract"
	"github.com/dtop/gin-alexa-skeleton/ginlexa/apps/sampleapp"
	"github.com/dtop/gin-alexa-skeleton/ginlexa/config"
	"github.com/dtop/gin-alexa-skeleton/ginlexa/ginext"
	"github.com/dtop/gin-alexa-skeleton/ginlexa/routes"
)

// Config holds the actual config. it should follow the Configurable interface
type Config struct {
	ServerPort int                              `env:"lex.port" json:"lex.port" yaml:"lex.port"`
	Apps       map[string]ginalexa.Configurable `env:"lex.app" envSeparator:":" json:"apps" yaml:"apps"`
}

// GinLexa is the actual gin app to run
type GinLexa struct {
	app  *ginext.App
	cfg  *Config
	apps map[string]*ginalexa.EchoApplication
}

// Prepare prepares the app fro running
func (gl GinLexa) Prepare() {

	gl.apps = make(map[string]*ginalexa.EchoApplication)

	gl.cfg = &Config{}
	config.New(gl.cfg)

	app := ginext.NewApp(ginext.NewDi())

	// register config
	app.Di().Register(abstract.DiNameConfig, gl.cfg)

	// setup apps
	gl.setupApps()

	// bootstrap the app and emit the routes
	app.Bootstrap(routes.New(gl.apps))
	gl.app = app
}

func (gl GinLexa) setupApps() {

	gl.apps["SampleApp"] = &ginalexa.EchoApplication{
		AppID:          "<<<SomeAppId>>>", // can be ommited if config is present
		Config:         gl.cfg.Apps["SampleApp"],
		OnAuthCheck:    sampleapp.OnAuth,
		OnLaunch:       sampleapp.OnLaunch,
		OnIntent:       sampleapp.OnIntent,
		OnSessionEnded: sampleapp.OnSessionEnded,
	}.Inject()
}

// Run runs the actual server
func (gl GinLexa) Run() {
	gl.app.Run(gl.cfg.ServerPort)
}
