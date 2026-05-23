package path_size

import (
	"fmt"
	"os"
)

func GetPathSize(path string, human bool) (string, error) {
	// Проверяем путь на существование, доступ и ошибки
	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("файл или директория не существует: %s", path)
		} else if os.IsPermission(err) {
			return "", fmt.Errorf("нет прав доступа к %s", path)
		} else {
			return "", fmt.Errorf("неизвестная ошибка при доступе к %s: %w", path, err)
		}
	}

	// Проверяем тип на файл/директорию
	if !info.IsDir() { // файл - отпраляем объём, nil
		if human {
			return formatSize(info.Size()), nil
		} else {
			return fmt.Sprintf("%d", info.Size()), nil
		}
	}

	// Если директория — суммируем размеры файлов первого уровня
	var totalSize int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения директории '%s': %w", path, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue // пропускаем поддиректории
		}

		entryInfo, err := entry.Info()
		if err != nil {
			continue // продолжаем обработку остальных файлов при ошибке с одним файлом
		}
		totalSize += entryInfo.Size()
	}
	if human {
		return formatSize(totalSize), nil
	} else {
		return fmt.Sprintf("%d", totalSize), nil
	}
}

// formatSize преобразует байты в удобочитаемый формат
func formatSize(bytes int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
		PB = 1024 * TB
		EB = 1024 * PB
	)

	switch {
	case bytes < KB:
		return fmt.Sprintf("%dB", bytes)
	case bytes < MB:
		return fmt.Sprintf("%.1fKB", float64(bytes)/KB)
	case bytes < GB:
		return fmt.Sprintf("%.1fMB", float64(bytes)/MB)
	case bytes < TB:
		return fmt.Sprintf("%.1fGB", float64(bytes)/GB)
	case bytes < PB:
		return fmt.Sprintf("%.1fTB", float64(bytes)/TB)
	default:
		return fmt.Sprintf("%.1fPB", float64(bytes)/PB)
	}
}
