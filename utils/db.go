package utils

import (
	"github.com/objectbox/objectbox-go/objectbox"
	"log"
)

var ObjectBox *objectbox.ObjectBox
var Stories *StoryBox
var Dictionaries *DictionaryBox
var err error
func InitObjectBox() {
	ObjectBox, err = objectbox.NewBuilder().Model(ObjectBoxModel()).Build()
	if err != nil {
		log.Fatal(err)
	}
	Stories = BoxForStory(ObjectBox)
	Dictionaries = BoxForDictionary(ObjectBox)
}
