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
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var urlPath string
var staging bool
var admin bool

// loadallCmd represents the loadall command
var loadallCmd = &cobra.Command{
	Use:   "loadall",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand(urlPath, staging)

	},
}

func init() {
	rootCmd.AddCommand(loadallCmd)

	loadallCmd.Flags().StringVarP(&urlPath, "urlPath", "p", "", "URL Path to add to the end of URLs")
	loadallCmd.Flags().BoolVarP(&staging, "staging", "s", false, "Bool to do staging or not")
	loadallCmd.Flags().BoolVarP(&admin, "admin", "a", false, "Bool to do admin or not")
}

func getWebsiteURLS() []string {
	return []string{
		"https://www.afsrepair.com/",
		"https://www.aquaguard.net/",
		"https://www.bakerswaterproofing.com/",
		"https://www.completebasementsystems.net/",
		"https://www.dryprosystems.com/",
		"https://www.floridafoundationauthority.com/",
		"https://www.foundationrecoverysystems.com/",
		"https://www.foundationrepairwesterncolorado.com/",
		"https://www.drymich.com/",
		"https://www.dripfree.com/",
		"https://www.indianafoundation.com/",
		"https://www.innovativebasementauthority.com/",
		"https://www.itgbasements.com/",
		"https://www.fixmyfoundation.com/",
		"https://www.ohiobasementauthority.com/",
		"https://www.ohiobasementsystems.com/",
		"https://www.tarheelbasementsystems.com/",
		"https://staging.groundworkscompanies.com/",
	}
}

func runCommand(urlPath string, staging bool) {
	websiteURLS := getWebsiteURLS()

	for _, url := range websiteURLS {
		// Staging?
		if staging {
			url = strings.Replace(url, "www", "staging", 1)
		}

		// Admin?
		if admin {
			url = url + "wp-admin/"
		}

		// Add URL Path?
		if urlPath != "" {
			url = url + urlPath
		}

		// fmt.Println(url)

		// Load Websites
		err := exec.Command("open", url).Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}
