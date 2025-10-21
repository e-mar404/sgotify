package set

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		Use:   "set",
		Short: "set default values for sgotify app",
	}
)

func init() {
	RootCmd.AddCommand(deviceCmd)
}
