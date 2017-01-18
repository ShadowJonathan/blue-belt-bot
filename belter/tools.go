package Belt

import (
	"io"
	"net/http"
	"os"
	"time"

	"strconv"

	"encoding/json"

	"strings"

	"io/ioutil"

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

func FindProxy(ID string) string {

}

func GetBot(B *discordgo.User) (*BotInfo, bool) {
	for _, SB := range bbb.SeenBots {
		if SB.BotID == B.ID {
			return SB, false
		}
	} // if no bot was found
	response, err := QUERY("GETBOTINFO", true, "")
	if err == "" {
		_, C, RE, para := DECOMP(response)
		Bot := &BotInfo{}
		if C == '>' && RE == "BOTINFO" {
			json.Unmarshal([]byte(para), Bot)
			addBot(Bot)
			return Bot, true
		}
	}
	Bot := &BotInfo{}
	return Bot, false
}

func WaitForFriendship(ID string) (bool, bool) {
	Check := IsFriends(ID)
	if Check {
		return true, false
	}
	bbb.dg.RelationshipFriendRequestSend(ID)
	for i := 0; 60 < i; i++ {
		if IsFriends(ID) {
			return true, false
		}
		time.Sleep(1 * time.Second)
	}
	return false, true
}

func IsFriends(ID string) bool {
	for _, R := range bbb.dg.State.Relationships {
		if ID == R.ID {
			return true
		}
	}
	return false
}

func GetFreeCh() byte {
	return <-bbb.FreeSessions
}

func ToLetter(I int) string {
	return strconv.FormatInt(int64(I), 10)
}

func ToNumber(S string) int {
	Int, _ := strconv.ParseInt(S, 10, 0)
	return int(Int)
}

func BtN(B byte) int {
	I, _ := strconv.ParseInt(string(B), 10, 0)
	return int(I)
}

func NtB(N int) byte {
	B := strconv.FormatInt(int64(N), 10)
	return B[0]
}

func GetUser(ID string) (*discordgo.User, bool) {
	GS := bbb.dg.State.Guilds
	for _, G := range GS {
		for _, M := range G.Members {
			if M.User.ID == ID {
				return M.User, true
			}
		}
	}
	return nil, false
}

func GetUserChannel(ID string) (*discordgo.Channel, error) {
	Ch, err := bbb.dg.UserChannelCreate(ID)
	return Ch, err
}

func Compile(T interface{}) string {
	data, _ := json.Marshal(T)
	return string(data)
}

func DECOMP(S string) (byte, byte, string, string) {
	s := S[0]
	C := S[1]
	Ss := strings.Split(S[2:], " ")
	CMD := Ss[0]
	para := S[4+len(CMD):]
	return s, C, CMD, para
}

func GFS(S int) string {
	switch S {
	case 1:
		return <-bbb.SessionOne
	case 2:
		return <-bbb.SessionTwo
	case 3:
		return <-bbb.SessionThree
	case 4:
		return <-bbb.SessionFour
	}
	return ""
}

func addBot(Bot *BotInfo) {
	for _, SB := range bbb.SeenBots {
		if SB.BotID == Bot.BotID {
			return
		}
	}
	bbb.SeenBots = append(bbb.SeenBots, Bot)
	go WriteBots()
}

func WriteBots() {
	Bots, err := json.Marshal(bbb.SeenBots)
	if err != nil {
		ioutil.WriteFile("../seenbots", Bots, 0777)
	}
}
