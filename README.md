### Hexlet tests and linter status:
[![Actions Status](https://github.com/VictoriscominG/go-from-scratch-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/VictoriscominG/go-from-scratch-project-242/actions)

### My tests and linter status:
[![Run Go Tests](https://github.com/VictoriscominG/go-from-scratch-project-242/actions/workflows/go-tests.yml/badge.svg)](https://github.com/VictoriscominG/go-from-scratch-project-242/actions/workflows/go-tests.yml)

### Демонстрация работы
Посмотрите интерактивную запись терминала с примерами использования:

[![Демонстрация GetPathSize](https://asciinema.org/a/k6suBotyo2UH9Z4L.svg)](https://asciinema.org/a/k6suBotyo2UH9Z4L)

### Примеры использования
* Базовый запуск:
  ```bash
  go run cmd/hexlet-path-size/main.go -h
  go run cmd/hexlet-path-size/main.go testdata/
  go run cmd/hexlet-path-size/main.go testdata/ --human
  go run cmd/hexlet-path-size/main.go testdata/ --human --all
  go run cmd/hexlet-path-size/main.go testdata/ --human --all --recursive
  go run cmd/hexlet-path-size/main.go testdata/file_1kb -H
  go run cmd/hexlet-path-size/main.go testdata/file_2kb -H
  go run cmd/hexlet-path-size/main.go testdata/another_file -H
  go run cmd/hexlet-path-size/main.go testdata/another_file
  go run cmd/hexlet-path-size/main.go testdata/.hidden_file_1kb -H
  go run cmd/hexlet-path-size/main.go testdata/.hidden_file_1kb -H -a
  make test