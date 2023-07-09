package main

import (
	"log"
	"net/http"
	"readlearnapp/routes"
	"readlearnapp/utils"
	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/cors"

)

func init(){
	utils.InitObjectBox()

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	defer utils.ObjectBox.Close()
	router := gin.New()
	router.Use(CORSMiddleware())
	router.GET("/addData",routes.AddData)
	router.GET("/getStories",routes.GetStories)
	router.GET("/serveStoryData", routes.ServeStoryData)
	router.GET("/deleteStory", routes.DeleteStory)
	
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
