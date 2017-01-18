package Belt

import (
	fmt "fmt"
	"io/ioutil"
	"time"

	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

// Vars after this

var bbb *BlueBot

//BotFunc Fuction of the bot, can be set to "bot" or "user"
var BotFunc string

// Functions after this

// Handlers

func BBFriendship(F *discordgo.RelationshipAdd) {
	Bot, _ := GetBot(F.User)
	if Bot.WatchOut == false {
		bbb.dg.RelationshipFriendRequestAccept(F.ID)
	}
}

func BBReady(s *discordgo.Session, r *discordgo.Ready) {
	bbb.OwnID = r.User.ID
	bbb.OwnAV = r.User.Avatar
	bbb.OwnName = r.User.Username
	bbb.OwnInfo, _ = GetBot(r.User)
	bbb.OwnInfo.Version = bbb.version
	fmt.Println("Discord: Ready message received\nBB: I am '" + bbb.OwnName + "'!\nBB: My User ID: " + bbb.OwnID)
}

func BBMessageCreate(Ses *discordgo.Session, MesC *discordgo.MessageCreate) {
	// stuff here
	Mes := MesC.Message
	if Mes.Content[1:] == ">open" {
		CALLDIAL(Mes.Author, Mes.ChannelID, Mes.Content[0])
		return
	}
	if bbb.Sessions[Mes.Author.ID].Dailed == true {
		switch bbb.Sessions[Mes.Author.ID].OwnSession {
		case '1':
			bbb.SessionOne <- Mes.Content
		case '2':
			bbb.SessionTwo <- Mes.Content
		case '3':
			bbb.SessionThree <- Mes.Content
		case '4':
			bbb.SessionFour <- Mes.Content
		}
	}
	CI, IsCMD := SwitchCMDType(Mes)
	if IsCMD == false {
		return
	}
	switch CI {
	case '0':
		Processcommand(Mes)
	case '1':
		Mesedit := Mes
		Mesedit.Content = Mesedit.Content[1:]
		Processcommand(Mesedit)
	case '2':
		ProcessCMD(Mes)
	case '3':
		ProcessQuery(Mes)
	}
}

// init

func Initialize(Token string) {
	isdebug, err := ioutil.ReadFile("debugtoggle")
	var BI []*BotInfo
	BIB, err2 := ioutil.ReadFile("seenbots")
	err3 := json.Unmarshal(BIB, BI)
	if err2 != nil {
		fmt.Println("SeenBots file not read: " + err2.Error())
	}
	if err3 != nil {
		fmt.Println("Seenbots incorrectly Unmarshald: " + err3.Error())
	}
	var MB = &MasterInfo{}
	MBB, err2 := ioutil.ReadFile("../Master")
	err3 = json.Unmarshal(MBB, MB)
	if err2 != nil {
		fmt.Println("Master file not read: " + err2.Error())
	}
	if err3 != nil {
		fmt.Println("Master incorrectly Unmarshald: " + err3.Error())
	}
	bbb = &BlueBot{
		version:      Version{0, 1, 0, false, 0},
		Debug:        (err == nil && len(isdebug) > 0),
		Stop:         false,
		FreeSessions: make(chan byte, 4),
		SessionOne:   make(chan string),
		SessionTwo:   make(chan string),
		SessionThree: make(chan string),
		SessionFour:  make(chan string),
		SeenBots:     BI,
		Master:       MB,
	}
	bbb.FreeSessions <- '1'
	bbb.FreeSessions <- '2'
	bbb.FreeSessions <- '3'
	bbb.FreeSessions <- '4'
	UserFile, err := ioutil.ReadFile("IsUser")
	if UserFile != nil {
		bbb.dg, err = discordgo.New(Token)
	} else {
		bbb.dg, err = discordgo.New("Bot " + Token)
	}
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
