// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/flyhard/versioner/version"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var output *string
var packageName *string
var skipGenerate *bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate a version file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		self := os.Args[0]
		w, err := os.Create(*output)
		if err != nil {
			log.Fatal("Failed to write to file: ", err)
			return
		}
		defer w.Close()
		fmt.Fprintf(w, "package %s\n\n", *packageName)
		if !*skipGenerate {
			fmt.Fprintf(w, "//go:generate %s generate --package=%s --output=%s\n", self, *packageName, *output)
		}
		fmt.Fprintf(w, "var (\n")
		fmt.Fprintf(w, "\tVersion\t\tstring = %q\n", version.GetVersion())
		fmt.Fprintf(w, "\tRevision\tstring = %q\n", version.GetRevision())
		fmt.Fprintf(w, "\tBuild\t\tstring = %q\n", time.Now().UTC().Format(time.RFC3339))
		fmt.Fprintf(w, ")\n")

	},
}

func init() {
	RootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	output = generateCmd.PersistentFlags().String("output", "version.go", "output filename")
	packageName = generateCmd.PersistentFlags().String("package", "main", "package for resulting go file")
	skipGenerate = generateCmd.PersistentFlags().Bool("skip-generate", false, "skip the //go:generate line in the file")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
