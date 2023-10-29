package main

import (
	"database/sql"
	"fmt"
	"log"
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
				rows, err := db.Query(`SELECT ID, SENDER_NAME, MESSAGE, MESSAGE_TIME FROM ASYNC_MESSAGE WHERE RECEIVED_TIME IS NULL AND SENDER_NAME != $1 LIMIT 1`, serverIP)
				if err != nil {
					log.Println("Error querying message: ", err)
					continue
				}

				if rows.Next() {
					var id int
					var senderName, message string
					var sentTime time.Time
					if err := rows.Scan(&id, &senderName, &message, &sentTime); err != nil {
						log.Println("Error scanning message:", err)
					} else {
						fmt.Printf("Sender %s sent: %s at time %s\n", senderName, message, sentTime.Format("2006-01-02 15:04:05"))
						_, err := db.Exec(`UPDATE ASYNC_MESSAGE SET RECEIVED_TIME = $1 WHERE ID = $2`, time.Now(), id)
						if err != nil {
							log.Println("Error updating message:", err)
						}
					}
				}
				rows.Close()
			}
		}(serverIP)
	}

	wg.Wait()
}
