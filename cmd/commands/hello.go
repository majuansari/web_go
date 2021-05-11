package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"web/cmd"
)

var Name string

// helloCmd represents the serve command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello " + Name)
		//Do the work here
	},
}

func init() {
	cmd.RootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringVarP(&Name, "name", "n", "Tim", "Just give some name")
}
