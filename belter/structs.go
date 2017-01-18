package Belt

import "github.com/bwmarrin/discordgo"

type Version struct {
	Major               byte
	Minor               byte
	Build               byte
	Experimental        bool
	ExperimentalVersion byte
}

type BlueBot struct {
	dg           *discordgo.Session
	Debug        bool
	version      Version
	OwnID        string
	OwnAV        string
	OwnName      string
	Stop         bool
	SeenBots     []*BotInfo
	Master       *MasterInfo
	Sessions     map[string]*Session
	FreeSessions chan byte
	SessionOne   chan string
	SessionTwo   chan string
	SessionThree chan string
	SessionFour  chan string
	OwnInfo      *BotInfo
}

type MasterInfo struct {
	ID               string
	DMChannel        string
	Fullname         string
	directconnection bool // aka, have a "friendship" open with master
}

type BotInfo struct {
	BotUName   string
	BotHash    string
	BotID      string
	BotSC      string // currently only "BB", bluebelt
	BotType    string // "Master", "Slave", "Relay"
	Version    Version
	FromMaster bool // info comes from the master bot
	IsEdited   bool // pushed bool from master bots that detected a change or weird behaviour from the bot
	WatchOut   bool // pushed bool from master bots and Root users that signaled that the bot is potenbtionally dangerous
}

type Session struct {
	IsBot  bool
	Dailed bool
	Dail   struct {
		Timeout      int
		LastActivity int
		Lastfailed   bool
		RetriesLeft  int
	}
	User         *discordgo.User
	IsOpener     bool
	OtherSession byte
	OwnSession   byte
	SessionSync  bool
}
