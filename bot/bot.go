package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/to4to/Go-Discord-Bot1/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is Running")

}


func messageHandler(s *discordgo.Session,m *discordgo.MessageCreate){

if m.Author.ID==BotID{
	return 
}

if m.Content=="who made you"{
_,_  =s.ChannelMessageSend(m.ChannelID,"visit him at:  github.com/to4to  ")
}




}