package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// commandHelp is the text you see when you type /feed help
const commandHelp = `* |/feed subscribe url| - Connect your Mattermost channel to an rss feed 
 * |/feed list| - Lists the rss feeds you have subscribed to
 * |/feed unsubscribe url| - Unsubscribes the Mattermost channel from the rss feed`

// + `* |/feed initiate| - initiates the rss feed subscription poller`

const displayName = "妹子紳士"
const iconURL = "https://pngimg.com/uploads/trollface/trollface_PNG28.png"

func getCommand() *model.Command {
	return &model.Command{
		Trigger:     "看妹子",
		DisplayName: "LookGirls",
		Description: "看妹子圖的小工具",
	}
}

func getCommandResponse(responseType, text string) *model.CommandResponse {
	return &model.CommandResponse{
		ResponseType: responseType,
		Text:         text,
		Username:     displayName,
		IconURL:      iconURL,
		Type:         model.POST_DEFAULT,
	}
}

// ExecuteCommand will execute commands ...
func (p *LookGirlPlugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	split := strings.Fields(args.Command)
	command := split[0]

	if command != "/看妹子" {
		return &model.CommandResponse{}, nil
	}
	return getCommandResponse(model.COMMAND_RESPONSE_TYPE_EPHEMERAL, "![](https://api.ooopn.com/image/beauty/api.php?type=jump)"), nil
}
