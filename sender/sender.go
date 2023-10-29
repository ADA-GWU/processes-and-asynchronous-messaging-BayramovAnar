package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connectionString := "user=anar dbname=message_db sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var serverIPs = []string{"db1"}
	var wg sync.WaitGroup

	for _, serverIP := range serverIPs {
		wg.Add(1)
		go func(serverIP string) {
			defer wg.Done()
			for {
				var senderName, message string
				var reader = bufio.NewReader(os.Stdin)

				fmt.Printf("Enter your name: ")
				senderName, _ = reader.ReadString('\n')
				senderName = senderName[:len(senderName)-1]

				fmt.Printf("Enter message: ")
				message, _ = reader.ReadString('\n')
				message = message[:len(message)-1]

				sentTime := time.Now()
				_, err := db.Exec(`INSERT INTO ASYNC_MESSAGE (SENDER_NAME, MESSAGE, MESSAGE_TIME) VALUES ($1, $2, $3)`, senderName, message, sentTime.Format("2006-01-02 15:04:05"))
				if err != nil {
					log.Println("Can't send a message: ", err)
				}
			}
		}(serverIP)
	}

	wg.Wait()
}
