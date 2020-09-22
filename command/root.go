package command

import "github.com/spf13/cobra"

func Execute() error {
	return root.Execute()
}

var root = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}
