package cmd

import "github.com/spf13/cobra"

var (
	album string
	artist string
	track string
	sreachType []string

	searchCmd = &cobra.Command{
		Use: "search",
		Short: "search spotify for media",
		Run: func(cmd *cobra.Command, args []string) {
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
