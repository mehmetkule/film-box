package exxen

import "github.com/spf13/cobra"


func NewExxenCmd() *cobra.Command{
	rootCmd := &cobra.Command{
		Use:   "exxen",
		Short: "run exxen",
		RunE: exxen,
	}

	//Added Flags
	return rootCmd
}

func exxen(cmd *cobra.Command, args []string) error{
	//Do something
	return nil
}