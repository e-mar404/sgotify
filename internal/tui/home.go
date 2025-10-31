package tui

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	constants "github.com/e-mar404/sgotify/internal/const"
	"github.com/spf13/viper"
)

type HomeUI struct {
	loaded  bool
	help    help.Model
	profile api.ProfileReply 
}

type profileMsg api.ProfileReply 
type profileErrMsg error

func fetchProfileCmd() tea.Cmd {
	return func() tea.Msg {
		args := api.ProfileArgs {
			AccessToken: viper.GetString("access_token"),
			BaseUrl: viper.GetString("spotify_api_url"),
		}

		var reply api.ProfileReply 
		err := client.Call("User.Profile", &args, &reply)
		if err != nil {
			return profileErrMsg(err)
		}
		return profileMsg(reply)
	}
}

func (h HomeUI) Init() tea.Cmd {
	return fetchProfileCmd()
}

func (h HomeUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Quit):
			return h, tea.Quit
		case key.Matches(msg, constants.Keymap.Help):
			h.help.ShowAll = !h.help.ShowAll
		}

	case profileMsg:
		h.profile = api.ProfileReply(msg)
		h.loaded = true

	case profileErrMsg:
		log.Error("failed to fetch profile", "error", msg)
		h.profile = api.ProfileReply{} 
	}

	return h, nil
}

func (h HomeUI) View() string {
	if !h.loaded { 
		return "loading..."
	}

	borderOffset := 2
	helpView := h.help.View(constants.Keymap)

	// +1 accounts for the newline between content and helpView
	helpOffset := strings.Count(helpView, "\n") + 1
	
	var profileText string
	profileText += "username: " + h.profile.Username + "\n"
	profileText += "followers: " + strconv.Itoa(h.profile.Followers) + "\n"
	profileText += "top artist (last 30 days): " + h.profile.TopArtist + "\n"
	profileText += "top track (last 30 days): " + h.profile.TopTrack + "\n"

	content := constants.WindowStyle.
		Height(constants.WindowSize.Height - borderOffset - helpOffset).
		Width(constants.WindowSize.Width - borderOffset).
		Render(profileText)

	return content + "\n" + helpView
}

func newHomeUI() HomeUI {
	log.Debug("new home ui created")
	return HomeUI{
		help: help.New(),
	}
}
