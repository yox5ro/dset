/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/yox5ro/dset/internal"
)

// intersectCmd represents the intersect command
var intersectCmd = &cobra.Command{
	Use:   "intersect",
	Short: "Perform set intersection on lexicographically sorted text files.",
	Long: `Usage:
dset intersect [file1] [file2] [file3]...
	
Example:
dset intersect file1.txt file2.txt file3.txt
`,
	Run: func(cmd *cobra.Command, args []string) {
		files := make([]*os.File, len(args))
		for i, arg := range args {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			files[i] = file
			defer file.Close()
		}
		readers := make([]io.ReadSeeker, len(files))
		for i, file := range files {
			readers[i] = file
		}
		if err := internal.Intersect(os.Stdout, readers...); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(intersectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// intersectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// intersectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
