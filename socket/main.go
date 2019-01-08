package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

type socketData struct {
	Name int
	Time time.Time
}

const mqttTopic = "my-topic"
const mqttServer = "mqtt"
const delay = 10 * time.Second

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

	time.Sleep(7 * time.Second)
	//Code starts here
	for i := 0; i < 5; i++ {
		data, _ := json.Marshal(socketData{Name: 1, Time: time.Now()})
		fmt.Println(string(data))
		token := c.Publish(mqttTopic, 0, false, data)
		token.Wait()
		time.Sleep(delay)
	}

	c.Disconnect(250)
}
