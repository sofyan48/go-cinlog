package main

import (
	"fmt"

	"github.com/sofyan48/go-cinlog/cinlog"
	"github.com/sofyan48/go-cinlog/config"
)

func main() {
	cfg := config.NewConfig().SetClient("localhost:3000", "payment")
	log := cinlog.NewSession(cfg).V1()
	data := map[string]interface{}{
		"code": "200",
		"result": map[string]interface{}{
			"bank":      "bca",
			"confirmed": "1",
			"user_id":   "00001",
		},
	}
	result, _ := log.SaveLogger("12345", "success", data)
	fmt.Println(result)
}
