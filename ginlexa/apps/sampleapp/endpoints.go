package sampleapp

import (
    echoRequest "github.com/go-alexa/alexa/parser"
    echoResponse "github.com/go-alexa/alexa/response"
    "github.com/dtop/gin-alexa"
)

// OnAuth handles any kind of user authentication
func OnAuth (*ginalexa.EchoContext, *echoRequest.Event, *echoResponse.Response) {

}

// OnLaunch should handle the amazon echo on launch request
func OnLaunch (*ginalexa.EchoContext, *echoRequest.Event, *echoResponse.Response) {

}

// OnIntent should handle all intent requests
func OnIntent (*ginalexa.EchoContext, *echoRequest.Event, *echoResponse.Response) {

}

// OnSessionEnded should hndle the on session ended request
func OnSessionEnded (*ginalexa.EchoContext, *echoRequest.Event, *echoResponse.Response) {

}
