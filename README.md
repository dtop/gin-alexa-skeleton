# Alexa Skeleton App
Golang Skeleton Application to start an Alexa Skill

Since Alexa APIs are only accepting one endpoint per skill you can easily
add several apps to one server. Even if you need to handle them separrately you
dont need to implement oauth stuff over and over.

# Information

for building the skeleton I am using the following libs

[gin-gonic/gin](https://github.com/gin-gonic/gin) gin http framework
[dtop/gin-alexa](https://github.com/dtop/gin-alexa) alexa enhancements for gin
[go-alexa/alexa](https://github.com/go-alexa/alexa) actual alexa handling lib
[caarlos0/env](https://github.com/caarlos0/env) env config lib
[go-oauth2/gin-server](https://github.com/go-oauth2/gin-server) gin oauth server


# Usage

since you need to edit the actual files to create your own skill(s)
I would recommend to download the release version of this skeleton
and create a new repo instead of forking it.

Chances that improvements on this skeleton will be mergable are
not very good =(

do not forget to recreate the dependencies after downloading using
[glide](https://github.com/Masterminds/glide)

### 1) Create (or enhance) the config in ginlexa.go

```go

type Config struct {

    
    ServerPort int `....`
    Apps map[string]ginalexa.Configurable `...`
    
    // your stuff
}

```

Advisory: all config into Apps (inside of the map) will be passed to the actual EchoApplication obj

### 2) Setup all applications (1 per skill)

in ginlexa.go

```go

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

```

Advisory: do not forget the .Inject() at end. this will give the appid to the obj if you did not give it to the struct

### 3) Create another directory under apps and start your skill programming

tbc: create wrappers for your data storage stuff along with the provided interfaces in /abstract/types.go

# Hints

tbc