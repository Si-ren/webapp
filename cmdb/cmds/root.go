package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var rootCommand = &cobra.Command{
	Use:   "cmdb",
	Short: "cmdb program",
	Long:  "This is cmdb program",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("cmdb")
		return nil
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
