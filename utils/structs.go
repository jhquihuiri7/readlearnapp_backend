package utils

type Story struct {
	Id          string `bson:"_id"`
	Title       string	`bson:"title"`
	Image       string	`bson:"image"`
	Text        string	`bson:"text"`
	Scroll      float64	`bson:"scroll"`
	Progress    float64`bson:"progress"`
	IsCompleted bool	`bson:"isCompleted"`
}
type Dictionary struct {
	Id          string	`bson:"_id"`
	Ref         string	`bson:"ref"`
	Word        string	`bson:"word"`
	Meaning     string	`bson:"meaning"`
	Translation string	`bson:"translation"`
	Example     string	`bson:"example"`
}




//, err := os.Open("story.txt")
//f err != nil {
//	log.Fatal(err)
//
//, err := io.ReadAll(s)
//f err != nil {
//	log.Fatal(err)
//

//tory := Story{
//	Id:          uuid.NewV4().String(),
//	Title:       "La Sirenita",
//	Image:       "imag.com",
//	Text:        string(b),
//	}
//	f, err := os.Open("dataCleaned.csv")
//	if err != nil {
//		log.Fatal(err)
//	}
//	reader := csv.NewReader(f)
//	records, err := reader.ReadAll()
//	if err != nil {
//		log.Fatal(err)
//	}
//	var allDict []interface{}
//	for _, v := range records[1:] {
//		dict := Dictionary{
//			Id:          uuid.NewV4().String(),
//			Ref:         story.Id,
//			Word:        v[0],
//			Meaning:     "",
//			Translation: v[1],
//			Example:     v[2],
//			}
//			allDict = append(allDict, dict)
//	}