package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
)

var (
	Token          string
	BotPrefix      string
	BotPrefixRegex string
	HelpMessage    string

	config *configStruct
)

type configStruct struct {
	Token       string `json:"Token"`
	BotPrefix   string `json:"BotPrefix"`
	HelpMessage string `json:"HelpMessage"`
}

func ReadConfig() error {
	log.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	//log.Println(string(file)) //file with secret token should not be printed

	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	HelpMessage = config.HelpMessage
	BotPrefixRegex = regexp.QuoteMeta(config.BotPrefix)

	return nil

}
