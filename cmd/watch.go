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
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/typositoire/grocy-alerts/watcher"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:     "watch",
	Short:   "Fetch products expiering in the next `X` days and notify users throught `backend`.",
	Version: rootCmd.Version,
	Long: `Fetch products from grocy api and can notify users using multiple backend

Currently supported backend are:
- Sendgrid Dynamic Templates
- Stdout
- More to come.
`,
	Run: func(cmd *cobra.Command, args []string) {
		grocyUrl := viper.GetString("grocy_api-url")
		grocyApiKey := viper.GetString("grocy_api-key")
		sgApiKey := viper.GetString("sg_api-key")
		notifierBackend := viper.GetString("notifier-backend")

		w, err := watcher.New(grocyUrl+"/api", grocyApiKey, sgApiKey, notifierBackend)
		if err != nil {
			panic(err)
		}

		doEvery(time.Duration(viper.GetInt("interval")*int(time.Second)), w.Run)
	},
}

func doEvery(d time.Duration, f func()) {
	f()

	if d.Seconds() != -1 {
		for range time.Tick(d) {
			f()
		}
	}
}

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.PersistentFlags().Int("interval", -1, "Interval to check. If `-1`, will run once.")
	watchCmd.PersistentFlags().String("notifier-backend", "stdout", "Notifier backend.")
	watchCmd.PersistentFlags().String("grocy_api-url", "http://grocy.example.com", "Grocy url without `/api`.")
	watchCmd.PersistentFlags().String("grocy_api-key", "APIKEY", "Grocy API Key to gather products due soon.")
	watchCmd.PersistentFlags().String("grocy_due-soon-max", "5", "How far due to fetch.")
	watchCmd.PersistentFlags().String("sg_template-id", "d-dXXXXXXXXXXXXXXX", "Dynamic template ID for notification.")
	watchCmd.PersistentFlags().String("sg_from-email", "grocy-alerts@example.com", "Dynamic template ID for notification.")
	watchCmd.PersistentFlags().String("sg_api-key", "APIKEY", "Sendgrid api key.")
	watchCmd.PersistentFlags().String("sg_recipients", "", "Comma seperated list of recipient for sendgrid backend.")

	viper.BindPFlags(watchCmd.PersistentFlags())
	viper.BindPFlags(watchCmd.Flags())
}
