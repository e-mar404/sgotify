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

- [x] sgotify help
    - using the default help menu for now
    - at some point use formatted output with [lipgloss](https://github.com/charmbracelet/lipgloss) of all 
    available commands and what they do

- [x] sgotify login
    - will start login process with http server
    - save all the necessary fields to the config
    - flag `-v` will show logs, without debug

- [x] sgotify logout
    - remove auth & refresh tokens, client id and secret and default device id

- [x] sgotify player
    - flag `--list-devices | -l` will list available devices
    - flag `--set-device | -s` will set a device defined by flag`

- [x] sgotify play 
    - if no arg then continue playing wtv was playing before on the set
      player
    - takes in a spotify uri and then plays it on the set player

- [x] sgotify pause
    - pause player

- [x] sgotify prev
    - go to the previously played song

- [x] sgotify next
    - go to the next song 

- [ ] sgotify search
    - search for things in spotify
    - needs the following flags:
        - album
        - artist
        - track
        - type (default to album,artist,playlist,track)

- [x] sgotify server
    - should start the rpc server
    - will have regular logging by default 
    - flag `-v` will show all the logs, including debug

## Default client

Ill be writing the progress of it in here since there is no clear outline yet.

11/10/2025
================================================================================

Wrote a few new commands to interact with what is currently playing. Having to
write commands that dont really interact with the server other than making a new
function to handle that new command it is showing some cracks on the design.

I need to make sure I go through the client and the services one more time
before finalizing the design because there is some uneccessary ceremony that I
have been doing to get some things to work.

After finalizing the design I can go back and write tests to make sure that any
new features I add later one dont break the minimum functionality I have as of
now.


11/08/2025
================================================================================

Implemented play cmd for resuming already playing media. So does not take any
args for the spotify uri yet. 

The /me/player/play endpoint is kind of janky because it says a successful
action will return 204 with no content but instead it return 200 with random
string. Had to add a specific check on the generic do[T]\(...\) function to just
return nil as the reply if conditions applied.

There is some error checking but I feel it could be better. That will be ironed
out with testing whenever I get to that.

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

