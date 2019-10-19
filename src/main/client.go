// Create topic
package main

/**
 * Copyright 2018 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"kafka/consumer"
	"kafka/producer"
	"kafka/stats"
	"log"
	"net/http"
	"kafka/admin"
)
var broker = "k8scale-kafka-1:9092"


func createTopic(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating topic")
	topic := "test"
	numParts :=1
	replicationFactor :=1
	admin.CreateTopic(broker, topic, numParts, replicationFactor)
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending message")

	producer.SendData(broker, "test", []byte("this is a message"))
}

func receiveMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Consuming message ")
	topics := []string{"test"}
	consumer.Consume(broker, "", topics)
}

func receiveStats(w http.ResponseWriter, r *http.Request) {
	log.Println("Consuming message ")
	topics := []string{"test"}
	stats.GetStats(broker, "", topics)
}


func main() {
	http.HandleFunc("/create-topic", createTopic)
	http.HandleFunc("/send-message", sendMessage)
	http.HandleFunc("/receive-message", receiveMessage)
	http.HandleFunc("/stats", receiveStats)
	log.Fatal(http.ListenAndServe(":4040", nil))
}