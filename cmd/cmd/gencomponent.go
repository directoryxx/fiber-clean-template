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
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// gencomponentCmd represents the gencomponent command
var gencomponentCmd = &cobra.Command{
	Use:   "gencomponent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		genComponent(args, cmd)
	},
}

func init() {
	rootCmd.AddCommand(gencomponentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gencomponentCmd.PersistentFlags().String("foo", "", "A help for foo")
	gencomponentCmd.PersistentFlags().String("name", "", "Component Name")
	gencomponentCmd.PersistentFlags().String("target", "", "Target Name")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gencomponentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func genComponent(args []string, cmd *cobra.Command) {
	strName, errName := cmd.Flags().GetString("name")
	if errName != nil {
		fmt.Println("Name parameter is empty")
	}

	strTarget, errTarget := cmd.Flags().GetString("target")
	if errTarget != nil {
		fmt.Println("Target parameter is empty")
		return
	}

	_, errCheck := os.Stat("/app/app/schema/" + strName + ".go")
	if errCheck != nil {
		fmt.Println("Schema " + strName + " not found")
		return
	}

	if strTarget == "repository" || strTarget == "service" {
		// fmt.Println(strName)
		// fmt.Println(strTarget)
		newFile := copyFile(strName,strTarget)
		changePattern(newFile, strName)

	} else {
		fmt.Println("Target Unknown")
		return
	}
	// err := filepath.Walk("/app/cmd/stub/", visit)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

func copyFile(new string,target string) string {
	newFilePath := "/app/app/"+target+"/" + new + target +".go"
	sourceFile, err := os.Open("/app/cmd/stub/role"+target+".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create(newFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)

	return newFilePath
}

func changePattern(new string, schemaname string) {
	input, err := ioutil.ReadFile(new)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte("{Ubah}"), []byte(strings.Title(schemaname)), -1)

	if err = ioutil.WriteFile(new, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input1, err := ioutil.ReadFile(new)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output1 := bytes.Replace(input1, []byte("{Ubah1}"), []byte(schemaname), -1)

	if err = ioutil.WriteFile(new, output1, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
