package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

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
	imageURL, code := getImangURL()

	if strings.Compare(code, "200") != 0 {
		return &model.CommandResponse{}, nil
	}
	return getCommandResponse(model.COMMAND_RESPONSE_TYPE_EPHEMERAL, "![]("+imageURL+")"), nil
}
