/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os/exec"
	"unicode"

	"github.com/spf13/cobra"
)

var hasUpper bool

// generateschemaCmd represents the generateschema command
var generateschemaCmd = &cobra.Command{
	Use:   "generateschema",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateSchema(args, cmd)
	},
}

func init() {
	rootCmd.AddCommand(generateschemaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	generateschemaCmd.PersistentFlags().String("name", "", "Schema Name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateschemaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateSchema(args []string, cmd *cobra.Command) {
	str, err := cmd.Flags().GetString("name")
	if err != nil {
		panic(err)
	}

	hasUpper = false
	for _, r := range str {
		if unicode.IsUpper(r) {
			hasUpper = true
			break
		}
	}

	if !hasUpper {
		fmt.Println("String must capitalize")
		return
	}

	cmdExec := exec.Command("ent", "init", "--target", "app/schema", str)
	cmdExec.Dir = "/app"
	out, err := cmdExec.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
