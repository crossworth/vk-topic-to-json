package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	vkapi "github.com/himidori/golang-vk-api"

	topicToJSON "github.com/crossworth/vk-topic-to-json"
)

func main() {
	email := flag.String("email", "", "Seu email do VK")
	password := flag.String("senha", "", "Sua senha do VK")
	groupID := flag.String("comunidadeID", "", "O ID da comunidade")
	topicID := flag.String("topicoID", "", "O ID do tópico")

	flag.Parse()

	if *email == "" {
		log.Fatalf("Você deve informar o email")
	}

	if *password == "" {
		log.Fatalf("Você deve informar a senha")
	}

	if *groupID == "" {
		log.Fatalf("Você deve informar o ID da comunidade")
	}

	if *topicID == "" {
		log.Fatalf("Você deve informar o ID do tópico")
	}

	client, err := vkapi.NewVKClient(vkapi.DeviceIPhone, *email, *password)
	if err != nil {
		log.Fatalf("Não foi possível criar o cliente do VK, %v", err)
	}

	groupIDInt, err := strconv.Atoi(*groupID)
	if err != nil {
		log.Fatalf("O ID da comunidade não é válido, %v", err)
	}

	topicIDInt, err := strconv.Atoi(*topicID)
	if err != nil {
		log.Fatalf("O ID do tópico não é válido, %v", err)
	}

	topic, err := topicToJSON.SaveTopic(client, groupIDInt, topicIDInt)

	output, err := json.Marshal(topic)
	if err != nil {
		log.Fatalf("Não foi possível conveter o tópico para JSON, %v", err)
	}

	err = ioutil.WriteFile("backup_topico_"+*groupID+"_"+*topicID+".json", output, os.ModePerm)
	if err != nil {
		log.Fatalf("Não foi possível salvar o arquivo no disco, %v", err)
	}
}
