package db

import (
	"database/sql"
	"fmt"

	"github.com/connordotfun/slack-off-Backend/message"
)

// Get all the Message structs from the current round
func GetBracketMessages(round int, database *sql.DB) {
	command := `SELECT * FROM voting WHERE "bracket_count"='$1'`
	row := database.QueryRow(command, round)
	message := convertRowToMessage(row)

	fmt.Println("Read message: " + message)
}

// Convert a row from the messages SQL database to a message struct
func convertRowToMessage(row *sql.Row) message.Message {
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
