package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func rootHandler(ctx *gin.Context) {
	if len(ctx.Query("fail")) > 0 {
		ctx.String(http.StatusBadRequest, "Error")
		return
	}

	client_ip := ctx.ClientIP()

	log.Println("Requst from: " + client_ip)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "-"
	}
	output := fmt.Sprintf("HOSTNAME : %s\nCLIENT IP: %s\n", hostname, client_ip)

	ctx.String(http.StatusOK, output)
}
