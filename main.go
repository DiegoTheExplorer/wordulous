package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("wordulous")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file: ")
	}
	BOT_ACCESS_TOKEN := os.Getenv("DISCORD_BOT_ACCESS_TOKEN")

	session, err := discordgo.New("Bot " + BOT_ACCESS_TOKEN)

	if err != nil {
		log.Fatal("Error while creating discord session: ", err)
	}

	//session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate))

	err = session.Open()

	if err != nil {
		log.Fatal("Error while opening discord session: ", err)
	}

	defer session.Close()

	// GAME disc chat channel id: 844133446438748171
	msgs, err := session.ChannelMessages("844133446438748171", 5, "", "", "")

	if err != nil {
		log.Fatal("Discord message retrieval failure")
	}

	for index, element := range msgs {
		fmt.Println("message no. ", index)
		fmt.Println(element.Author.Username, ": ", element.ContentWithMentionsReplaced())
	}

	fmt.Println("Wordulous is online")
}
