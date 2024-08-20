package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
)

type Metadata struct {
	Name      string            `json:"name" yaml:"name"`
	Namespace string            `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Labels    map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}

type Secret struct {
	APIVersion string            `json:"apiVersion" yaml:"apiVersion"`
	Kind       string            `json:"kind" yaml:"kind"`
	Metadata   Metadata          `json:"metadata" yaml:"metadata"`
	Data       map[string]string `json:"data" yaml:"-"`
	StringData map[string]string `json:"stringData" yaml:"stringData,omitempty"`
	Type       string            `json:"type" yaml:"type"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: kubectl-decodesecret <secret-name>")
		os.Exit(1)
	}

	secretName := os.Args[1]

	// Run kubectl get secret command
	cmd := exec.Command("kubectl", "get", "secret", secretName, "-o", "json")
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running kubectl: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal the JSON output into a Secret struct
	var secret Secret
	if err := json.Unmarshal(output, &secret); err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	// Decode the base64 data and populate stringData
	secret.StringData = make(map[string]string)
	for key, value := range secret.Data {
		decodedValue, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding base64: %v\n", err)
			os.Exit(1)
		}
		secret.StringData[key] = string(decodedValue)
	}

	// Clear the original Data field to avoid printing base64 data
	secret.Data = nil

	// Marshal the Secret struct to YAML
	secretYAML, err := yaml.Marshal(&secret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding YAML: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(secretYAML))
}

