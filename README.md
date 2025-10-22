# Sgotify
*Spotify cli interface writen in go*
---

## Vision

The original plan was to make the cli its own stand alone spotify client, one to
which you can trasnfer music playback to, but due to some constraints imposed by
the spotify api it will have to be a little different.

The way this cli will work is that it will be based on prior existing devices 
that already have spotify open and running and this cli will just offer an 
interface to control those devices through the cli, and in the future through a 
tui.

### Commands to write

Commands and configuration is handled by cobra and viper due to its very easy
configuration and extensibility.

Since this is a project that is in development the commands will be outlined
here with notes on its state and any sub commands that will need to be written.
Once it is ready to be used a lot of this will get simplified.

- [ ] sgotify
    - will start the tui
    - for now it is just a charm model that displays a "coming soon" msg

- [ ] sgotify -h
    - formatted output with [lipgloss](https://github.com/charmbracelet/lipgloss) of all available commands and what they do
    - as of now it gets automatically created done with cobra but it is not 
      pretty

- [ ] sgotify login
    - will prompt for client id & secret if needed
    - extensibility ideas:
        - either take a --config flag that points to a file with client id &
          secret or have 2 flags --d & --secret to bypass initial prompt
    - updates/creates config file with fields needed for authentication

- [ ] sgotify logout
    - should delete anything related to a user's account:
        - access_token
        - refresh_token
        - device_name
        - device_id
        - client_id
        - client_secret

- [ ] sgotify profile
    - show some stats similar to neofetch
    - for now it has the following stats:
        - Username
        - Followers
        - Top Artist (in the past month)
        - Top Track (in the past month)

- [ ] sgotify list
    - fill list different objects from spotify
    - [ ] devices
        - list available devices, this info will be used to set the output
          device
    - [ ] playlists
        - list playlist in your library

- [ ] sgotify set
    - set different values used for playback
    - [ ] device \[param\]
        - will set the device passed to be the default audio output, can be
          either id or device name
    - [ ] volume \[param\]
        - sets the volume, 0-100 

- [ ] sgotify play
    - [ ] track
        - will play track with specified `--id` or `--name`
    - [ ] playlist
        - will play playlist with specified `--id` or `--name`
        - can be shuffled with the `--shuffle` / `-s` flag

## TODO

There are a few things that I already know need some improvements like logging
and how the program errors and making sure that the `requireAuth` function
actually gets ran every time a command that needs it has it in its `PreRun` 
field. As I find things like this they will be added in github issues and I will
create prs for them. A bit overkill but I find it easier to work on issues like
that instead of running a grep command to find all the "// TODO" comments
across my code.

### Using charm

I have been thinking of the purpose of the tui in an entirely wrong way. I
should make the tui an extension of the cli. There needs to be a well defined
"api" and then the tui will just be the graphical way of doing it. I should not
be passing commands from the cli to the tui but instead passing keyboard input
to the cli and displaying the output. (by passing things to the cli i mean
passing the same data that would be used on the cli to the same api calls that
the cli uses).

With that said I will not start on the tui until I am sure that I like the
design of the api/cli. After that creating the tui should not be too hard,
ideally.

