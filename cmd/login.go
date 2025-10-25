package cmd

func loginHandler(_ command) error {
	// ask user for client id and secret
	// start http server for spotify auth
	// wait for code from spotify login
	// get access + refresh token
	// save everything config to ~/.config/sgotigy/conf.json
	return nil
}

func init() {
	availableCommands.AddCommand("login", loginHandler)
}
