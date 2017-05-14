package cmd

import "github.com/spf13/cobra"

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test",
	Long:  "Test",
	RunE:  test,
}

func init() {
	RootCmd.AddCommand(testCmd)
}

func test(cmd *cobra.Command, args []string) error {
	return nil
}
