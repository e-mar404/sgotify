package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	searchArgs api.SearchArgs

	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "search spotify for media",
		Run: func(cmd *cobra.Command, args []string) {
			var reply api.SearchReply
			searchArgs.BaseURL = viper.GetString("spotify_api_url")
			searchArgs.AccessToken = viper.GetString("access_token")

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

	searchCmd.Flags().StringVar(
		&searchArgs.Album,
		"album",
		"",
		"filter down results to items with this album title",
	)

	searchCmd.Flags().StringVar(
		&searchArgs.Artist,
		"artist",
		"",
		"filter down results to items with from this artist",
	)

	searchCmd.Flags().StringVar(
		&searchArgs.Track,
		"track",
		"",
		"filter down results to items with this string on the track name",
	)

	searchCmd.Flags().StringSliceVar(
		&searchArgs.Type,
		"type",
		defaultSearchType,
		"list of item types to search across",
	)

	searchCmd.MarkFlagsOneRequired("album", "artist", "track")

	rootCmd.AddCommand(searchCmd)
}

func printResults(results api.SearchReply) {
	printArtists(results.Artists.Items)
	printPlaylists(results.Playlist.Items)
	printAlbums(results.Albums.Items)
	printTracks(results.Tracks.Items)
}

func printTracks(tracks []api.TrackItem) {
	fmt.Printf("Tracks:\n\n")
	for _, track := range tracks {
		fmt.Printf("Name: %s\n", track.Name)

		artists := ""
		for _, artist := range track.Artists {
			artists += artist.Name + ", "
		}
		fmt.Printf("By: %s\n", artists)

		fmt.Printf("URI: %s\n\n", track.URI)
	}
}

func printArtists(artists []api.ArtistItem) {
	fmt.Printf("Artists:\n\n")
	for _, artist := range artists {
		fmt.Printf("Name: %s\n", artist.Name)
		fmt.Printf("URI: %s\n\n", artist.URI)
	}
}

func printPlaylists(playlists []api.PlaylistItem) {
	fmt.Printf("Playlists:\n\n")
	for _, playlist := range playlists {
		fmt.Printf("Name: %s\n", playlist.Name)
		fmt.Printf("URI: %s\n\n", playlist.URI)
	}
}

func printAlbums(albums []api.AlbumItem) {
	fmt.Printf("Albums:\n\n")
	for _, album := range albums {
		fmt.Printf("Name: %s\n", album.Name)

		artists := ""
		for _, artist := range album.Artists {
			artists += artist.Name + ", "
		}
		fmt.Printf("By: %s\n", artists)

		fmt.Printf("URI: %s\n\n", album.URI)
	}
}
