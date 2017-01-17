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
