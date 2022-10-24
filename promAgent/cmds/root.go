package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	Verbose     bool
	Help        bool
	Path        string
	rootCommand = &cobra.Command{
		Use:   "promAgent",
		Short: "Prometheus Agent",
		Long:  "Prometheus Agent read config and upload to cmdb-test",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("This is cobra RunE")
			fmt.Println("Prom Agent")
			return nil
		},
	}
)

func init() {
	rootCommand.Flags().BoolVarP(&Verbose, "verbose", "v", false, "verbose details")
	rootCommand.Flags().StringVarP(&Path, "config", "c", "./etc/prometheus.yaml", "set prometheus config file path")
	//-h --help 参数与生俱来,可以不用定义,定义了也只是覆盖
	//rootCommand.Flags().BoolVarP(&Help, "help", "h", false, " Use \"go help <command>\" for more information about a command.")
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatalln(err)
	}
}
