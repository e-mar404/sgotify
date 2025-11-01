# Sgotify
*Spotify interface written in go*
---

## Vision

There will be 3 parts to the service, the server, the cli, and the viewer 
(client). The client and the cli will be very similar, they will both be able to
do the same things since they will both be calling the same api functions it 
just differs on how they look. The client and server will communicate with 
rpc-msgpack so the client can be written to use in other places as well (ie. 
neovim).

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

- [x] sgotify logout
    - remove auth & refresh tokens, client id and secret and default device id

- [ ] sgotify player
    - flag `--list-devices | -l` will list available devices
    - flag `--set-device | -s` will set a device defined by flag `--deviceID=id`
    - [ ] play 
        - continue playing wtv was playing before
        - if song is given as an arg then it will search that song and play the
          first result with option to pass flag `--select` then you have to
          select the song out of the first few search results
    - [ ] pause: pause player
    - [ ] prev: go to the previously played song
    - [ ] next: go to the next song 

- [x] sgotify server
    - should start the rpc server
    - will have regular logging by default 
    - flag `-v` will show all the logs, including debug

## Default client

Ill be writing the progress of it in here since there is no clear outline yet.

11/01/2025
================================================================================

Refactoring in such a way that the cli is just a way to interact with the
server.

This means that I will have to create a client whenever the cli starts and. 

Also because I will be focusing on the cli aspect of it right now i will delete
the tui part of it, that can be done later.

10/26/2025
================================================================================
Lets start with the home page showing a users profile. The profile will consist
of:

- User pfp (if kitty protocols are available) or the spotify logo
- username
- follower count
- top track (past 30 days)
- top artist (past 30 days)

