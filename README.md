# Sgotify
*Spotify interface written in go*
---

## Vision

I want to create a way to interact with spotify whenever I am in the terminal
since I do almost everything through there. I was thinking of either making a
tui or a neovim plugin but then I decided to make a client agnostic rpc server
and then I can decide on which one I would like to implement. 

This project is divided in 2 stages:

1. Develop the api interface through the rpc server and cli commands
2. Create the "frontend" of the interface on either a tui or a neovim plugin

As of now the first stage is complete enough. I am able to control playback and
search spotify with the current cli interface. The second stage will be done
later.

## Installation

You can install the project with the following command:

```shell
go install github.com/e-mar404/sgotify
```

## Usage

### Initial setup

Before using the cli create a project in [spotify developers dashboard](https://developer.spotify.com/documentation/web-api/concepts/apps) and get a set of Client ID and Client Secret, 
take note of them since they will be used for the login process.

Note: make sure the redirect uri is `http://127.0.0.1:8080/callback`.

### Login

In order to use the cli or any future tuis the server needs to be running in the
background. To do that just run the command `sgotify server`. 

Now that the server is running we can start to log in and use the cli. Use the
login command, `sgotify login`, and add the client id and client secret from the
initial setup. Then go to the linked site to finalize log in with spotify.

### Playback

In order for playback to run you need to specify which device will be the target
of the playback. To do that list all the active players with `sgotify player
--list-devices` copy the device id of the device you want to use and then set
the device with `sgotify player --set-device <device-id>`.

Note: If you don't see the device you want on the list make sure the app is
open and running on that device.

## Commands 

The available commands are defined bellow
```
Usage:
  sgotify [flags]
  sgotify [command]

Available Commands:
  checkhealth verify that all the services and resources are in working condition
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  login       start login process
  logout      will remove any identifiable information from configuration
  next        skip to next track
  pause       Pause playback on current active device
  play        will start/resume playback on the set player
  player      command to interact with a spotify player state
  prev        go to the previous track
  search      search spotify for media
  server      start rpc server

Flags:
  -h, --help      help for sgotify
  -v, --verbose   set verbose output

Use "sgotify [command] --help" for more information about a command.
```
