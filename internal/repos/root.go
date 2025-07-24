package repos

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ReposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Manage repositories",
	Long:  `Commands to manage and view repositories in a GitHub user account.`,
	RunE:  listRepos,
}

func listRepos(cmd *cobra.Command, args []string) error {
	// This function would typically fetch and list repositories for a user.
	// For demonstration, we will just print a message.
	fmt.Println("Listing repositories...")
	return nil
}

func init() {
	ReposCmd.Flags().String("user", "", "GitHub username to list repositories for")
	// reposCmd.MarkFlagRequired("user") // Make the user flag required
}
