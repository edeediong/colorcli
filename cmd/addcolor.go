/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"io/ioutil"
	"github.com/spf13/cobra"
)

// addcolorCmd represents the addcolor command
var addcolorCmd = &cobra.Command{
	Use:   "addcolor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		addColor(args)
	},
}

func init() {
	rootCmd.AddCommand(addcolorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addcolorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addcolorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addColor(args []string) {

	hex := args[0]
	colorName := args[1]
	// Read the file
	content, err := ioutil.ReadFile("colornames.min.json")

	if err != nil {
		fmt.Printf("Error while reading a file %v", err)
	}

	// map hex value to color name
	var hexMap map[string]string

	// unmarshap to hexMap
	_ = json.Unmarshal(content, &hexMap)

	// check if hex exist
	name, ok := hexMap[hex]

	if ok {
		fmt.Printf("Hex already exist. Color Name is :%s\n", name)
	} else {
		hexMap[hex] = colorName

		//marshal to json or convert to json
		hexJSON, _ := json.Marshal(hexMap)

		// write to colornames.min.json
		err = ioutil.WriteFile("colornames.min.json", hexJSON, 0777)

		if err != nil {
			fmt.Printf("Error while writing a file %v", err)
		}

		fmt.Printf("\nHex to color added successfully!")
	}
}