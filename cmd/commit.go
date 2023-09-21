/*
Copyright © 2023 MihaK mihak09@gmail.com

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
	"strings"

	"github.com/mihakralj/opnsense/internal"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: `Commit changes from the 'staging.xml' to the active 'config.xml'`,
	Long: `The 'commit' command finalizes the staged changes made to the 'staging.xml' file, making them the active configuration for the OPNsense firewall system. This operation is the last step in a sequence that typically involves the 'set' and optionally 'discard' commands. The 'commit' action creates a backup of the active 'config.xml', moves 'staging.xml' to 'config.xml', and reloads the 'configd' service.
	`,

Example: `  opnsense commit          Commit the changes in 'staging.xml' to become the active 'config.xml'
  opnsense commit --force  Commit the changes without requiring interactive confirmation.`,
	Run: func(cmd *cobra.Command, args []string) {

		// check if staging.xml exists
		internal.Checkos()
		bash := `if [ -f "` + stagingfile + `" ]; then echo "exists"; fi`
		fileexists := internal.ExecuteCmd(bash, host)
		if strings.TrimSpace(fileexists) != "exists" {
			internal.Log(1, "no staging.xml detected - nothing to commit.")
		}
		internal.Log(2,"modifying "+configfile)

		// copy config.xml to /conf/backup dir
		backupname := generateBackupFilename()
		bash = `sudo cp -f ` + configfile + ` /conf/backup/` + backupname + ` && sudo mv -f /conf/staging.xml `+configfile
		//internal.ExecuteCmd(bash, host)
		fmt.Println(bash)

		bash = `if [ -f "` + configfile + `" ]; then echo "ok"; else echo "error"; fi`
		fileexists = internal.ExecuteCmd(bash, host)
		if fileexists == "ok" {
			bash = ``
		} else {
			//error
			bash = ``
		}
		// config reload - full or partial?

		fmt.Println(fileexists)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
