package main

import (
	"fmt"
	"meeting_scheduler_api/route"
	"net/http"
	"os"
)

func main() {
	os.Setenv("PORT","8080")
	port:=os.Getenv("PORT")
	router:=route.GetRouter()
	fmt.Println("PORT : "+port)
	err:=http.ListenAndServe(":"+port,router)
	if err!=nil {
		fmt.Printf("error while listening: %v",err)
	}
}