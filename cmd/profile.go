package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
)

const logo string = 
`⠀⠀⠀⠀⠀⠀⠀⢀⣠⣤⣤⣶⣶⣶⣶⣤⣤⣄⡀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢀⣤⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣤⡀⠀⠀⠀⠀
⠀⠀⠀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⠀⠀⠀
⠀⢀⣾⣿⡿⠿⠛⠛⠛⠉⠉⠉⠉⠛⠛⠛⠿⠿⣿⣿⣿⣿⣿⣷⡀⠀
⠀⣾⣿⣿⣇⠀⣀⣀⣠⣤⣤⣤⣤⣤⣀⣀⠀⠀⠀⠈⠙⠻⣿⣿⣷⠀
⢠⣿⣿⣿⣿⡿⠿⠟⠛⠛⠛⠛⠛⠛⠻⠿⢿⣿⣶⣤⣀⣠⣿⣿⣿⡄
⢸⣿⣿⣿⣿⣇⣀⣀⣤⣤⣤⣤⣤⣄⣀⣀⠀⠀⠉⠛⢿⣿⣿⣿⣿⡇
⠘⣿⣿⣿⣿⣿⠿⠿⠛⠛⠛⠛⠛⠛⠿⠿⣿⣶⣦⣤⣾⣿⣿⣿⣿⠃
⠀⢿⣿⣿⣿⣿⣤⣤⣤⣤⣶⣶⣦⣤⣤⣄⡀⠈⠙⣿⣿⣿⣿⣿⡿⠀
⠀⠈⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣾⣿⣿⣿⣿⡿⠁⠀
⠀⠀⠀⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠀⠀⠀
⠀⠀⠀⠀⠈⠛⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠛⠁⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠈⠙⠛⠛⠿⠿⠿⠿⠛⠛⠋⠁⠀⠀⠀⠀⠀⠀⠀`

var (
	statsStyle = lipgloss.NewStyle().Padding(2)
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("40"))
	profileCmd  = &cobra.Command {
		Use: "profile",
		Short: "display some stats about your spotify profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			profile, err := api.UserProfile()
			if err != nil {
				log.Error("could not get user profile", "error", err)
				return err
			}
			log.Info("got back a user profile", "profile", profile)

			topArtist, err := api.TopArtist()
			if err != nil {
				log.Error("could not get top artist", "error", err)
				return err
			}
			log.Info("got back top artist", "top item", topArtist)

			topTrack, err := api.TopTrack()
			if err != nil {
				log.Error("could not get top track", "error", err)
				return err
			}
			log.Info("got back top track", "top item", topTrack)
			
			// stats to display
			// ✓ username (maybe see if the pfp pic can be downloaded and show that instead of the logo)
			// ✓ follower count 
			// - playlist count (idk if ill do this one)
			// ✓ top artist (this month)
			// ✓ top track (this month)

			stats := strings.Builder{}
			stats.WriteString(fmt.Sprintf("Username: %s\n", profile.DisplayName))
			stats.WriteString(fmt.Sprintf("Followers: %d\n", profile.Followers.Total))
			stats.WriteString(fmt.Sprintf("Top Artist (this month): %s\n", topArtist.Items[0].Name))
			stats.WriteString(fmt.Sprintf("Top Track (this month): %s\n", topTrack.Items[0].Name))
			
			output := lipgloss.JoinHorizontal(
				lipgloss.Top,
				logoStyle.Render(logo),
				statsStyle.Render(stats.String()),
			)

			fmt.Println(output)

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(profileCmd)
}

