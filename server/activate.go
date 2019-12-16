package main

// OnActivate implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
func (p *LookGirlPlugin) OnActivate() error {
	return p.API.RegisterCommand(getCommand())
}

// OnDeactivate implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
func (p *LookGirlPlugin) OnDeactivate() error {

	return nil
}
