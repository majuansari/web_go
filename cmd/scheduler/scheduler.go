package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"time"
	"web/cmd"
)

// schedulerCmd represents the scheduler command
var SchedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

var task = func() { fmt.Println("Running job") }

func Run() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Second().Do(task)
	s.StartBlocking()
}

func init() {
	cmd.RootCmd.AddCommand(SchedulerCmd)
}
