package path_size

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// parseSizeString преобразует строковое представление размера в int64 (в байтах)
func parseSizeString(sizeStr string) (int64, error) {
	// Убираем пробелы и переводим в нижний регистр
	sizeStr = strings.TrimSpace(strings.ToLower(sizeStr))
	if len(sizeStr) == 0 {
		return 0, nil
	}

	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
		PB = 1024 * TB
		EB = 1024 * PB
	)

	var multiplier int64 = 1
	lastChar := sizeStr[len(sizeStr)-1]

	switch lastChar {
	case 'b':
		// Проверяем, есть ли перед 'b' ещё один символ (например, 'k', 'm' и т. д.)
		if len(sizeStr) > 1 {
			prevChar := sizeStr[len(sizeStr)-2]
			switch prevChar {
			case 'k':
				multiplier = KB
				sizeStr = sizeStr[:len(sizeStr)-2] // убираем 'kb'
			case 'm':
				multiplier = MB
				sizeStr = sizeStr[:len(sizeStr)-2] // убираем 'mb'
			case 'g':
				multiplier = GB
				sizeStr = sizeStr[:len(sizeStr)-2] // убираем 'gb'
			case 't':
				multiplier = TB
				sizeStr = sizeStr[:len(sizeStr)-2] // убираем 'tb'
			case 'p':
				multiplier = PB
				sizeStr = sizeStr[:len(sizeStr)-2] // убираем 'pb'
			case 'e':
				multiplier = EB
				sizeStr = sizeStr[:len(sizeStr)-2] // убираем 'eb'
			default:
				// Если перед 'b' не суффикс, то это просто байты — убираем только 'b'
				sizeStr = sizeStr[:len(sizeStr)-1]
			}
		} else {
			// Одиночный 'b' — байты, multiplier = 1, убираем 'b'
			sizeStr = sizeStr[:len(sizeStr)-1]
		}
	}

	// Если после удаления суффикса строка пустая, считаем её равной "0"
	if len(sizeStr) == 0 {
		sizeStr = "0"
	}

	num, err := strconv.ParseFloat(sizeStr, 64)
	if err != nil {
		return 0, fmt.Errorf("не удалось распарсить число '%s': %w", sizeStr, err)
	}

	return int64(num * float64(multiplier)), nil
}

// Проверяет, что функция GetPathSize корректно возвращает размер файла в байтах (1024 байт)
func TestGetPathSize_File_1KB(t *testing.T) {
	testFile := "testdata/file_1kb"
	expectedSizeBytes := int64(1024)

	sizeStr, err := GetPathSize(testFile, false)
	if err != nil {
		t.Errorf("Ошибка GetPathSize() для файла %s: %v", testFile, err)
	}

	actualSizeBytes, err := parseSizeString(sizeStr)
	if err != nil {
		t.Errorf("Не удалось распарсить строку размера '%s': %v", sizeStr, err)
	}

	if actualSizeBytes != expectedSizeBytes {
		t.Errorf("GetPathSize(%s) = %s (%d байт), ожидается %d байт",
			testFile, sizeStr, actualSizeBytes, expectedSizeBytes)
	}
}

// Проверяет, что функция GetPathSize корректно возвращает размер файла в байтах (2048 байт)
func TestGetPathSize_File_2KB(t *testing.T) {
	testFile := "testdata/file_2kb"
	expectedSizeBytes := int64(2048)

	sizeStr, err := GetPathSize(testFile, false)
	if err != nil {
		t.Errorf("Ошибка GetPathSize() для файла %s: %v", testFile, err)
	}

	actualSizeBytes, err := parseSizeString(sizeStr)
	if err != nil {
		t.Errorf("Не удалось распарсить строку размера '%s': %v", sizeStr, err)
	}

	if actualSizeBytes != expectedSizeBytes {
		t.Errorf("GetPathSize(%s) = %s (%d байт), ожидается %d байт",
			testFile, sizeStr, actualSizeBytes, expectedSizeBytes)
	}
}

