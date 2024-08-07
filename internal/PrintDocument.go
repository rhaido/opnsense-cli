/*
Copyright © 2023 Miha miha.kralj@outlook.com

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
package internal

import (
	"fmt"
	"os"

	"github.com/beevik/etree"
)

func PrintDocument(doc *etree.Document, path string) {
	var output string
	switch {
	case xmlFlag:
		output = ConfigToXML(doc, path)
	case jsonFlag:
		output = ConfigToJSON(doc, path)
	case yamlFlag:
		output = ConfigToYAML(doc, path)
	default:
		output = ConfigToTTY(doc, path)
	}

	if len(outfile) > 0 {
		f, err := os.Create(outfile)

		if err != nil {
			Log(1, "%v", err)
		}

		fmt.Fprintln(f, output)
	} else {
		fmt.Println(output)
	}
}
