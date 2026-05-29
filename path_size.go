package code

import (
	"code/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// sizeUnit описывает единицу измерения размера.
type sizeUnit struct {
	Threshold int64
	Suffix    string
}

// sizeUnits — пороги и суффиксы в порядке возрастания.
var sizeUnits = []sizeUnit{
	{config.EB, "EB"},
	{config.PB, "PB"},
	{config.TB, "TB"},
	{config.GB, "GB"},
	{config.MB, "MB"},
	{config.KB, "KB"},
}

// GetPathSize возвращает размер файла или директории.
func GetPathSize(path string, human, all, recursive bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			return "", fmt.Errorf("файл или директория не существует: %s", path)
		case os.IsPermission(err):
			return "", fmt.Errorf("нет прав доступа к %s", path)
		default:
			return "", fmt.Errorf("неизвестная ошибка при доступе к %s: %w", path, err)
		}
	}

	// Проверка скрытости для начального пути
	isHidden, err := IsHidden(path)
	if err != nil {
		return "", fmt.Errorf("ошибка при проверке скрытности файла %s: %v", path, err)
	}
	if isHidden && !all {
		return "", fmt.Errorf("скрытый файл '%s' пропущен (используйте флаг -a для включения)", path)
	}

	// Если это файл — сразу возвращаем размер
	if !info.IsDir() {
		return formatResult(info.Size(), human), nil
	}

	// Директория — суммируем размеры содержимого
	total, err := walkDir(path, all, recursive)
	if err != nil {
		return "", err
	}
	return formatResult(total, human), nil
}

// walkDir рекурсивно обходит директорию и возвращает суммарный размер в байтах.
func walkDir(path string, all, recursive bool) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения директории '%s': %w", path, err)
	}

	var total int64
	for _, entry := range entries {
		if isHiddenName(entry.Name()) && !all {
			continue
		}

		entryPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			if recursive {
				subSize, err := walkDir(entryPath, all, true)
				if err != nil {
					log.Printf("ПРЕДУПРЕЖДЕНИЕ: не удалось обработать поддиректорию %s: %v", entryPath, err)
					continue
				}
				total += subSize
			}
			continue
		}

		info, err := entry.Info()
		if err != nil {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: не удалось получить информацию о %s: %v", entryPath, err)
			continue
		}
		total += info.Size()
	}
	return total, nil
}

// formatResult возвращает размер в нужном формате.
func formatResult(bytes int64, human bool) string {
	if human {
		return formatSize(bytes)
	}
	return fmt.Sprintf("%dB", bytes)
}

// formatSize преобразует байты в удобочитаемый формат.
func formatSize(bytes int64) string {
	for _, u := range sizeUnits {
		if bytes >= u.Threshold {
			return fmt.Sprintf("%.1f%s", float64(bytes)/float64(u.Threshold), u.Suffix)
		}
	}
	return fmt.Sprintf("%dB", bytes)
}

// isHiddenName проверяет, начинается ли имя файла с точки.
func isHiddenName(name string) bool {
	return len(name) > 0 && name[0] == '.'
}

// IsHidden проверяет, является ли файл скрытым (Unix-системы).
func IsHidden(path string) (bool, error) {
	return isHiddenName(filepath.Base(path)), nil
}
