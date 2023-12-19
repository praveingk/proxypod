// Copyright 2023 The ClusterLink Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/google/uuid"
)

const (
	k8sTemplate = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.name}}
  labels:
    app: {{.name}}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{.name}}
  template:
    metadata:
      labels:
        app: {{.name}}
    spec:
      containers:
        - name: {{.name}}
          image: praveingk/proxypod:latest
          imagePullPolicy: Always
          command: ["/bin/sh"]
          args:
            - -c
            - >-
                /usr/local/bin/proxypod  --port {{.port}} \
                          --target {{.target}} &&
                sleep infinity`
)

// K8SConfig returns a kubernetes deployment file.
func k8SConfig(name string, port string, target string) ([]byte, error) {

	args := map[string]interface{}{
		"name":   name,
		"port":   port,
		"target": target,
	}

	var k8sConfig bytes.Buffer
	t := template.Must(template.New("").Parse(k8sTemplate))
	if err := t.Execute(&k8sConfig, args); err != nil {
		return nil, fmt.Errorf("cannot create k8s configuration off template: %w", err)
	}

	return k8sConfig.Bytes(), nil
}

func main() {
	name := flag.String("name", "", "Name of the proxy pod")
	listenPort := flag.String("port", "", "Local port to listen to")
	target := flag.String("target", "", "target service to redirect connections to")
	flag.Parse()
	yaml, _ := k8SConfig(*name, *listenPort, *target)
	fileString := "/tmp/" + *name + "-" + uuid.New().String()[:5] + ".yaml"
	os.WriteFile(fileString, yaml, 0600)
	cmd := exec.Command("kubectl", "apply", "-f", fileString)
	cmd.Run()
}
