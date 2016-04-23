// Copyright © 2016 Kevin Kirsche <kev.kirsche@gmail.com>
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
	"strings"

	"github.com/spf13/cobra"
)

// shake256Cmd represents the shake256 command
var shake256Cmd = &cobra.Command{
	Use:   "shake256",
	Short: "Generate a SHAKE256 variant SHA3 hash",
	Long:  `Generate a SHAKE256 variant SHA3 hash`,
	Run: func(cmd *cobra.Command, args []string) {
		passwordHash := GenerateSHA3ShakeSum256FromString(strings.Join(args, " "))
		fmt.Println(passwordHash)
	},
}

func init() {
	RootCmd.AddCommand(shake256Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shake256Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shake256Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
