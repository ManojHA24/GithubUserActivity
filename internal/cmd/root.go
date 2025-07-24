package cmd

import (
	"fmt"
	"os"

	"github.com/ManojHA24/GithubUserActivity/internal/issues"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "GithubUserActivity",
	Short: "A CLI tool to fetch GitHub user activity",
	Long:  `A Command Line Interface (CLI) tool to fetch and display GitHub user activity statistics.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.AutomaticEnv()
		viper.SetEnvPrefix("MY")
		token := viper.GetString("GITHUB_TOKEN")
		if token == "" {
			fmt.Print("🔑 Enter API key: ")
			fmt.Scanln(&token)
			viper.Set("GITHUB_TOKEN", token)
			if token == "" {
				return fmt.Errorf("api key is required")
			}
		}
		viper.Set("GITHUB_TOKEN", token)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func Init() {
	rootCmd.PersistentFlags().String("token", "", "GitHub Token (or set GITHUB_TOKEN)")
	viper.BindPFlag("GITHUB_TOKEN", rootCmd.PersistentFlags().Lookup("token"))

	rootCmd.AddCommand(commits.commitsCmd)
	rootCmd.AddCommand(issues.IssuesCmd)
	rootCmd.AddCommand(pullrequests.pullRequestsCmd)
	rootCmd.AddCommand(repos.reposCmd)
}
