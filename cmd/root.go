package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	// cfgFile string
	// userLicense string

	rootCmd = &cobra.Command{
		Use:   "rtc",
		Short: "Run the command",
		Long: `RTC is a CLI library for Go that will run a repeatedly command
		and group the execution time by status code.track the status code.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().IntP("asdf", "n", 10, "the number of times to run the command")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "Adam Israel <adam@adamisrael.com>")
	// viper.SetDefault("license", "apache")

	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
}

// func initConfig() {
// 	// if cfgFile != "" {
// 	// 	// Use config file from the flag.
// 	// 	viper.SetConfigFile(cfgFile)
// 	// } else {
// 	// 	// Find home directory.
// 	// 	home, err := os.UserHomeDir()
// 	// 	cobra.CheckErr(err)

// 	// 	// Search config in home directory with name ".cobra" (without extension).
// 	// 	viper.AddConfigPath(home)
// 	// 	viper.SetConfigType("yaml")
// 	// 	viper.SetConfigName(".cobra")
// 	// }

// 	// viper.AutomaticEnv()

// 	// if err := viper.ReadInConfig(); err == nil {
// 	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	// }
// }