// Проверяет, что функция GetPathSize корректно возвращает размер файла в байтах (1536 байт)
func TestGetPathSize_AnotherFile_1_5KB(t *testing.T) {
	testFile := "testdata/another_file"
	expectedSizeBytes := int64(1536)

	sizeStr, err := GetPathSize(testFile, false)
	if err != nil {
		t.Errorf("Ошибка GetPathSize() для файла %s: %v", testFile, err)
	}

	actualSizeBytes, err := parseSizeString(sizeStr)
	if err != nil {
		t.Errorf("Не удалось распарсить строку размера '%s': %v", sizeStr, err)
	}

	if actualSizeBytes != expectedSizeBytes {
		t.Errorf("GetPathSize(%s) = %s (%d байт), ожидается %d байт",
			testFile, sizeStr, actualSizeBytes, expectedSizeBytes)
	}
}

// Проверяет, что функция GetPathSize корректно возвращает размер файла в байтах (512 байт)
func TestGetPathSize_SmallFile_512B(t *testing.T) {
	testFile := "testdata/subdir/small_file"
	expectedSizeBytes := int64(512)

	sizeStr, err := GetPathSize(testFile, false)
	if err != nil {
		t.Errorf("Ошибка GetPathSize() для файла %s: %v", testFile, err)
	}

	actualSizeBytes, err := parseSizeString(sizeStr)
	if err != nil {
		t.Errorf("Не удалось распарсить строку размера '%s': %v", sizeStr, err)
	}

	if actualSizeBytes != expectedSizeBytes {
		t.Errorf("GetPathSize(%s) = %s (%d байт), ожидается %d байт",
			testFile, sizeStr, actualSizeBytes, expectedSizeBytes)
	}
}

// Проверяет, что функция GetPathSize корректно возвращает размер файлов
// первого уровня директории в байтах (4568 байт)
func TestGetPathSize_Directory_Root(t *testing.T) {
	testDir := "testdata"
	// Ожидаемый размер: только файлы первого уровня (file_1kb + file_2kb + another_file)
	// 1024 + 2048 + 1536 = 4568 байт ≈ 4.5K
	expectedSizeBytes := int64(4568)

	sizeStr, err := GetPathSize(testDir, false)
	if err != nil {
		t.Errorf("Ошибка GetPathSize() для директории %s: %v", testDir, err)
	}

	actualSizeBytes, err := parseSizeString(sizeStr)
	if err != nil {
		t.Errorf("Не удалось распарсить строку размера '%s': %v", sizeStr, err)
	}

	// Допускаем погрешность ±100 байт из‑за форматирования (например, 4.5K вместо точного 4568B)
	if actualSizeBytes < expectedSizeBytes-100 || actualSizeBytes > expectedSizeBytes+100 {
		t.Errorf("GetPathSize(%s) = %s (%d байт), ожидается приблизительно %d байт",
			testDir, sizeStr, actualSizeBytes, expectedSizeBytes)
	}
}

// Проверяет, что функция GetPathSize корректно возвращает размер файлов
// первого уровня поддиректории в байтах (512 байт)
func TestGetPathSize_Subdirectory(t *testing.T) {
	subDir := "testdata/subdir"
	expectedSizeBytes := int64(512)

	sizeStr, err := GetPathSize(subDir, false)
	if err != nil {
		t.Errorf("Ошибка GetPathSize() для поддиректории %s: %v", subDir, err)
	}

	actualSizeBytes, err := parseSizeString(sizeStr)
	if err != nil {
		t.Errorf("Не удалось распарсить строку размера '%s': %v", sizeStr, err)
	}

	if actualSizeBytes != expectedSizeBytes {
		t.Errorf("GetPathSize(%s) = %s (%d байт), ожидается %d байт",
			subDir, sizeStr, actualSizeBytes, expectedSizeBytes)
	}
}

