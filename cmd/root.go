package main

import (
	"os"

	"github.com/mehmetkule/film-box/cmd/netflix"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {

	rootCmd := &cobra.Command{
		Use:   "film-box",
		Short: "Netflix,ve Exxen Popüler Film Listesi",
	}

	rootCmd.AddCommand(netflix.NetFlixCmd())
	return rootCmd
}

func main() {
	if err := NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
