# Sgotify
*Spotify client for neovim writen in go*
---

## Using charm

how can i utilize the capabilities of charm in a better way?

I could try to make every command just pass directly to the charm main model and
then that will create any additional views and execute the proper code for what 
needs to be done??

This would require me write the command struct conform to the msg interface but 
that shouldnt be hard just have to look at docs, that way it makes the actuall
callback of what i want to execute (in the case of me wanting to start a server
or something) easy and still keep each model view to display what I want to see
instead of the command executing?

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
