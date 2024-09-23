package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/adamisrael/rtc/internal/output"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Running %s\n", strings.Join(args, " "))
		// fmt.Printf("%v\n", args)

		command := args[0]
		args = args[1:]
		c := exec.Command(command, args...)

		// c.Stdout = os.Stderr
		// c.Stderr = os.Stderr

		// benchmark how long the command takes to run
		now := time.Now()
		err := c.Run()
		runtime := time.Since(now)

		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				fmt.Println(exitError.ExitCode())
			}
		} else {
			// 0
		}

		fmt.Printf("Runtime: %v\n", runtime)

		output.DisplayTable()
	},
}
