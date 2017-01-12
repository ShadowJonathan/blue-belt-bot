package Belt

import (
	fmt "fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Processcommand(m *discordgo.Message) {
	args = GetArgs(m.Content[1:])
	CMD = strings.ToLower(args[0])
}

func ProcessCMD(m *discordgo.Message) {
	Args := GetArgs(m.Content[1:])
	CMD := Args[0]
}

func ProcessQuery(m *discordgo.Message) {
	Args := GetArgs(m.Content[1:])
	QU := Args[0]
}
