package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	vkapi "github.com/himidori/golang-vk-api"

	topicToJSON "github.com/crossworth/vk-topic-to-json"
)

func main() {
	email := flag.String("email", "", "Your VK Email")
	password := flag.String("password", "", "Your VK Password")
	groupID := flag.Int("group", 0, "GroupID")
	topicID := flag.Int("topic", 0, "TopicID")

	flag.Parse()

	if *email == "" {
		log.Fatalf("You must provide an email")
	}

	if *password == "" {
		log.Fatalf("You must provide a password")
	}

	if *groupID == 0 {
		log.Fatalf("You must provide a group ID")
	}

	if *topicID == 0 {
		log.Fatalf("You must provide a topic ID")
	}

	client, err := vkapi.NewVKClient(vkapi.DeviceIPhone, *email, *password)
	if err != nil {
		log.Fatalf("Could not create the VK client, %v", err)
	}

	topic, err := topicToJSON.SaveTopic(client, *groupID, *topicID)

	output, err := json.Marshal(topic)
	if err != nil {
		log.Fatalf("Could not marshal topic to JSON, %v", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("vk_topic_backup_%d_%d.json", *groupID, *topicID), output, os.ModePerm)
	if err != nil {
		log.Fatalf("Could not save topic to disc, %v", err)
	}
}
