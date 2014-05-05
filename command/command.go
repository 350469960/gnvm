package command

import (
	// lib
	"github.com/spf13/cobra"

	// go
	"fmt"
	"strconv"
	"strings"

	// local
	"gnvm/config"
)

var (
	global bool
	remote bool
)

// defind root cmd
var gnvmCmd = &cobra.Command{
	Use:   "gnvm",
	Short: "Gnvm is Node.js version manage by GO on Win",
	Long: `Gnvm is Node.js version manage by GO on Win
           like nvm. Complete documentation is available at https://github.com/kenshin/gnvm`,
	Run: func(cmd *cobra.Command, args []string) {
		// TO DO
	},
}

// sub cmd
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gnvm",
	Long:  `Print the version number of Gnvm`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v" + config.VERSION)
	},
}

// sub cmd
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install any node.js version",
	Long: `install any node.js version like
		    'gnvm install latest'
		    'gnvm install x.xx.xx'
		    'gnvm install x.xx.xx --global`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gnvm install args include " + strings.Join(args, " "))
		fmt.Println("global flag is " + strconv.FormatBool(global))
		//TO DO
	},
}

// sub cmd
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall local node.js version",
	Long: `uninstall local node.js version like
		   'gnvm uninstall x.xx.xx`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gnvm uninstall args include " + strings.Join(args, " "))
		//TO DO
	},
}

// sub cmd
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "use the specific version by current cmd( temp )",
	Long: `use the specific version by current cmd( temp ) like
		   'gnvm use x.xx.xx
		   'gnvm use x.xx.xx --global`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gnvm use args include " + strings.Join(args, " "))
		fmt.Println("global flag is " + strconv.FormatBool(global))
		//TO DO
	},
}

// sub cmd
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update global node.js",
	Long: `update global node.js like
		   'gnvm update`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gnvm update args include " + strings.Join(args, " "))
		//TO DO
	},
}

// sub cmd
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list show all local | remote node.js version",
	Long: `list show all local | remote node.js version like
		   'gnvm ls
		   'gnvm ls --remote`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gnvm ls args include " + strings.Join(args, " "))
		fmt.Println("remote flag is " + strconv.FormatBool(remote))
		//TO DO
	},
}

// sub cmd
var nodeVersionCmd = &cobra.Command{
	Use:   "node-version",
	Short: "show global | current | latest node.js version",
	Long: `show global | current | latest node.js version like
		   'gnvm node-version
		   'laest version is x.xx.xx
		   'global version is x.xx.xx
		   'current version is x.xx.xx`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gnvm node-version args include " + strings.Join(args, " "))
		//TO DO
	},
}

func Exec() {

	// add sub cmd to root
	gnvmCmd.AddCommand(versionCmd)
	gnvmCmd.AddCommand(installCmd)
	gnvmCmd.AddCommand(uninstallCmd)
	gnvmCmd.AddCommand(useCmd)
	gnvmCmd.AddCommand(updateCmd)
	gnvmCmd.AddCommand(lsCmd)
	gnvmCmd.AddCommand(nodeVersionCmd)

	// flag
	gnvmCmd.PersistentFlags().BoolVarP(&global, "global", "g", false, "get this version global version")
	gnvmCmd.PersistentFlags().BoolVarP(&remote, "remote", "r", false, "get remote all node.js version list")

	// exec
	gnvmCmd.Execute()
}
