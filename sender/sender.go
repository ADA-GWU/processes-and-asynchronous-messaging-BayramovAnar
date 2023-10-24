package sender

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

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

func insetMessageToDB(db *sql.DB, message Message) {
	query := "INSERT INTO ASYNC_MESSAGE (SENDER_NAME, MESSAGE_TEXT, TIMESTAMP) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, message.senderName, message.messageText, message.timeStamp)
	if err != nil {
		log.Fatalf("Failed to insert a message: %v", err)
	}
}

func main() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password = %s dbname=%s sslmode=disable", dbHost, dbPort, dbPassword, dbName)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect to the database: %v", err)
	}

	dbServers := []string{"DBServer1", "DBServer2", "DBServer3"}

	var wg sync.WaitGroup

	for _, dbServer := range dbServers {
		wg.Add(1)

		go func(server string) {
			defer wg.Done()
			for {
				var messageText string
				fmt.Printf("Enter a message to send to %s: ", server)
				fmt.Scanln(&messageText)

				message := Message{
					senderName:  "Anar",
					messageText: messageText,
					timeStamp:   time.Now(),
				}
				insetMessageToDB(db, message)
			}
		}(dbServer)
	}
	wg.Wait()
}