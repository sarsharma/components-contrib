/*
Copyright 2021 The Dapr Authors
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

package bindings

import (
	"errors"

	"gopkg.in/yaml.v3"
)

// Metadata represents a set of binding specific properties.
type Metadata struct {
	Name       string
	Properties map[string]string `json:"properties"`
}

// SpecYAML is the data read from the spec.yaml file.
type SpecYAML []byte

// SpecMetada represents the entire metadata for a binding.
type SpecMedataData struct {
	Name                   string                 `json:"name" yaml:"name"`
	CertStatus             string                 `json:"cert-status" yaml:"cert-status"`
	Version                string                 `json:"version" yaml:"version"`
	BindingType            []string               `json:"binding-type" yaml:"binding-type"`
	SpecConnectionMetadata SpecConnectionMetadata `json:"connection-metadata" yaml:"connection-metadata"`
	Operations             Operations             `json:"operations" yaml:"operations"`
	InputBindingMetadata   InputBindingMetadata   `json:"input-binding-metadata" yaml:"input-binding-metadata"`
}

// SpecConnectionMetadata represents the collection of all possible connection metadata fields.
type SpecConnectionMetadata []SpecConnectionMetadataField

// SpecConnectionMetadataField represents details of a connection metadata field.
type SpecConnectionMetadataField struct {
	Name           string         `json:"name" yaml:"name"`
	BindingSupport BindingSupport `json:"binding-support" yaml:"binding-support"`
	Description    string         `json:"description" yaml:"description"`
	Example        string         `json:"example" yaml:"example"`
}

/* BindingSupport tells whether a SpecConnectionMetadataField is supported by the binding type and if it is optional or required.
Possible values: false/optional/required */
type BindingSupport struct {
	InputBinding  string `json:"input-binding" yaml:"input-binding"`
	OutputBinding string `json:"output-binding" yaml:"output-binding"`
}

// Operations represents the collection of metadata of all possible operations for a binding.
type Operations []OperationMetadata

// OperationMetadata contains the metadata of an operation for a binding.
type OperationMetadata struct {
	Name             string           `json:"name" yaml:"name"`
	Description      string           `json:"description" yaml:"description"`
	OperationInputs  OperationInputs  `json:"inputs" yaml:"inputs"`
	OperationOutputs OperationOutputs `json:"outputs" yaml:"outputs"`
}

// OperationInputs represents the details of all inputs for an operation.
type OperationInputs struct {
	Data     []SpecOperationInput `json:"data" yaml:"data"`
	Metadata []SpecOperationInput `json:"metadata" yaml:"metadata"`
}

// SpecOperationInput represents an input field for an operation.
type SpecOperationInput struct {
	Name        string `json:"name" yaml:"name"`
	Required    bool   `json:"required" yaml:"required"`
	Description string `json:"description" yaml:"description"`
}

// OperationOutputs represents outputs received as response from an operation.
type OperationOutputs struct {
	Data []ResponseMetadataField `json:"data" yaml:"data"`
}

// ResponseMetadataField represents the details of an output response of an operation.
type ResponseMetadataField struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}

// InputBindingMetadata represents the data received when the input binding is invoked.
type InputBindingMetadata struct {
	Data []ResponseMetadataField `json:"data" yaml:"data"`
}

// Method to unmarshal the SpecYAML read from spec.yaml file.
func (sp *SpecMedataData) UnmarshalYAML(sy SpecYAML) error {
	if len(sy) == 0 {
		return errors.New("unable to read spec metadata")
	}
	err := yaml.Unmarshal(sy, &sp)
	if err != nil {
		return errors.New("error in resolving spec metadata")
	}
	return nil
}
