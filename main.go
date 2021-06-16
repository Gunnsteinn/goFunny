package main

import (
	//"io"
	//"log"
	"net/http"
	//"os"
  
  	"github.com/gin-gonic/contrib/static"
  	"github.com/gin-gonic/gin"
)

// Hello World API
func main() {
	// Set the router as the default one shipped with Gin
	  router := gin.Default()

	  // Serve frontend static files
	  router.Use(static.Serve("/", static.LocalFile("./views", true)))

	  // Setup route group for the API
	  api := router.Group("/api")
	  {
	    api.GET("/", func(c *gin.Context) {
	      c.JSON(http.StatusOK, gin.H {
		"message": "pong",
	      })
	    })
	  }
	
	  // Start and run the server
  	router.Run(":3000")
	
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "8000" //localhost
	//}

	//helloHandler := func(w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "Colgate de esta TanKO!")
	//}

	//http.HandleFunc("/TanKO", helloHandler)

	//log.Fatal(http.ListenAndServe(":"+port, nil))
}
