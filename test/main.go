package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"os"
	"time"
)

type socketData struct {
	Name int
	Time time.Time
}

const mqttTopic = "my-topic"
const mqttServer = "mqtt"
const delay = 1 * time.Second

func main() {
	fmt.Println("Socket starting...")

	clientID, _ := os.Hostname()
	connOpts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:1883", mqttServer)).SetClientID(clientID).SetCleanSession(true)

	c := MQTT.NewClient(connOpts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe(mqttTopic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	timeArray := []time.Time{
		time.Date(2019, time.January, 13, 9, 0, 0, 0, time.UTC),
		time.Date(2019, time.January, 14, 9, 0, 0, 0, time.UTC),
		time.Date(2019, time.January, 15, 9, 0, 0, 0, time.UTC),
		time.Date(2019, time.January, 16, 9, 0, 0, 0, time.UTC),
		time.Date(2019, time.January, 17, 9, 0, 0, 0, time.UTC),
		time.Date(2019, time.January, 18, 9, 0, 0, 0, time.UTC),
		time.Date(2019, time.January, 19, 9, 0, 0, 0, time.UTC),
		time.Date(2019, time.January, 20, 9, 0, 0, 0, time.UTC),
	}
	socketIDs := []int{1, 2, 3, 4}
	//Code starts here
	for _, t := range timeArray {
		for _, id := range socketIDs {

			data, _ := json.Marshal(socketData{Name: id, Time: t})
			rand.Seed(time.Now().UTC().UnixNano())
			numberOfTimes := rand.Intn(40)
			for i := 0; i < numberOfTimes; i++ {
				token := c.Publish(mqttTopic, 0, false, data)
				token.Wait()
				fmt.Printf("Sent: %v\n", data)
				time.Sleep(delay)
			}
		}
	}

	c.Disconnect(250)
}
