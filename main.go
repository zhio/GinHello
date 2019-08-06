package main

import (
	"GinHello/initRouter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := initRouter.SetupRouter()
	_ = router.Run()
}