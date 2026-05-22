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
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args().Slice() // получаем аргументы коммандной строки
			if len(args) == 0 {
				return fmt.Errorf("укажите путь к файлу или директории")
			}
			path := args[0] // берём первый аргумент (путь)

			human := cmd.Bool("human") // получаем значение флага human

			result, err := path_size.GetPathSize(path, human)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", result, path)
			return nil
		},
	}

	// Запускаем CLI‑команду с текущим контекстом и аргументами командной строки,
	// а при возникновении ошибки выводим её описание и аварийно завершаем программу
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
