package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version bool
)

var RootCmd = &cobra.Command{
	Use:   "restful-api",
	Long:  "restful-api-cli 后端API1",
	Short: "restful-api-cli 后端API2",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("restful-api-cli v1.0.0")
		}
		return nil
	},
}

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&version, "version", "v", false, "restful-api 版本信息")
}
