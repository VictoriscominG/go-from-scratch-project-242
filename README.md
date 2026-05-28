### Hexlet tests and linter status:
[![Actions Status](https://github.com/VictoriscominG/go-from-scratch-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/VictoriscominG/go-from-scratch-project-242/actions)

### My tests and linter status:
[![Run Go Tests](https://github.com/VictoriscominG/go-from-scratch-project-242/actions/workflows/go-tests.yml/badge.svg)](https://github.com/VictoriscominG/go-from-scratch-project-242/actions/workflows/go-tests.yml)

### Демонстрация работы
Посмотрите интерактивную запись терминала с примерами использования:

[![Демонстрация GetPathSize](https://asciinema.org/a/NEfYGp4g1bRoJS4B.svg)](https://asciinema.org/a/NEfYGp4g1bRoJS4B)

### Примеры использования
* Базовый запуск:
  ```bash
  go run cmd/hexlet-path-size/main.go -h
  go run cmd/hexlet-path-size/main.go .
  go run cmd/hexlet-path-size/main.go . --human
  go run cmd/hexlet-path-size/main.go . --human --recursive
  go run cmd/hexlet-path-size/main.go . --human --recursive --all
  go run cmd/hexlet-path-size/main.go Makefile -H
  go run cmd/hexlet-path-size/main.go testdata/ -H -r -a
  go run cmd/hexlet-path-size/main.go testdata/file_1kb -H
  go run cmd/hexlet-path-size/main.go testdata/file_2kb -H
  go run cmd/hexlet-path-size/main.go testdata/.hidden_file_1kb -H
  go run cmd/hexlet-path-size/main.go testdata/.hidden_file_1kb -H -a
  make test