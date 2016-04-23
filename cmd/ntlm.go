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

// ntlmCmd represents the md4 command
var ntlmCmd = &cobra.Command{
	Use:   "ntlm",
	Short: "Generate MD4-based Windows NT Lan Manager Passwords",
	Long:  `Generate MD4-based Windows NT Lan Manager Passwords`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(GenerateMD4WindowsNTLMFromString(strings.Join(args, " ")))
	},
}

func init() {
	RootCmd.AddCommand(ntlmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ntlmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ntlmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
