package Belt

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// three command types (for users)

func Processcommand(m *discordgo.Message) {
	args := GetArgs(m.Content[1:])
	CMD := strings.ToLower(args[0])
	fmt.Println(CMD + " called")
}

func ProcessCMD(m *discordgo.Message) {
	Args := GetArgs(m.Content[1:])
	CMD := Args[0]
	fmt.Println(CMD + " CMD-called")
}

func ProcessQuery(m *discordgo.Message) {
	Args := GetArgs(m.Content[1:])
	QU := Args[0]
	fmt.Println(QU + " Queried")
}

// more commands for bots

func SwitchBotCommand(m *discordgo.Message, relayID string) {
	Bot, success := GetBot(m.Author)
	if !success {
		return
	}
	switch m.Content[1] {
	case '?':
		PCQ(m, Bot, relayID)
	case '!':
		PCN(m, Bot, relayID)
	case '%':
		PEVAL(m, Bot, relayID)
	case '>':
		PCPUSH(m, Bot, relayID)
	case '<':
		PCPULL(m, Bot, relayID)
	case '^':
		PCGET(m, Bot, relayID)
	case '&':
		PCRELAY(m, Bot)
	}
}

// ? Query
func PCQ(m *discordgo.Message, Bot *BotInfo, relay string) {

}

// ! Notice

func PCN(m *discordgo.Message, Bot *BotInfo, relay string) {

}

// % Evaluate

func PEVAL(m *discordgo.Message, Bot *BotInfo, relay string) {

}

// > Push

func PCPUSH(m *discordgo.Message, Bot *BotInfo, relay string) {

}

// < Pull

func PCPULL(m *discordgo.Message, Bot *BotInfo, relay string) {

}

// ^ Get info

func PCGET(m *discordgo.Message, Bot *BotInfo, relay string) {

}

// & Relay

func PCRELAY(m *discordgo.Message, Bot *BotInfo) {
	S, T, ID, CMD := DECOMP(m.Content[2:])
	if S == bbb.Sessions[m.Author.ID].OwnSession && T == '&' {
		relays := strings.Split(CMD, " ") // relayto relayfrom
		// relayto is incoming, relayfrom is where the command originally is from
		if ID == bbb.OwnID {
			M := m
			M.Content = string(S) + strings.Join(relays[1:], "")
			SwitchBotCommand(M, relays[1])
		} else {
		}
	}
}
