package Belt

import (
	"encoding/json"

	"time"

	"github.com/bwmarrin/discordgo"
)

func DailLoop(HookID string, finishchannel byte) {
	var activityreported int
	for i := 0; i < bbb.Sessions[HookID].Dail.Timeout; i++ {
		time.Sleep(1 * time.Second)
		if activityreported != bbb.Sessions[HookID].Dail.LastActivity {
			i = 0
			activityreported = bbb.Sessions[HookID].Dail.LastActivity
		}
		if bbb.Sessions[HookID].Dailed == false {
			return
		}
	}
	DAILCLOSE(HookID, finishchannel)
}

func QUERY(CMD string, MASTER bool, ID string) (string, string) { // single query
	if MASTER {
		Dsuccess := DAIL(bbb.Master.ID)
		if !Dsuccess {
			var err = "DAIL FAIL"
			return "", err
		}
		Ses := bbb.Sessions[bbb.Master.ID]
		var Ch byte
		if Ses.IsOpener || Ses.SessionSync {
			Ch = Ses.OwnSession
		} else {
			Ch = Ses.OtherSession
		}
		bbb.dg.ChannelMessageSend(bbb.Master.DMChannel, string(Ch)+"?"+CMD)
		response := GFS(BtN(Ch))
		return response, ""

	} else {
		Dsuccess := DAIL(ID)
		if !Dsuccess {
			var err = "DAIL FAIL"
			return "", err
		}
		Ses := bbb.Sessions[ID]
		var Ch byte
		if Ses.IsOpener || Ses.SessionSync {
			Ch = Ses.OwnSession
		} else {
			Ch = Ses.OtherSession
		}
		Channel, _ := GetUserChannel(ID)
		bbb.dg.ChannelMessageSend(Channel.ID, string(Ch)+"?"+CMD)
		response := GFS(BtN(Ch))
		return response, ""
	}
	return "", ""
}

func DAILCLOSE(ID string, channel byte) {
	bbb.Sessions[ID].Dailed = false
	bbb.FreeSessions <- channel
}

func DAIL(ID string) bool {
	SB := GetFreeCh()
	Session := ToNumber(string(SB))
	User, found := GetUser(ID)
	if bbb.Sessions[ID] != nil {
		if bbb.Sessions[ID].Dailed == true {
			return true
		}
	}
	if found {
		Bot, frommaster := GetBot(User)
		if frommaster {
			addBot(Bot)
		}
	} else {
		yes, timeout := WaitForFriendship(ID)
		if !yes {
			if timeout {
				return false
			} else { // this should NEVER happen
				panic(yes)
			}
		}
		Bot, frommaster := GetBot(User)
		if frommaster {
			addBot(Bot)
		}
	}
	Ch, err := GetUserChannel(ID)
	if err.Error() == "" {
		bbb.dg.ChannelMessageSend(Ch.ID, ToLetter(Session)+">open")
	}
	response := GFS(Session)
	_, C, M, Info := DECOMP(response)
	if C == '>' && M == "opened" {
		GetBotInfo := &BotInfo{}
		json.Unmarshal([]byte(Info), GetBotInfo)
		alreadythere := false
		for _, Bot := range bbb.SeenBots {
			if Bot.BotID == GetBotInfo.BotID {
				alreadythere = true
			}
		}
		if !alreadythere {
			addBot(GetBotInfo)
		}
		bbb.Sessions[ID].Dailed = true
		bbb.Sessions[ID].IsOpener = true
		bbb.Sessions[ID].OwnSession = SB
		bbb.Sessions[ID].User, _ = GetUser(ID)
		bbb.Sessions[ID].IsBot = bbb.Sessions[ID].User.Bot
		bbb.Sessions[ID].Dail.Lastfailed = false
		bbb.Sessions[ID].Dail.Timeout = 360
		bbb.Sessions[ID].Dail.LastActivity = 0
		go DailLoop(ID, SB)
		return true
	}
	return false
}

func CALLDIAL(User *discordgo.User, Channel string, OtherSession byte) {
	Ch := GetFreeCh()
	var SS = false
	if OtherSession == byte(Ch) {
		SS = true
	}
	var NewSession = &Session{
		Dailed:       true,
		User:         User,
		IsOpener:     false,
		IsBot:        User.Bot,
		OwnSession:   byte(Ch),
		OtherSession: OtherSession,
		SessionSync:  SS,
	}
	bbb.Sessions[User.ID] = NewSession
	bbb.dg.ChannelMessageSend(Channel, string(OtherSession)+">opened "+Compile(bbb.OwnInfo))
}
