package main

import (
	"database/sql"
	"fmt"
	"log"
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

func getAvailableMessagesFromDB(db *sql.DB, dbServer string) []Message {
	query := "SELECT SENDER_NAME, MESSAGE_TEXT, TIMESTAMP FROM ASYNC_MESSAGE WHERE RECEIVED_TIME IS NULL AND SENDER_NAME != $1"
	rows, err := db.Query(query, "Anar")
	if err != nil {
		log.Fatalf("Failed to query messages: %v", err)
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.senderName, &message.messageText, &message.timeStamp); err != nil {
			log.Printf("Failed to scan message: %v", err)
			continue
		}
		messages = append(messages, message)
	}
	return messages
}

func markMessageAsReceived(db *sql.DB, message Message) {
	query := "UPDATE ASYNC_MESSAGE SET RECEIVED_TIME = $1 WHERE SENDER_NAME = $2 AND MESSAGE_TEXT = $3"
	_, err := db.Exec(query, time.Now(), message.senderName, message.messageText)
	if err != nil {
		log.Printf("Failed to mark message as received: %v", err)
	}
}

func main() {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password = %s dbname=%s sslmode=disablet", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbInfo)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	dbServers := []string{"DBServer1", "DBServer2", "DBServer3"}

	for _, dbServer := range dbServers {
		go func(server string) {
			for {
				messages := getAvailableMessagesFromDB(db, server)
				for _, message := range messages {
					// TODO: implement database locking
					time.Sleep(time.Second)
					fmt.Printf("Sender %s sent '%s' at time %s\n", message.senderName, message.messageText, message.timeStamp)
					markMessageAsReceived(db, message)
				}
			}
		}(dbServer)
	}

	select {}
}
