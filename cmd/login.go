// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "用户登陆",
	Long: 
	`
GoAgenda login -u username -p password

	各个参数分别对应:
	-u 用户名
	-p 用户密码
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("pass")
		// 处理参数
		fmt.Println("login called by ")
		fmt.Println("username: " + username)
		fmt.Println("password: " + password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("user", "u", "", "your username")
	loginCmd.Flags().StringP("pass", "p", "", "your password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
