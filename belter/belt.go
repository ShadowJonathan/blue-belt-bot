package Belt

import "github.com/bwmarrin/discordgo"
import "fmt"

type Version struct {
	Major               byte
	Minor               byte
	Build               byte
	Experimental        bool
	ExperimentalVersion byte
}

type BlueBot struct {
	dg      *discordgo.Session
	version Version
}

// Vars after this

var bbb *BlueBot

// Functions after this

func Initialize(Token string) {
	bbb = &Bluebot{
		version: Version{0, 0, 1, true, 1},
	}

	bbb.dg, err = discordgo.New(Token)
	if err != nil {
		fmt.Printl("Discord Session error, check token, error message: " + err.Error())
		return
	}
}
