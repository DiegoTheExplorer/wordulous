package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/psykhi/wordclouds"
)

func main() {
	fmt.Println("Wordulous is online")

	// Load .env file for the bot access token
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file: ")
	}
	BOT_ACCESS_TOKEN := os.Getenv("DISCORD_BOT_ACCESS_TOKEN")

	//	Initialize a new discord bot session using the access token
	session, err := discordgo.New("Bot " + BOT_ACCESS_TOKEN)
	if err != nil {
		log.Fatal("Error while creating discord session: ", err)
	}

	//session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate))
	//	Open/Stat the discord session
	err = session.Open()
	if err != nil {
		log.Fatal("Error while opening discord session: ", err)
	}
	defer session.Close()

	CHAT_CHANNEL_ID := os.Getenv("GAME_DISC_CHAT_ID")
	msgs, err := session.ChannelMessages(CHAT_CHANNEL_ID, 100, "", "", "")
	if err != nil {
		log.Fatal("Discord message retrieval failure: ", err)
	}

	dict := map[string]int{}
	corrupted := 0

	for _, msg := range msgs {
		content, err := msg.ContentWithMoreMentionsReplaced(session)
		if err != nil {
			corrupted++
			continue
		}

		//remove all non-alphanumeric characters
		//convert everything to lowercase
		content = strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9\t ]`).ReplaceAllString(content, ""))

		//tokenize each message into words
		tokens := strings.Fields(content)

		for _, word := range tokens {
			count := dict[word]
			if count == 0 {
				dict[word] = 1
				continue
			}
			dict[word]++
		}
	}

	fmt.Print(dict)

	w_cloud := wordclouds.NewWordcloud(
		dict,
		wordclouds.FontFile("fonts/MontserratBlack.ttf"),
		wordclouds.Height(2048),
		wordclouds.Width(2048),
		wordclouds.Colors([]color.Color{
			color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // Red
			color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // Green
			color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // Blue
		}),
		wordclouds.BackgroundColor(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}), // White background
	)

	// Draw the word cloud
	img := w_cloud.Draw()

	// Save the image to the pc
	f, err := os.Create("output.png")
	if err != nil {
		log.Fatal("Error while creating outpet image file: ", err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		log.Fatal("PNG Encoding error: ", err)
	}

	// for _, msg := range msgs {
	// 	fmt.Println(msg.Author.Username, ": ")
	// 	content, err := msg.ContentWithMoreMentionsReplaced(session)
	// 	if err != nil {
	// 		log.Fatal("Message content with more mentions replaced error: ", err)
	// 	}
	// 	fmt.Println(content)
	// }

}
