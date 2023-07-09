//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen
package utils

type Story struct {
	Id          uint64 `json:"id"`
	Title       string	`bson:"title"`
	Image       string	`bson:"image"`
	Text        string	`bson:"text"`
	Scroll      float64	`bson:"scroll"`
	Progress    float64`bson:"progress"`
	IsCompleted bool	`bson:"isCompleted"`
}
type Dictionary struct {
	Id          uint64	`json:"id"`
	Title         string`bson:"title"`
	Word        string	`bson:"word"`
	Meaning     string	`bson:"meaning"`
	Translation string	`bson:"translation"`
	Example     string	`bson:"example"`
}
