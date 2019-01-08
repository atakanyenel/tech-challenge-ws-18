package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type socketData struct {
	Name int
	Time time.Time
}

const mqttServer = "mqtt"
const mqttTopic = "my-topic"
const mysqlServer = "mysql"

func main() {
	fmt.Println("sphere starting...")

	onlyServer := flag.Bool("server", false, "run only server")
	flag.Parse()

	cc := make(chan os.Signal, 1)
	signal.Notify(cc, os.Interrupt, syscall.SIGTERM)

	if !*onlyServer {

		// defer the close till after the main function has finished
		// executing

		opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:1883", mqttServer)).SetClientID("sphere")

		c := MQTT.NewClient(opts)

		if token := c.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		if token := c.Subscribe(mqttTopic, 0, nil); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		c.Subscribe(mqttTopic, 0, onMessageReceived)
	}
	go startServer()
	<-cc
}

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	var receivedMessage socketData
	json.Unmarshal(message.Payload(), &receivedMessage)
	fmt.Println(receivedMessage.Name)
	writeToDB(receivedMessage)
}
func writeToDB(receivedMessage socketData) {

	db, err := sql.Open("mysql", "root:example@tcp(mysql:3306)/test")
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("INSERT measurements SET socket_id=?,time=?")
	checkErr(err)

	_, err = stmt.Exec(receivedMessage.Name, receivedMessage.Time)

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
