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
	"fmt"
	"strconv"

	"github.com/jasonblanchard/thecodelessctl/codeless"
	"github.com/jasonblanchard/thecodelessctl/store"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a single story at stdout",
	Long: `For a more enjoyably readin experience, pipe it to less:

	thecodelessctl read 42 | less
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var id int
		var err error
		bookmark, err := cmd.Flags().GetBool("bookmark")
		config, err := cmd.Flags().GetString("config")

		if err != nil {
			return err
		}

		if len(args) == 0 {
			id = store.GetBookmark()
		} else {
			idString := args[0]
			id, err = strconv.Atoi(idString)

			if err != nil {
				return err
			}
		}

		if bookmark == true {
			err := store.WriteBookmark(config, id)
			if err != nil {
				return err
			}
		}

		story := codeless.GetStoryById(id)
		fmt.Println(codeless.DecorateStory(story))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	readCmd.Flags().BoolP("bookmark", "b", false, "Set bookmark value to this story")
}
