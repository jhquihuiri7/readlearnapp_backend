package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"readlearnapp/utils"
)

func GetStories(c *gin.Context) {

	stories , err := utils.Stories.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	
	JSONresponse, err := json.Marshal(stories)
	if err != nil { log.Fatal(err)}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(JSONresponse)
}

func ServeStoryData(c *gin.Context){
	title := c.Query("title")
	dictionaryQuery := utils.Dictionaries.Query(utils.Dictionary_.Title.Equals(title, false))
	dictionary, err := dictionaryQuery.Find()
	fmt.Println(dictionary)
	JSONresponse, err := json.Marshal(dictionary)
	if err != nil { log.Fatal(err)}
	c.Writer.WriteHeader(404)
	c.Writer.Write(JSONresponse)
}
func AddData(c *gin.Context){
	var st []utils.Story
	var dc []*utils.Dictionary
	fst, err := os.Open("LearnReadApp.Stories.json")
	if err != nil {
		log.Fatal(err)
	}
	fdc, err := os.Open("LearnReadApp.Dictionary.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(fst)
	err = decoder.Decode(&st)
	if err != nil {
		log.Fatal(err)
	}
	decoder = json.NewDecoder(fdc)
	err = decoder.Decode(&dc)
	if err != nil {
		log.Fatal(err)
	}

	_, err = utils.Stories.Put(&st[0])
	if err != nil {
		log.Fatal(err)
	}

	_, err = utils.Dictionaries.PutMany(dc)
	if err != nil {
		log.Fatal(err)
	}
}