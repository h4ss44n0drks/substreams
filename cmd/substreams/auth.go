package main

import (
	"errors"
	"fmt"
	"github.com/streamingfast/cli"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:          "auth",
	Short:        "Login command for Substreams development",
	RunE:         runAuthE,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(authCmd)
}

func runAuthE(cmd *cobra.Command, args []string) error {
	localDevelopment := os.Getenv("LOCAL_DEVELOPMENT")

	fmt.Println("Open this link to authenticate on The Graph Market:")
	fmt.Println()
	if localDevelopment == "true" {
		fmt.Println("    " + cli.PurpleStyle.Render("http://localhost:3000/auth/substreams-devenv"))
	} else {
		fmt.Println("    " + cli.PurpleStyle.Render("https://thegraph.market/auth/substreams-devenv"))
	}
	fmt.Println("")

	var token string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				EchoMode(huh.EchoModePassword).
				Title("After retrieving your token, paste it here:").
				Inline(true).
				Value(&token).
				Validate(func(s string) error {
					if s == "" {
						return errors.New("token cannot be empty")
					}
					if strings.HasPrefix(s, "server_") {
						return errors.New("You've copied an API key, not a JWT token. Obtain a JWT from the link above, or by Generating an API Token in the The Graph Market Dashboard.")
					}
					return nil
				}),
		),
	)

	if err := form.Run(); err != nil {
		return fmt.Errorf("error running form: %w", err)
	}

	fmt.Println("Writing `./.substreams.env`.  NOTE: Add it to `.gitignore`.")
	fmt.Println("")

	err := os.WriteFile(".substreams.env", []byte(fmt.Sprintf("export SUBSTREAMS_API_TOKEN=%s\n", token)), 0644)
	if err != nil {
		return fmt.Errorf("writing .substreams.env file: %w", err)
	}

	fmt.Println("Load credentials in current terminal with the following command:")
	fmt.Println("")
	fmt.Println(cli.PurpleStyle.Render("       . ./.substreams.env"))
	fmt.Println()

	return nil
}
