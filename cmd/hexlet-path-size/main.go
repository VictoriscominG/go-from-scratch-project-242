package main

import (
	path_size "code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args()
			if args.Len() == 0 {
				return fmt.Errorf("укажите путь к файлу или директории")
			}
			path := args.First()

			result, err := path_size.GetPathSize(path)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s", result, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
