package cmds

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/spf13/cobra"
)

var webCommand = &cobra.Command{
	Use:   "web",
	Short: "Web console",
	Long:  "This is Web console",
	RunE: func(cmd *cobra.Command, args []string) error {
		models.CacheInit("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)
		beego.Run()
		fmt.Println("Web console")
		return nil
	},
}

func init() {
	rootCommand.AddCommand(webCommand)
}