// Проверяет, что функция GetPathSize корректно обрабатывает несуществующий путь,
// должна вернуть ошибку с конкретным сообщением
func TestGetPathSize_Nonexistent(t *testing.T) {
	nonexistent := "testdata/nonexistent_file"

	_, err := GetPathSize(nonexistent, false)
	if err == nil {
		t.Error("GetPathSize() должна возвращать ошибку для несуществующего пути")
	}

	if !strings.Contains(err.Error(), "файл или директория не существует") {
		t.Errorf("Ожидалась ошибка 'файл или директория не существует', получена: %v", err)
	}
}

// Проверяет, что функция GetPathSize корректно возвращает 0 байт для пустой директории
func TestGetPathSize_EmptyDirectory(t *testing.T) {
	emptyDir := "testdata/empty_dir"
	err := os.MkdirAll(emptyDir, 0755) //  создание пустой директории
	if err != nil {
		t.Fatalf("Не удалось создать пустую директорию: %v", err)
	}
	defer func() { //  настройка очистки после теста
		if err := os.RemoveAll(emptyDir); err != nil {
			t.Logf("Ошибка при удалении директории %s: %v", emptyDir, err)
		}
	}()

	sizeStr, err := GetPathSize(emptyDir, false)
	if err != nil {
		t.Errorf("Ошибка GetPathSize() для пустой директории: %v", err)
	}

	actualSizeBytes, err := parseSizeString(sizeStr)
	if err != nil {
		t.Errorf("Не удалось распарсить строку размера '%s': %v", sizeStr, err)
	}

	if actualSizeBytes != 0 {
		t.Errorf("GetPathSize(%s) = %s (%d байт), ожидается 0 байт", emptyDir, sizeStr, actualSizeBytes)
	}
}

// Проверяет, что функция GetPathSize корректно возвращает ошибку для директории без права доступа
func TestGetPathSize_ReadDirError(t *testing.T) {
	// Создаём директорию, из которой нельзя прочитать содержимое
	errorDir := "testdata/error_dir"
	err := os.MkdirAll(errorDir, 0000) // Права 0000 — нет доступа
	if err != nil {
		t.Fatalf("Не удалось создать директорию с ошибкой доступа: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(errorDir); err != nil {
			t.Logf("Ошибка при удалении директории %s: %v", errorDir, err)
		}
	}()

	_, err = GetPathSize(errorDir, false)
	if err == nil {
		t.Error("GetPathSize() должна возвращать ошибку, когда директорию невозможно прочитать")
	}

	if !strings.Contains(err.Error(), "ошибка чтения директории") {
		t.Errorf("Ожидалась ошибка 'ошибка чтения директории', получена: %v", err)
	}
}

// Проверяет, что функция GetPathSize корректно возвращает размер файла в байтах (1024 байт),
// в удобночитаемом формате
func TestGetPathSize_File_1KB_Human(t *testing.T) {
	testFile := "testdata/file_1kb"
	sizeStr, err := GetPathSize(testFile, true) // human = true
	if err != nil {
		t.Errorf("Ошибка GetPathSize() с флагом human для файла %s: %v", testFile, err)
	}

	// Проверяем, что результат содержит "KB" и не равен "0B"
	if !strings.Contains(sizeStr, "KB") || strings.Contains(sizeStr, "0B") {
		t.Errorf("Ожидался удобочитаемый формат (KB), получено: %s", sizeStr)
	}
}

// Проверяет, что функция GetPathSize корректно возвращает размер файлов первого
// уровня в байтах (4568 байт), в удобночитаемом формате
func TestGetPathSize_Directory_Root_Human(t *testing.T) {
	testDir := "testdata"
	sizeStr, err := GetPathSize(testDir, true) // human = true
	if err != nil {
		t.Errorf("Ошибка GetPathSize() с флагом human для директории %s: %v", testDir, err)
	}

	// Ожидаем "K" или "KB" для размера ~4.5 KB
	if !strings.Contains(sizeStr, "K") {
		t.Errorf("Ожидался удобочитаемый формат с 'K', получено: %s", sizeStr)
	}
}
