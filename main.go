package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
	"github.com/spf13/cobra"
)

var version string
var date string

func main() {
	cmd := &cobra.Command{
		Use: "uuid",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	var v int

	new := &cobra.Command{
		Use:     "new",
		Short:   "Generate a new UUID",
		Aliases: []string{"n"},
		RunE: func(cmd *cobra.Command, args []string) error {
			switch v {
			case 1:
				fmt.Print(uuid.Must(uuid.NewV1()))
			case 4:
				fmt.Print(uuid.Must(uuid.NewV4()))
			default:
				return errors.New("invalid version")
			}

			return nil
		},
	}

	new.Flags().IntVarP(&v, "version", "v", 4, "Speficy a UUID version as defined by RFC-4122[1] and DCE 1.1[2]")

	cmd.AddCommand(new)

	cmd.AddCommand(&cobra.Command{
		Use:     "version",
		Short:   "Print version information",
		Aliases: []string{"v"},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("uuid v%s (%s)\n", version, date)
			return nil
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:     "test [uuid]",
		Short:   "Test if a uuid is valid",
		Aliases: []string{"t"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := uuid.FromString(args[0])
			return err
		},
	})

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
