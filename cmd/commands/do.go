package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"web/cmd"
)

// helloCmd represents the serve command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do Something",
	Run: func(cmd *cobra.Command, args []string) {
		//Do the work here
		doWork()
		fmt.Print("Work done")
	},
}

func doWork() {
	fmt.Println("Processing ..")
}

func init() {
	cmd.RootCmd.AddCommand(doCmd)
}
