package Belt

import (
	fmt "fmt"
	"io/ioutil"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Version struct {
	Major               byte
	Minor               byte
	Build               byte
	Experimental        bool
	ExperimentalVersion byte
}

type BlueBot struct {
	dg      *discordgo.Session
	Debug   bool
	version Version
	OwnID   string
	OwnAV   string
	OwnName string
	Stop    bool
}

// Vars after this

var bbb *BlueBot

//BotFunc Fuction of the bot, can be set to "bot" or "user"
var BotFunc string

// Functions after this

func BBReady(s *discordgo.Session, r *discordgo.Ready) {
	bbb.OwnID = r.User.ID
	bbb.OwnAV = r.User.Avatar
	bbb.OwnName = r.User.Username
	fmt.Println("Discord: Ready message received\nBB: I am '" + bbb.OwnName + "'!\nBB: My User ID: " + bbb.OwnID)
}

func Initialize(Token string) {
	isdebug, err := ioutil.ReadFile("debugtoggle")
	bbb = &BlueBot{
		version: Version{0, 0, 1, true, 1},
		Debug:   (err == nil && len(isdebug) > 0),
		Stop:    false,
	}
	bbb.dg, err = discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Discord Session error, check token, error message: " + err.Error())
		return
	}
	// handlers
	bbb.dg.AddHandler(BBReady)

	fmt.Println("BB: Handlers installed")

	err = bbb.dg.Open()
	if err == nil {
		fmt.Println("Discord: Connection established")
		for !bbb.Stop {
			time.Sleep(400 * time.Millisecond)
		}
	} else {
		fmt.Println("Error opening websocket connection: ", err.Error())
	}
	fmt.Println("BBB: Blue belt stopping...")
	bbb.dg.Close()
}
