package cmd

import (
	"calf-restful-api/version"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	vers bool
)

var RootCmd = &cobra.Command{
	Use:   "restful-api",
	Long:  "restful-api-cli 后端API1",
	Short: "restful-api-cli 后端API2",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println("restful-api-cli v1.0.0")
			fmt.Println(version.FullVersion())
		}
		return nil
	},
}

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&vers, "version", "v", false, "restful-api 版本信息")
}
