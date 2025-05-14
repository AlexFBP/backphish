package common

import (
	"encoding/json"
	"os"
	"os/exec"
	// alternatives:
	// "github.com/robertkrimen/otto"
	// "github.com/dop251/goja"
)

// Runs a JS snippet. The snippet should only print to console a json serialized object.
// Uses installed nodejs to run the code.
func RunJS(jsCode string, result interface{}) (err error) {
	// Create a temporary js file
	var tmpFile *os.File
	if tmpFile, err = os.CreateTemp("", "bp-*.js"); err != nil {
		return err
	}
	close := func() {
		// no error checking as the file might had already been closed
		tmpFile.Close()
	}
	// Defer the closing of the file
	defer close()
	// Defer the removal of the file
	defer func() {
		if err = os.Remove(tmpFile.Name()); err != nil {
			return
		}
	}()

	// Write the js code to the file
	_, err = tmpFile.WriteString(missingJS() + jsCode)
	if err != nil {
		return
	}
	close()

	// eval the js code
	// -- prepare the command
	cmd := exec.Command("node", tmpFile.Name())
	// -- run the command
	var output []byte
	output, err = cmd.Output()
	if err != nil {
		return
	}

	// parse the output
	return json.Unmarshal(output, &result)
}

func missingJS() string {
	return `
function atob(encoded) {
    return Buffer.from(encoded, 'base64').toString('utf-8');
}
	`
}
