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
	"log"

	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
		conn, err := infrastructure.RabbitConn()
		channel, _ := conn.Channel()
		if err != nil {
			panic(err)
		}
		defer channel.Close()

		// Start delivering queued messages.
		messages, err := channel.Consume(
			"TestQueue",
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Println(err)
		}

		// Welcome message.
		log.Println("Successfully connected to RabbitMQ instance")
		log.Println("[*] - Waiting for messages")
		log.Println("[*] - Run Fiber API server and go to http://127.0.0.1:3000/send?msg=<YOUR TEXT HERE>")

		// Open a channel to receive messages.
		forever := make(chan bool)

		go func() {
			for message := range messages {
				// For example, just show received message in console.
				log.Printf("Received message: %s\n", message.Body)
			}
		}()

		// Close the channel.
		<-forever
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
