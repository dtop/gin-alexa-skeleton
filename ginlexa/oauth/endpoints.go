package oauth

import (
    "github.com/gin-gonic/gin"
    "github.com/dtop/gin-alexa"
)

// Authorize handles the actual oauth /authorize request
func Authorize(c *gin.Context) {

    apps := getApps(c)

}

// Token handles all token related requests (such as refresh, token, etc..)
func Token(c *gin.Context) {

    apps := getApps(c)
}

func getApps(c *gin.Context) (apps map[string]*ginalexa.EchoApplication) {

    xapps, ok := c.Get("apps")

    if !ok {
        return nil
    }

    apps = xapps.(map[string]*ginalexa.EchoApplication)
    return
}