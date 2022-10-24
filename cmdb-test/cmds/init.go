package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	db      string
	force   bool
	verbose bool
)

var initCommand = &cobra.Command{
	Use:   "init cmdb-test",
	Short: "init db and so on",
	Long:  "Init db and others before run cmdb-test",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(db, force, verbose)
		return nil
	},
}

func init() {
	rootCommand.AddCommand(initCommand)
	//go run main1.go init -d 123 -f -v,使用这些-f等长短参数
	initCommand.Flags().StringVarP(&db, "database", "d", "default", "database")
	initCommand.Flags().BoolVarP(&force, "force", "f", false, "force syncdb")
	//initCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose")
	//rootCommand.PersistentFlags为全局使用的,不会在子命令里使用
	rootCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose")
}
