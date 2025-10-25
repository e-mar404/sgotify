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
    - as of now it gets automatically created done with cobra but it is not 
      pretty

- [ ] sgotify server
    - should start the rpc server
    - will have regular logging by default 
    - flag `-v` will show the debug logs
