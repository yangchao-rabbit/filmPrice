package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version bool

var rootCmd = &cobra.Command{
	Use:   "filmPrice",
	Short: "胶片比价后端",
	Long:  `胶片比价后端`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "版本信息")
}
