package main

import (
	
	"database"
	"fmt"
	"os"
	"routers"
	"github.com/gin-gonic/gin"
	"net/http"
	"local_modules/configloader"
	"time"
	
	
)



func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--runserver" {

		
			
			config := configloader.New("config.yaml")
			gin.SetMode(gin.ReleaseMode)
			routerInit := routers.InitRouter()
			
			

			endPoint := fmt.Sprintf(":%d",config.ServerSet.Httpport)
			server := &http.Server{
				Addr:	endPoint,
				Handler:routerInit,
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}
		
			server.ListenAndServe()
		} else if os.Args[1] == "--dbinit" {
			dbinit := database.New()
			dbinit.InitDatabase()
			dbinit.CloseDBConnect()
		} else {
			fmt.Println("Please use:")
			fmt.Println("--runserver :run gin server.")
			fmt.Println("--dbinit : run mysql init.")
			os.Exit(3)
		}
	} else {
		fmt.Println("Please use:")
		fmt.Println("--runserver :run gin server.")
		fmt.Println("--dbinit : run mysql init.")
		os.Exit(3)
	}

}
