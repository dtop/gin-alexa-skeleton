package main

import "github.com/dtop/gin-alexa-skeleton/ginlexa"

func main() {

    gl := ginlexa.GinLexa{}
    gl.Prepare()
    gl.Run()
}
