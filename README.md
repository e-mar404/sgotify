# Sgotify
*Spotify client for neovim writen in go*
---

### Commands to write

This is the order that they should be written in too

- [x] sgotify help
    - formatted output with [lipgloss](https://github.com/charmbracelet/lipgloss) of all available commands and what they do

- [ ] sgotify login
    - use formated form with bubble tea to get client id & secret
    - during the same alt screen tui start log in flow with spotify 
    - use .env for clientID & clientSecret for initial load or prompt user for
      them

- [ ] sgotify logout
    - delete all fields from keyring

- [ ] sgotify profile
    - show some stats similar to neofetch

- [ ] sgotify start
    - will start rpc server and then this will connect to the neovim rpc client
      for the ui, this is the last thing that will be done
