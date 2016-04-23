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
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// sha512Cmd represents the sha512 command
var sha512Cmd = &cobra.Command{
	Use:   "sha512",
	Short: "Hashes the text provided into a SHA-512 password hash.",
	Long:  `Hashes the text provided into a SHA-512 password hash.`,
	Run: func(cmd *cobra.Command, args []string) {
		hash, err := GenerateSHA512FromString(strings.Join(args, " "))
		if err != nil {
			log.Panicln(err)
		}

		fmt.Println(hash)
	},
}

func init() {
	RootCmd.AddCommand(sha512Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sha512Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sha512Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
