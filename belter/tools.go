package Belt

import (
	"io"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func IsSpace(L byte) bool {
	return L == ' ' || L == '\t' || L == '\r'
}

func DownloadUrl(URL string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func CheckPCh(m *discordgo.Message) {
	return CheckPrivateChannel(m.ChannelID)
}

func CheckPrivateChannel(Ch string) { // probably very innefficient, but whatever
	return bbb.dg.State.Channel(Ch).IsPrivate
}

func SwitchCMDType(m *discordgo.Message) int {
	MC := m.Content
	switch MC[0] {
	case '!':
		return 0
	case ' ' || '\t' || '\r':
		if MC[1] == '!' {
			return 1
		}
	case '>':
		return 2
	case '?':
		return 3
	default:
		return nil
	}
}

func GetArgs(S string) []string {
	Args := []string{}
	Length := len(S)
	for i := 0; i < Length; i++ {
		Letter := S[i]
		if !IsSpace(Letter) {
			var startCMD int
			start = i
			i++
			for i < Length && IsSpace(s[i]) {
				i++
			}
			Args = append(Args, S[startCMD:i])
		}
	}
	return Args
}
