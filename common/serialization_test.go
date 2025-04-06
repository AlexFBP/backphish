package common_test

import (
	"os"
	"testing"

	"github.com/AlexFBP/backphish/common"
)

func TestFileRead(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write some data to the temporary file
	expectedContent := []byte("Hello, World!")
	if _, err := tempFile.Write(expectedContent); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	// Test FileRead
	data, err := common.FileRead(tempFile.Name())
	if err != nil {
		t.Errorf("FileRead returned an error: %v", err)
	}

	// Verify the content
	if string(data) != string(expectedContent) {
		t.Errorf("FileRead returned incorrect data. Got: %s, Want: %s", data, expectedContent)
	}
}

func TestFileRead_FileNotFound(t *testing.T) {
	// Test FileRead with a non-existent file
	_, err := common.FileRead("non_existent_file.txt")
	if err == nil {
		t.Error("FileRead should return an error for a non-existent file, but it did not")
	}
}

func TestFileWrite(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close() // Close the file so it can be written to later

	// Data to write
	content := []byte("Hello, FileWrite!")

	// Test FileWrite
	err = common.FileWrite(tempFile.Name(), content)
	if err != nil {
		t.Errorf("FileWrite returned an error: %v", err)
	}

	// Verify the content
	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read back the file: %v", err)
	}
	if string(data) != string(content) {
		t.Errorf("FileWrite wrote incorrect data. Got: %s, Want: %s", data, content)
	}
}

func TestFileWrite_CreateDirectories(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// File path with non-existent directories
	filePath := tempDir + "/nonexistent/subdir/testfile.txt"

	// Data to write
	content := []byte("Hello, Directories!")

	// Test FileWrite
	err = common.FileWrite(filePath, content)
	if err != nil {
		t.Errorf("FileWrite returned an error: %v", err)
	}

	// Verify the content
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read back the file: %v", err)
	}
	if string(data) != string(content) {
		t.Errorf("FileWrite wrote incorrect data. Got: %s, Want: %s", data, content)
	}
}

func TestFileWrite_MkdirAllFailure(t *testing.T) {
	// Create a file path with an invalid directory to simulate MkdirAll failure
	invalidDir := string([]byte{0}) // Invalid directory name
	filePath := invalidDir + "/testfile.txt"

	// Data to write
	content := []byte("This should fail")

	// Test FileWrite
	err := common.FileWrite(filePath, content)
	if err == nil {
		t.Error("FileWrite should return an error when MkdirAll fails, but it did not")
	}
}

func TestYamlParse(t *testing.T) {
	// Valid YAML data
	yamlData := []byte(`
key1: value1
key2: value2
nested:
  key3: value3
`)

	var result map[string]any

	// Test YamlParse with valid data
	err := common.YamlParse(yamlData, &result)
	if err != nil {
		t.Errorf("YamlParse returned an error: %v", err)
	}

	// Verify the parsed content
	expected := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"nested": map[string]any{
			"key3": "value3",
		},
	}
	if len(result) != len(expected) || result["key1"] != expected["key1"] || result["key2"] != expected["key2"] {
		t.Errorf("YamlParse returned incorrect data. Got: %v, Want: %v", result, expected)
	}

	// Invalid YAML data
	invalidYamlData := []byte(`
key1: value1
key2
`)

	// Test YamlParse with invalid data
	err = common.YamlParse(invalidYamlData, &result)
	if err == nil {
		t.Error("YamlParse should return an error for invalid YAML data, but it did not")
	}
}

func TestYamlSerialize(t *testing.T) {
	// Test YamlSerialize with valid input
	input := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"nested": map[string]any{
			"key3": "value3",
		},
	}

	expectedYaml := `key1: value1
key2: value2
nested:
    key3: value3
`

	// Test YamlSerialize
	output, err := common.YamlSerialize(input)
	if err != nil {
		t.Errorf("YamlSerialize returned an error: %v", err)
	}

	// Verify the serialized content
	if string(output) != expectedYaml {
		t.Errorf("YamlSerialize returned incorrect data.\nGot:\n%s\nWant:\n%s", output, expectedYaml)
	}

	// Test YamlSerialize with invalid input (e.g., a channel, which cannot be serialized)
	invalidInput := make(chan int)

	_, err = common.YamlSerialize(invalidInput)
	if err == nil {
		t.Error("YamlSerialize should return an error for invalid input, but it did not")
	}
}

func TestReadYamlFile(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile.yaml")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write valid YAML data to the temporary file
	yamlContent := []byte(`
key1: value1
key2: value2
nested:
  key3: value3
`)
	if _, err := tempFile.Write(yamlContent); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	// Test ReadYamlFile with valid YAML data
	var result map[string]any
	err = common.ReadYamlFile(tempFile.Name(), &result)
	if err != nil {
		t.Errorf("ReadYamlFile returned an error: %v", err)
	}

	// Verify the parsed content
	expected := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"nested": map[string]any{
			"key3": "value3",
		},
	}
	if len(result) != len(expected) || result["key1"] != expected["key1"] || result["key2"] != expected["key2"] {
		t.Errorf("ReadYamlFile returned incorrect data. Got: %v, Want: %v", result, expected)
	}

	// Test ReadYamlFile with a non-existent file
	err = common.ReadYamlFile("non_existent_file.yaml", &result)
	if err == nil {
		t.Error("ReadYamlFile should return an error for a non-existent file, but it did not")
	}

	// Test ReadYamlFile with invalid YAML data
	tempFile, err = os.CreateTemp("", "invalid_testfile.yaml")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	invalidYamlContent := []byte(`
key1: value1
key2
`)
	if _, err := tempFile.Write(invalidYamlContent); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	err = common.ReadYamlFile(tempFile.Name(), &result)
	if err == nil {
		t.Error("ReadYamlFile should return an error for invalid YAML data, but it did not")
	}
}

func TestWriteYamlFile(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile.yaml")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close() // Close the file so it can be written to later

	// Data to write
	input := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"nested": map[string]any{
			"key3": "value3",
		},
	}

	// Test WriteYamlFile
	err = common.WriteYamlFile(tempFile.Name(), input)
	if err != nil {
		t.Errorf("WriteYamlFile returned an error: %v", err)
	}

	// Verify the content
	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read back the file: %v", err)
	}

	expectedYaml := `key1: value1
key2: value2
nested:
    key3: value3
`

	if string(data) != expectedYaml {
		t.Errorf("WriteYamlFile wrote incorrect data.\nGot:\n%s\nWant:\n%s", data, expectedYaml)
	}

	// Test WriteYamlFile with invalid input (e.g., a channel, which cannot be serialized)
	invalidInput := make(chan int)

	err = common.WriteYamlFile(tempFile.Name(), invalidInput)
	if err == nil {
		t.Error("WriteYamlFile should return an error for invalid input, but it did not")
	}
}
