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
			args := cmd.Args() // получаем аргументы коммандной строки
			if args.Len() == 0 {
				return fmt.Errorf("укажите путь к файлу или директории")
			}
			path := args.First() // берём первый аргумент(путь)

			result, err := path_size.GetPathSize(path)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", result, path)
			return nil
		},
	}
	// Запускаем CLI‑команду с текущим контекстом и аргументами командной строки, а при
	// возникновении ошибки выводим её описание и аварийно завершаем программу
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
