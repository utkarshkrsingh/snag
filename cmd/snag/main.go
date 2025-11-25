package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	human bool
)

var rootCmd = &cobra.Command{
	Use:   "snag",
	Short: "Snag is a very fast automation tool.",
	Long: `A fast and flexible way to hot-reload the project.
No need to manually build and run the project over and over
after modification.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Snag - Rewrite")
	},
}

func init() {
	// Remove auto-generated completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	timeCmd := &cobra.Command{
		Use:   "time",
		Short: "current time",
		Long:  "Fetch the current time from the operating system",

		Run: func(cmd *cobra.Command, args []string) {
			if human {
				fmt.Println("Time:", time.Now().Format("Monday, January 2, 2006 PM"))
			} else {
				fmt.Println("Time:", time.Now())
			}
		},
	}

	timeCmd.Flags().BoolVarP(&human, "human", "H", false, "Print time in human readable format")
	rootCmd.AddCommand(timeCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("error: %v\n", err.Error())
	}
}
