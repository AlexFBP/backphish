package common

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func FileRead(filePath string) (data []byte, err error) {
	return os.ReadFile(filePath)
}

func FileWrite(filePath string, data []byte) (err error) {
	if err = os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return
	}
	return os.WriteFile(filePath, data, 0644)
}

func YamlParse(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}

func YamlSerialize(v any) (out []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to serialize YAML: %v", r)
		}
	}()
	out, err = yaml.Marshal(v)
	return
}

func ReadYamlFile(filePath string, v any) (err error) {
	var data []byte
	if data, err = FileRead(filePath); err == nil {
		err = YamlParse(data, v)
	}
	return
}

func WriteYamlFile(filePath string, v any) (err error) {
	var data []byte
	if data, err = YamlSerialize(v); err != nil {
		return
	}
	return FileWrite(filePath, data)
}
