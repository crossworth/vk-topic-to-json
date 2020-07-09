package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"

	vkapi "github.com/himidori/golang-vk-api"

	topicToJSON "github.com/crossworth/vk-topic-to-json"
)

var (
	urlRegex = regexp.MustCompile(`topic-([0-9]+)_([0-9]+)`)
)

func main() {
	email := flag.String("email", "", "Your VK Email")
	password := flag.String("password", "", "Your VK Password")
	groupID := flag.Int("group", 0, "GroupID (if no url is provided)")
	topicID := flag.Int("topic", 0, "TopicID (if no url is provided)")
	url := flag.String("url", "", "Topic URL")

	flag.Parse()

	if *email == "" {
		log.Fatalf("You must provide an email")
	}

	if *password == "" {
		log.Fatalf("You must provide a password")
	}

	if *url != "" {
		matches := urlRegex.FindStringSubmatch(*url)
		if len(matches) == 3 {
			*groupID = mustInt(matches[1])
			*topicID = mustInt(matches[2])
		}
	}

	if *groupID == 0 {
		log.Fatalf("You must provide a group ID or URL")
	}

	if *topicID == 0 {
		log.Fatalf("You must provide a topic ID or URL")
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

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("error parsing %q as integer, %v", s, err)
	}

	return i
}
