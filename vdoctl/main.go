/*
Copyright 2021.

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
package main

import (
	"github.com/vmware-tanzu/vsphere-kubernetes-drivers-operator/vdoctl/cmd"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "generate-doc" {
			cmd.GenerateMarkdownDoc(os.Args[2])
			os.Exit(0)
		}
	}

	cmd.Execute()
}
