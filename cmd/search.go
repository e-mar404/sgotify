package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	album      string
	artist     string
	track      string
	sreachType []string

	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "search spotify for media",
		Run: func(cmd *cobra.Command, args []string) {
			var reply api.SearchReply
			searchArgs := api.SearchArgs{
				BaseURL:     viper.GetString("spotify_api_url"),
				AccessToken: viper.GetString("access_token"),
				Track:       track,
				Type:        sreachType,
			}

			err := client.Call("Search.Search", &searchArgs, &reply)
			if err != nil {
				fmt.Printf("unexpected error occurred\n")
				log.Fatal("unexpected error", "error", err)
			}

			printResults(reply)
		},
	}
)

func init() {
	defaultSearchType := []string{"album", "artist", "track"}

	searchCmd.Flags().StringVar(&album, "album", "", "filter down results to items with this album title")
	searchCmd.Flags().StringVar(&artist, "artist", "", "filter down results to items with from this artist")
	searchCmd.Flags().StringVar(&track, "track", "", "filter down results to items with this string on the track name")
	searchCmd.Flags().StringSliceVar(&sreachType, "type", defaultSearchType, "list of item types to search across")

	searchCmd.MarkFlagsOneRequired("album", "artist", "track")
	rootCmd.AddCommand(searchCmd)
}

func printResults(results api.SearchReply) {
	fmt.Printf("Tracks:\n\n")
	for _, track := range results.Tracks.Items {
		fmt.Printf("Name: %s\n", track.Name)

		artists := ""
		for _, artist := range track.Artists {
			artists += artist.Name + ", "
		}
		fmt.Printf("By: %s\n", artists)

		fmt.Printf("URI: %s\n\n", track.URI)
	}
}
