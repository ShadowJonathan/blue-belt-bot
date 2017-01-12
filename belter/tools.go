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

func CheckPCh(m *discordgo.Message) bool {
	return CheckPrivateChannel(m.ChannelID)
}

func CheckPrivateChannel(Ch string) bool { // probably very innefficient, but whatever
	Chstate, _ := bbb.dg.State.Channel(Ch)
	return Chstate.IsPrivate
}

func SwitchCMDType(m *discordgo.Message) (int, bool) {
	MC := m.Content
	switch MC[0] {
	case '!':
		return 0, false
	case ' ':
		if MC[1] == '!' {
			return 1, false
		}
	case '>':
		return 2, false
	case '?':
		return 3, false
	}
	return -1, true
}

func GetArgs(S string) []string {
	Args := []string{}
	Length := len(S)
	for i := 0; i < Length; i++ {
		Letter := S[i]
		if !IsSpace(Letter) {
			var startCMD int
			startCMD = i
			i++
			for i < Length && IsSpace(S[i]) {
				i++
			}
			Args = append(Args, S[startCMD:i])
		}
	}
	return Args
}
