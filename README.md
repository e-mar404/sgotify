# Sgotify
*Spotify tui writen in go*
---

Todo:

- [ ] figure out what how to structure the repo
    rn there are a few things going on:
        - there is the temporary auth server
        - the auth server also needs an http client
        - the auth server needs to save data that will live past its lifetime
          (and that it is cached somewhere securely)
        - whenever the auth token expires we need to be able to reauthenticate
          but this can be done with just a http client and not spin up the auth
          server again

### Idea of org

/cmd/sgotify/main.go
1. get env vars
2. create cfg
3. start app

/cmd/sgotify/app.go
1. check for auth token
2. start auth server if no auth token found
3. start ui?
...
more steps to come

/cmd/sgotify/auth.go
1. create new router here with cfg
2. do auth process and either store or refresh auth token

/pkg/router/router.go
1. create router which will really be a mux to use in auth.go

/internal/handlers/...
- write all the handlers needed for auth in here

### Commands to write

These are in the order that they should be written in

- [ ] sgotify login
    - check keyring for existance of clientID, clientSecret, auth token, and
      auth refresh token
    - use .env for clientID & clientSecret for initial load or prompt user for
      them

- [ ] sgotify logout
    - delete all fields from keyring

- [ ] sgotify profile
    - show some stats similar to neofetch

- [ ] sgotify
    - this will be the actual ui and music playback, will deal with that
      whenever the above is done
