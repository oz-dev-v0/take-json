package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var flags struct {
	filePath      string
	decodeURI     bool
	fallbackPrint bool
}

var flagsName = struct {
	file, fileShort           string
	decodeURI, decodeURIShort string
	fbPrint, fbPrintShort     string
}{
	"file", "f",
	"decodeURI", "d",
	"fallbackPrint", "p",
}

type takeJSON struct{}

func newTakeJSON() *takeJSON {
	return &takeJSON{}
}

// Run the command logic.
func Run() {
	c := &cobra.Command{
		Use:   "take-json",
		Short: "take json and pretty-print",
		Long:  "pretty-print json of linux pipes or file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return newTakeJSON().excute()
		},
	}

	c.Flags().StringVarP(
		&flags.filePath,
		flagsName.file,
		flagsName.fileShort,
		"", "path to the file.")

	c.PersistentFlags().BoolVarP(
		&flags.decodeURI,
		flagsName.decodeURI,
		flagsName.decodeURIShort,
		false, "print, decodeURI.")

	c.PersistentFlags().BoolVarP(
		&flags.fallbackPrint,
		flagsName.fbPrint,
		flagsName.fbPrintShort,
		false, "fallback, just input text print.")

	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
