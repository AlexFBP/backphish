package common

import (
	"errors"
)

func FileRead(filePath string) (data []byte, err error) {
	return nil, errors.New("FileRead: not implemented")
}

func FileWrite(filePath string, data []byte) (err error) {
	return errors.New("FileWrite: not implemented")
}

func YamlParse(data []byte, v any) error {
	return errors.New("YamlParse: not implemented")
}

func YamlSerialize(v any) (out []byte, err error) {
	return nil, errors.New("YamlSerialize: not implemented")
}

func ReadYamlFile(filePath string, v any) (err error) {
	return errors.New("ReadYamlFile: not implemented")
}

func WriteYamlFile(filePath string, v any) (err error) {
	return errors.New("WriteYamlFile: not implemented")
}
