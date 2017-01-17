package Belt

func QUERY(CMD string, ARGS []string, MASTER bool, ID string) string {
	if MASTER {
		bbb.dg.ChannelMessageSend(MasterInfo.DMChannel)
	} else {

	}
}
