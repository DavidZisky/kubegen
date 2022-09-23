package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kubegen",
	Long:  `All software has versions. This is kubegens's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kubegen Kubernetes Manifest Generator v0.0.1 -- HEAD")
	},
}
