package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"chatting/config" //importing our config package which we have created above

	"github.com/bwmarrin/discordgo" //discordgo package from the repo of bwmarrin .
)

var (
	BotId      string
	goBot      *discordgo.Session
	copypastas *CopypastasStruct
)

func loadCopypastas() {
	// open and load copypastas
	fmt.Println("Reading copypasta file...")
	file, err := ioutil.ReadFile("./bot/data/copypasta.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal(file, &copypastas)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func initBot() error {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// Making our bot a user using User function .
	user, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Storing our id from u to BotId .
	BotId = user.ID

	// Adding handler function to handle our messages using AddHandler from discordgo package
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return err
	}

	return nil
}

func Start() {

	err := initBot()
	if err != nil {
		fmt.Println(err.Error())
	}

	loadCopypastas()

	fmt.Println("Bot is running !")
}

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	// if messageInput == config.BotPrefix+"" {
	// 	_, err = session.ChannelMessageSend(message.ChannelID, "")
	// 	fmt.Println(err)
	// 	fmt.Println("")
	// }

	messageInput := strings.ToLower(message.Content)
	var err error

	if message.Author.ID == BotId {
		return
	}

	if messageInput == config.BotPrefix+"ping" {
		_, err = session.ChannelMessageSend(message.ChannelID, "pong")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("pong")
	}

	if messageInput == config.BotPrefix+"debug" {
		_, err = session.ChannelMessageSend(message.ChannelID, "Debug:")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("debug")
	}

	if messageInput == config.BotPrefix+"help" {
		_, err = session.ChannelMessageSend(message.ChannelID, "Currently I have the following commands: \n!ping\n!scaryping\n!existentialping\nMore commands are coming pogu")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("help")
	}

	// https://tutorialedge.net/golang/parsing-json-with-golang/
	// https://www.sohamkamani.com/golang/json/
	// https://gobyexample.com/structs

	// if messageInput == config.BotPrefix+"scaryping" {
	// 	_, err = session.ChannelMessageSend(message.ChannelID, scaryPing)
	// 	fmt.Println(err)
	// 	fmt.Println("scaryping")
	// }

	// if messageInput == config.BotPrefix+"existentialping" {
	// 	_, err = session.ChannelMessageSend(message.ChannelID, existentialPing)
	// 	fmt.Println(err)
	// 	fmt.Println("existentialping")
	// }

	// if messageInput == config.BotPrefix+config.CopypastaPrefix+"scaryping" {

	// 	_, err = session.ChannelMessageSend(message.ChannelID, "")
	// 	fmt.Println(err)
	// 	fmt.Println("")
	// }

}
