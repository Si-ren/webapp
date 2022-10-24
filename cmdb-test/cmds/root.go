package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var rootCommand = &cobra.Command{
	Use:   "cmdb-test",
	Short: "cmdb-test program",
	Long:  "This is cmdb-test program",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("cmdb-test")
		return nil
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
