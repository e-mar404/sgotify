package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use: "list",
		Short: "list various objects from spotify",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("listing things")
		},
	}
)

func init() {
	RootCmd.AddCommand(devicesCmd)
}


