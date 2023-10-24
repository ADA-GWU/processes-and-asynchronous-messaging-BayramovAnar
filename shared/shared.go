package shared

import "time"

type Message struct {
	senderName  string
	messageText string
	timeStamp   time.Time
}

const (
	dbHost     = ""
	dbPort     = 8081
	dbUser     = "anar"
	dbPassword = ""
	dbName     = "template1"
)
