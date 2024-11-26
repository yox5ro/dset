/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yox5ro/dset/internal"
)

// unionCmd represents the union command
var unionCmd = &cobra.Command{
	Use:   "union",
	Short: "Perform set union on lexicographically sorted text files.",
	Long: `Usage:
dset union [file1] [file2] [file3]...

Example:
dset union file1.txt file2.txt file3.txt
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := internal.UnionWrapper(os.Stdout, args...); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(unionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
