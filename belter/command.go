package Belt

import (
	fmt "fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Processcommand(m *discordgo.Message, g *discordgo.Guild) {
	if m.Content[0] == '!' || (m.Content[0] == ' ' && m.Content[1] == '!') {
		var args []string
		var CMD string
		if m.Content[0] == '!' {
			args = GetArgs(m.Content[1:])
		} else {
			if m.Content[1] == '!' {
				args = GetArgs(m.Content[2:])
			}
		}
		CMD = strings.ToLower(args[0])
	}
}
