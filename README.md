# Sgotify
*Spotify interface written in go*
---

## Vision

There will be 2 parts to the service, the server and the viewer (client). The
client and server will communicate with rpc-msgpack so the client can be written
to use in other places as well (ie. neovim).

### Commands to write

Will use viper for configuration

- [ ] sgotify 
    - will start the tui client
    - for now it is just a charm model that displays a "coming soon" msg

- [ ] sgotify help
    - formatted output with [lipgloss](https://github.com/charmbracelet/lipgloss) of all available commands and what they do

- [x] sgotify login
    - will start login process with http server
    - save all the necessary fields to the config
    - flag `-v` will show logs, without debug

- [x] sgotify server
    - should start the rpc server
    - will have regular logging by default 
    - flag `-v` will show all the logs, including debug

## Default client

Ill be writing the progress of it in here since there is no clear outline yet.

10/26/2025
================================================================================
Lets start with the home page showing a users profile. The profile will consist
of:

- User pfp (if kitty protocols are available) or the spotify logo
- username
- follower count
- top track (past 30 days)
- top artist (past 30 days)

