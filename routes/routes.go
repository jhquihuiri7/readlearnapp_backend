package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"readlearnapp/utils"
	"github.com/gin-gonic/gin"
)

func GetStories(c *gin.Context) {
	cursor, err := utils.DbCollections.Stories.Find(context.TODO(), bson.D{})
	if err != nil { log.Fatal(err)}
	var stories []utils.Story
	for cursor.Next(context.TODO()){
		story := utils.Story{}
		err = cursor.Decode(&story)
		if err != nil { log.Fatal(err)}
		stories = append(stories, story)
	}
	
	JSONresponse, err := json.Marshal(stories)
	if err != nil { log.Fatal(err)}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(JSONresponse)
}

func ServeStoryData(c *gin.Context){
	id := c.Query("id")
	resultS := utils.DbCollections.Stories.FindOne(context.TODO(), bson.D{{"_id",id}})
	opts := options.FindOptions{Projection: bson.D{{"_id",0}}}
	cursorD, err := utils.DbCollections.Dictionary.Find(context.TODO(), bson.D{{"ref",id}},&opts)
	if err != nil { log.Fatal(err)}
	
	story := utils.Story{}
	err = resultS.Decode(&story)
	if err != nil { log.Fatal(err)}
	
	var dictionary []utils.Dictionary
	for cursorD.Next(context.TODO()){
		word := utils.Dictionary{}
		err = cursorD.Decode(&word)
		if err != nil { log.Fatal(err)}
		dictionary = append(dictionary,word)
	}
	
	fmt.Println(story)
	fmt.Println(dictionary)
	
	JSONresponse, err := json.Marshal(dictionary)
	if err != nil { log.Fatal(err)}
	c.Writer.WriteHeader(404)
	c.Writer.Write(JSONresponse)
}