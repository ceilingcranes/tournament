package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/connordotfun/slack-off-Backend/message"
	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	fmt.Println(os.Getenv("DB_URL"))
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ReadData() []message.Message {
	messages := make([]message.Message, 0)
	js, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(js, &messages)
	if err != nil {
		panic(err)
	}
	return messages
}

func FillDB(db *sql.DB, messages []message.Message) {
	command := "INSERT INTO messages (channel, author, text, file, id) VALUES ($1, $2, $3, $4, $5);"
	count := 0
	for _, message := range messages {
		_, err := db.Exec(command, message.Channel, message.Author, message.Text, message.File, message.ID)
		if err != nil {
			panic(err)
		}
		count++
	}
	fmt.Println("Final Count: " + strconv.Itoa(count))
}
