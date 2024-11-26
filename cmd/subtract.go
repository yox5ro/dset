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

// subtractCmd represents the subtract command
var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "Perform set subtraction on lexicographically sorted text files.",
	Long: `Usage:
deet subtract [file1] [file2]
	
Example:
deet subtract file1.txt file2.txt
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Fprintf(os.Stderr, "Specify two files to subtract, got %d files\n", len(args))
			os.Exit(1)
		}
		if err := internal.SubtractWrapper(os.Stdout, args[0], args[1]); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subtractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subtractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
