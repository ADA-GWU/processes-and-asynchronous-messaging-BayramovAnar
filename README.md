# Please read the instructions carefully and see the GIF at the end 
### Once you have cloned the repo, do the following:   

1. Install Go from the official website (https://go.dev/doc/install) <br> 
2. Type **"go mod init"** in the terminal to track your dependency <br> 
3. Type **"go get -u github.com/lib/pq"** to import a database from the GitHub <br> 
4. Create the following PostgreSQL database in your system: <br>

**CREATE TABLE ASYNC_MESSAGE(** <br>
**ID SERIAL PRIMARY KEY,** <br>
**SENDER_NAME TEXT NOT NULL,** <br>
**MESSAGE TEXT NOT NULL,** <br>
**CURRENT_TIME TIMESTAMP NOT NULL DEFAULT NOW(),** <br>
**RECEIVED_TIME TIMESTAMP);** <br>

5. Go to the terminal and do the following, as shown in the GIF below. <br>
6. In both **sender.go** and **reader.go** change the values of **connectionString := "user=anar dbname=message_db sslmode=disable"** to whatever your database called and your username. <br>

![Video to GIF](https://github.com/ADA-GWU/processes-and-asynchronous-messaging-BayramovAnar/assets/98649599/ad9250b2-b630-414a-af7d-3b6385aaa7ac)
