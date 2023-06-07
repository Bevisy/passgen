/*
Copyright © 2023 Binbin Zhang binbin36520@gmail.com

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
	"math/rand"
	"strings"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generate a random password",
	Long:  `Generate a random password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(generatePassword())
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	randomCmd.Flags().IntP("length", "l", 12, "The length of the password")
}

func generatePassword() string {
	length := 12
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	charsByte := []byte(chars)
	var result strings.Builder

	for i := 0; i < length; i++ {
		result.WriteByte(charsByte[rand.Intn(len(charsByte))])
	}
	return result.String()
}
