package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	annotationDeleteIDs []uint
)

var annotationDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete annotations",
	Run: func(cmd *cobra.Command, args []string) {
		c := connectGrafana()

		errFlag := false
		for _, id := range annotationDeleteIDs {
			err := c.AnnotationDelete(uint64(id))
			if err != nil {
				errFlag = true
				fmt.Printf("Fail to delete annotation %v due to %s", id, err.Error())
			}
		}
		if errFlag {
			os.Exit(1)
		}
	},
}

func init() {
	annotationDeleteCmd.Flags().UintSliceVar(&annotationDeleteIDs, "id", []uint{}, "Annotation IDs in format --id 1,2,3,... or --id 1 --id 2 ...")
	annotationDeleteCmd.MarkFlagRequired("id")

	annotationCmd.AddCommand(annotationDeleteCmd)
}
