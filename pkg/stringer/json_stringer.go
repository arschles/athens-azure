package stringer

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ericchiang/k8s"
)

// ToJSON encodes i to json and returns a string of the encoded JSON.
//
// If the encoding failed, it returns a string that indicates there was an error
func ToJSON(i k8s.Resource, name string) string {
	b, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf("error marshaling Deployment %s", name)
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, b, "", "    "); err != nil {
		return fmt.Sprintf("error indenting JSON for job %s", name)
	}
	return string(buf.Bytes())
}
