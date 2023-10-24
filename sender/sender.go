package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"../shared"
)

func insetMessageToDB(db *sql.DB, message shared.Message) {
	query := "INSERT INTO ASYNC_MESSAGE (SENDER_NAME, MESSAGE_TEXT, TIMESTAMP) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, message.senderName, message.messageText, message.timeStamp)
	if err != nil {
		log.Fatalf("Failed to insert a message: %v", err)
	}
}

func main() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password = %s dbname=%s sslmode=disable", shared.dbHost, shared.dbPort, shared.dbPassword, shared.dbName)
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

				message := shared.Message{
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
