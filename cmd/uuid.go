/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

var format bool

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a UUID",
	Long:  `Generate a UUID`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(generateUUID())
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)

	uuidCmd.Flags().BoolVarP(&format, "format", "f", false, "Format the UUID")
}

// generate UUID
func generateUUID() string {
	u := make([]byte, 16)
	rand.Read(u)
	uuid := hex.EncodeToString(u)
	if !format {
		return uuid
	} else {
		return fmt.Sprintf("%s-%s-%s-%s-%s", uuid[0:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:])
	}
}
