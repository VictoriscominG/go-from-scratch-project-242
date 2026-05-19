package path_size

import (
	"fmt"
	"os"
)

func GetPathSize(path string) (string, error) {
	// Проверяем путь на существование и ошибки
	info, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("путь не существует: %s", path)
	} else if err != nil {
		return "", fmt.Errorf("ошибка доступа: %w", err)
	}

	// Проверяем тип на файл/директорию
	if !info.IsDir() {
		return formatSize(info.Size()), nil // Файл - отпраляем объём, nil
	}

	// Если директория — суммируем размеры файлов первого уровня
	var totalSize int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения директории '%s': %w", path, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue // Пропускаем поддиректории
		}

		entryInfo, err := entry.Info()
		if err != nil {
			continue // Продолжаем обработку остальных файлов при ошибке с одним файлом
		}
		totalSize += entryInfo.Size()
	}
	return formatSize(totalSize), nil
}

// formatSize преобразует байты в удобочитаемый формат
func formatSize(bytes int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)

	switch {
	case bytes < KB:
		return fmt.Sprintf("%dB", bytes)
	case bytes < MB:
		return fmt.Sprintf("%.1fK", float64(bytes)/KB)
	case bytes < GB:
		return fmt.Sprintf("%.1fM", float64(bytes)/MB)
	default:
		return fmt.Sprintf("%.1fG", float64(bytes)/GB)
	}
}
