# Sgotify
*Spotify tui writen in go*
---

### Commands to write

These are in the order that they should be written in

- [ ] sgotify login
    - check keyring for existence of clientID, clientSecret, auth token, and
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
