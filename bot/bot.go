package bot

import (
	"fmt"
	"irekean-discord-direr/config"
	"irekean-discord-direr/dice"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotId string

//var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		log.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		log.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	var RegexCommand = regexp.MustCompile(`^` + config.BotPrefixRegex + `[a-z]+`) //Regex to match the first part of the message like matching "/roll" in a message like "/roll 1d20+4"
	//var RegexContent = regexp.MustCompile(`^\/[a-z]+`)

	var MatchCommand = RegexCommand.FindString(m.Content)
	MatchCommand = strings.Replace(MatchCommand, config.BotPrefix, "", 1)

	switch MatchCommand {
	case "ping":
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	case "help":
		_, _ = s.ChannelMessageSend(m.ChannelID, config.HelpMessage)
	case "roll", "r":
		rollResult, err := dice.Roll(m.Content)
		if err == nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, rollResult)
		} else {
			log.Default().Println("Error occurred while rolling " + err.Error()) //TODO write the text of the roll
			_, _ = s.ChannelMessageSend(m.ChannelID, "An error occurred: "+err.Error())
		}
	}
}
