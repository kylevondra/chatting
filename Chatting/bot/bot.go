package bot

import (
	"chatting/config" //importing our config package which we have created above
	"fmt"             //to print errors
	"strings"

	"github.com/bwmarrin/discordgo" //discordgo package from the repo of bwmarrin .
)

var BotId string
var goBot *discordgo.Session

type Copypastas struct {
	Copypastas []Copypasta `json:"copypastas"`
}

type Copypasta struct {
	name string `json:"name"`
	data string `json:"data"`
}

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Making our bot a user using User function .
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Storing our id from u to BotId .
	BotId = u.ID

	// Adding handler function to handle our messages using AddHandler from discordgo package. We will declare messageHandler function later.
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	//If every thing works fine we will be printing this.
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
		fmt.Println(err)
		fmt.Println("pong")
	}

	if messageInput == config.BotPrefix+"scaryping" {
		_, err = session.ChannelMessageSend(message.ChannelID, scaryPing)
		fmt.Println(err)
		fmt.Println("scaryping")
	}

	if messageInput == config.BotPrefix+"existentialping" {
		_, err = session.ChannelMessageSend(message.ChannelID, existentialPing)
		_, err = session.ChannelMessageSend(message.ChannelID, existentialPing2)
		fmt.Println(err)
		fmt.Println("existentialping")
	}

	if messageInput == config.BotPrefix+"help" {
		_, err = session.ChannelMessageSend(message.ChannelID, "Currently I have the following commands: \n!ping\n!scaryping\n!existentialping\nMore commands are coming pogu")
		fmt.Println(err)
		fmt.Println("help")
	}

}
