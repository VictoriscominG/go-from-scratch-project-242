package path_size

import (
	"code/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func GetPathSize(path string, human, all, recursive bool) (string, error) {
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

	// Проверяем путь на скрытность файла/директории
	isHidden, err := IsHidden(path)
	if err != nil {
		return "", fmt.Errorf("ошибка при проверке скрытности файла %s: %v", path, err)
	}

	// Проверяем тип на файл, выводим информацию учитывая флаги all | human
	if !info.IsDir() {
		if isHidden && !all {
			return "", fmt.Errorf("скрытый файл '%s' пропущен (используйте флаг -a для включения)", path)
		}
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
		if entry == nil {
			log.Printf("ПРОПУСКАЕМ: entry == nil в каталоге %s", path)
			continue
		}

		// Создаёт полный путь, автоматически выбирая разделитель (\ для Windows, / для Unix‑систем)
		entryPath := filepath.Join(path, entry.Name())

		// Проверяем путь на скрытность файла/директории
		isHiddenEntry, err := IsHidden(entryPath)
		if err != nil {
			log.Printf("ПРЕДУПРЕЖДЕНИЕ: ошибка проверки скрытости %s: %v", entryPath, err)
			continue
		}

		// Пропускаем скрытый файл/директорию при отсутствии флага all
		if isHiddenEntry && !all {
			continue
		}

		// Получаем информацию о файле, добавляем в вывод
		if !entry.IsDir() {
			entryInfo, err := entry.Info()
			if err != nil {
				log.Printf("ПРЕДУПРЕЖДЕНИЕ: не удалось получить информацию о %s: %v", entryPath, err)
				continue
			}
			totalSize += entryInfo.Size()
		}

		// Если директория и включен флаг рекурсивного обхода (recursive)
		if entry.IsDir() && recursive {
			// При рекурсии запрашиваем размер в байтах (human=false)
			subSizeStr, err := GetPathSize(entryPath, false, all, true)
			if err != nil {
				log.Printf("ПРЕДУПРЕЖДЕНИЕ: не удалось обработать поддиректорию %s: %v", entryPath, err)
				continue
			}

			// Преобразуем строчное кол-во байт в int64
			subSizeValue, err := strconv.ParseInt(subSizeStr, 10, 64)
			if err != nil {
				log.Printf("ПРЕДУПРЕЖДЕНИЕ: ошибка парсинга размера %s для %s: %v", subSizeStr, entryPath, err)
				continue
			}
			totalSize += subSizeValue
		}
	}

	if human {
		return formatSize(totalSize), nil
	} else {
		return fmt.Sprintf("%d", totalSize), nil
	}
}

// formatSize преобразует байты в удобочитаемый формат
func formatSize(bytes int64) string {
	switch {
	case bytes < config.KB:
		return fmt.Sprintf("%dB", bytes)
	case bytes < config.MB:
		return fmt.Sprintf("%.1fKB", float64(bytes)/config.KB)
	case bytes < config.GB:
		return fmt.Sprintf("%.1fMB", float64(bytes)/config.MB)
	case bytes < config.TB:
		return fmt.Sprintf("%.1fGB", float64(bytes)/config.GB)
	case bytes < config.PB:
		return fmt.Sprintf("%.1fTB", float64(bytes)/config.TB)
	case bytes < config.EB:
		return fmt.Sprintf("%.1fPB", float64(bytes)/config.PB)
	default:
		return fmt.Sprintf("%.1fEB", float64(bytes)/config.EB)
	}
}

// IsHidden проверяет, является ли файл скрытым в Unix‑системах
func IsHidden(path string) (bool, error) {
	filename := filepath.Base(path)
	if len(filename) == 0 {
		return false, nil
	}
	return filename[0] == '.', nil
}
