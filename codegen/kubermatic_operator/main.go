/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This application generates some manifests in the Kubermatic Helm chart
// based on the canonical source of truth, the Kubermatic Operator package.

package main

import (
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"

	"k8c.io/kubermatic/v2/pkg/controller/operator/defaults"
)

func main() {
	writeYAML(defaults.DefaultBackupStoreContainer, "charts/kubermatic/static/store-container.yaml")
	writeYAML(defaults.DefaultBackupCleanupContainer, "charts/kubermatic/static/cleanup-container.yaml")
	writeYAML(defaults.DefaultNewBackupStoreContainer, "charts/kubermatic/static/store-container-new.yaml")
	writeYAML(defaults.DefaultNewBackupDeleteContainer, "charts/kubermatic/static/delete-container.yaml")
	writeYAML(defaults.DefaultKubernetesAddons, "charts/kubermatic/static/master/kubernetes-addons.yaml")
	writeJSON(defaults.DefaultUIConfig, "charts/kubermatic/static/master/ui-config.json")

	markup, err := yaml.Marshal(map[string]interface{}{
		"addons": defaults.DefaultAccessibleAddons,
	})
	if err != nil {
		log.Fatalf("Failed to encode accessible addons as YAML: %v", err)
	}

	writeYAML(string(markup), "charts/kubermatic/static/master/accessible-addons.yaml")
}

func writeYAML(container string, filename string) {
	log.Printf("Generating %s...", filename)

	container = strings.Join([]string{
		"# This file has been generated using hack/update-kubermatic-chart.sh, do not edit.",
		"",
		strings.TrimSpace(container),
		"",
	}, "\n")

	if err := ioutil.WriteFile(filename, []byte(container), 0664); err != nil {
		log.Fatalf("Failed to write: %v", err)
	}
}

func writeJSON(data string, filename string) {
	log.Printf("Generating %s...", filename)

	data = strings.TrimSpace(data) + "\n"

	if err := ioutil.WriteFile(filename, []byte(data), 0664); err != nil {
		log.Fatalf("Failed to write: %v", err)
	}
}
