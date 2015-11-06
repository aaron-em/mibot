package channels

import (
	"github.com/nyubis/mibot/ircmessage"
	"github.com/nyubis/mibot/core"
	"github.com/nyubis/mibot/utils"
	"github.com/nyubis/mibot/modules/admin"
	"github.com/nyubis/mibot/modules/floodcontrol"
)

var autojoin []string
var blacklist []string
var bot *core.Bot

func Init(ircbot *core.Bot) {
	bot = ircbot

	LoadCfg()
}

func LoadCfg() {
	autojoin = core.Config.Channels.Autojoin
	blacklist = core.Config.Channels.Blacklist
}

func Autojoin(msg ircmessage.Message) {
	for _, channel := range autojoin {
		bot.SendJoin(channel)
	}
}

func InviteJoin(msg ircmessage.Message) {
	if len(msg.Content) > 0 && verify_channel(msg.Content) {
		if !floodcontrol.FloodCheck("invite", msg.Nick, msg.Channel) {
			bot.SendJoin(msg.Content)
		}
	}
}

func HandleJoinCommand(channels []string, sender string, fromchannel string) {
	if admin.CheckAdmin(sender) {
		for _, channel := range channels {
			if verify_channel(channel) {
				bot.SendJoin(channel)
			}
		}
	}
}

func verify_channel(channel string) bool {
	return channel[0] == '#' && !utils.Contains(blacklist, channel)
}
