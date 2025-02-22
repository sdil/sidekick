/*
Copyright © 2024 Mahmoud Mosua <m.mousa@hey.com>

Licensed under the GNU AGPL License, Version 3.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
https://www.gnu.org/licenses/agpl-3.0.en.html

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mightymoud/sidekick/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Prepare env variable secrets by encrypting them before deployment",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ViperInit()
		envFile, envFileErr := os.ReadFile("./.env")
		if envFileErr != nil {
			pterm.Error.Println("Unable to process your env file")
		}
		pterm.Info.Println("Detected the following env variables in your project")
		for _, line := range strings.Split(string(envFile), "\n") {
			fmt.Println(strings.Split(line, "=")[0])
		}
		pterm.Info.Println("-----------")
		pterm.Info.Println("Encrypting env vars using VPS public key")
		envCmd := exec.Command("sh", "-s", "-", viper.Get("publicKey").(string))
		envCmd.Stdin = strings.NewReader(utils.EnvEncryptionScript)
		if envCmdErr := envCmd.Run(); envCmdErr != nil {
			panic(envCmdErr)
		}
	},
}

func init() {
	launchCmd.AddCommand(envCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// envCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// envCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
