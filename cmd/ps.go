/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ps called")
		ssh_ip_list := viper.GetStringSlice("host_ip_list")
		ssh_username := viper.GetString("host_username")
		ssh_password := viper.GetString("host_password")

		// var hostKey ssh.PublicKey
		config := &ssh.ClientConfig{
			User: ssh_username,
			Auth: []ssh.AuthMethod{
				ssh.Password(ssh_password),
			},
			// HostKeyCallback: ssh.FixedHostKey(hostKey),
		}
		config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
		if len(ssh_ip_list) > 0 {
			for _, v := range ssh_ip_list {
				fmt.Println(v)
				ssh_ip2 := v + ":" + "22"
				client, err := ssh.Dial("tcp", ssh_ip2, config)
				if err != nil {
					log.Fatal("Failed to dial: ", err)
				}
				defer client.Close()
				
				session, err := client.NewSession()
				if err != nil {
					log.Fatal("Failed to create session: ", err)
				}
				
			
				// Once a Session is created, you can execute a single command on
				// the remote side using the Run method.
				var b bytes.Buffer
				session.Stdout = &b
				if err := session.Run("docker ps -a"); err != nil {
					// log.Fatal("Failed to run: " + err.Error())
					log.Flags()
				}
				fmt.Println(b.String())
				session.Close()
			}
			os.Exit(0)


			// go func(hostname string, cr chan ConnectionResult) {
			// 	defer wg.Done()
			// 	_, err := net.Dial("tcp", host+":22")
			// 	if err != nil {
			// 		cr <- ConnectionResult{host, "failed"}
			// 	} else {
			// 		cr <- ConnectionResult{host, "succeeded"}
			// 	}
			// }(host, cnres)
		}
		fmt.Println("there is no ip address on the list")
	
	},
}

func init() {
	rootCmd.AddCommand(psCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psCmd.Flags.BoolVarP( "all", "a", false, "Get all containers")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
