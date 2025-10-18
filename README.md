# Sgotify
*Spotify client for neovim writen in go*
---

## Using charm

I have been thinking of the purpose of the tui in an entirely wrong way. I
should make the tui an extension of the cli. There needs to be a well defined
"api" and then the tui will just be the graphical way of doing it. I should not
be passing commands from the cli to the tui but instead passing keyboard input
to the cli and displaying the output. (by passing things to the cli i mean
passing the same data that would be used on the cli to the same api calls that
the cli uses.

### Commands to write

Now that I am using cobra and viper it should be a lot easier to write the
commands needed for the cli

This is the order that the very first commands should be written in:

- [ ] sgotify
    - will start the tui
    - this is first just because it will be the rootCmd for cobra
    - for now just do a charm model that will display a "coming soon" msg

- [ ] sgotify help
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
