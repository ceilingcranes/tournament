package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/connordotfun/slack-off-Backend/message"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

type MessageStore struct {
	db *sql.DB
}

func NewDB() MessageStore {
	fmt.Println(os.Getenv("DB_URL"))
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return MessageStore{db: db}
}

//TODO: post votes, GetWinner

// Get all the Message structs from the current round
func (database MessageStore) GetBracketMessages(round int) []message.Message {
	command := `SELECT * FROM voting WHERE "bracket_count"='$1'`
	var messages []message.Message

	rows, err := database.db.Query(command, round)
	handle(err)

	for rows.Next() {
		messages = append(messages, convertRowToMessage(rows))
	}

	return messages
}

// Convert a row from the messages SQL database to a message struct
func convertRowToMessage(row *sql.Rows) message.Message {
	var channel string
	var id string
	var author string
	var text string
	var file string

	err := row.Scan(&channel, &author, &text, &file, &id)
	if err != nil {
		panic(err)
	}
	return message.Message{
		Channel: channel,
		Author:  author,
		ID:      id,
		Text:    text,
		File:    file}
}
