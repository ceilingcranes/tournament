package bracket

import (
	"tournament/db"

	"github.com/connordotfun/slack-off-Backend/message"
)

/*
Each bracket is used to compare two contestants during the tournament. Each bracket contains two contestants, and contains the number of points for each.
*/
type Bracket struct {
	ID          int
	Contestant1 *message.Message
	Contestant2 *message.Message
	Points1     int
	Points2     int
}

/* given the total number of messages and the current round, return a list of Bracket objects in order.
So, if it's 64 total messages, and it's round 3, there will be a slice of Brackets of length 16. */
func (database db.MessageStore) GetBrackets(round int, totNumMessages int) []Bracket {
	roundRows := database.GetBracketMessages(round)
	
}
