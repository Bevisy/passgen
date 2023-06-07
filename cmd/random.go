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

var passwdLength int

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

	randomCmd.Flags().IntVarP(&passwdLength, "length", "l", 12, "The length of the password")
}

func generatePassword() string {
	length := passwdLength
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	charsByte := []byte(chars)
	var result strings.Builder

	for i := 0; i < length; i++ {
		result.WriteByte(charsByte[rand.Intn(len(charsByte))])
	}
	return result.String()
}
