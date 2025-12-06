package wordlist

import (
	"bufio"
	"os"
	"strings"
)

// Load 加载字典，如果指定了文件路径则从文件加载，否则使用内置字典
func Load(filepath string) ([]string, error) {
	if filepath == "" {
		return BuiltinWordlist, nil
	}
	return LoadFromFile(filepath)
}

// LoadFromFile 从文件加载字典
func LoadFromFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" && !strings.HasPrefix(word, "#") {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// GetBuiltin 获取内置字典
func GetBuiltin() []string {
	return BuiltinWordlist
}
