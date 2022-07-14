// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package gotype

import "strings"

// tagOptions represents additional per field options that have
// been parsed from the struct tags.
//
// Available options are:
//   - squash, inline: The field must be a map[string]... or another struct.
//     All fields from the struct will be treated like they belong to the current
//     struct.
//   - omit: The field is never reported
//   - omitempty: The field is not reported if the value is assumed to be empty.
//     Strings, arrays, maps of 0 length and nil pointers are always assumed to be empty.
//     Custom types can implement the `IsZero() bool` method (IsZeroer interface). If the
//     `IsZero` method is true and `omitempty` has been set, the field will be ignored.
type tagOptions struct {
	squash    bool
	omitEmpty bool
	omit      bool
}

var defaultTagOptions = tagOptions{
	squash:    false,
	omitEmpty: false,
	omit:      false,
}

func parseTags(tag string) (string, tagOptions) {
	s := strings.Split(tag, ",")
	if len(s) == 0 {
		return "", defaultTagOptions
	}

	if s[0] == "-" {
		opts := defaultTagOptions
		opts.omit = true
		return "", opts
	}

	opts := defaultTagOptions
	for _, opt := range s[1:] {
		switch strings.TrimSpace(opt) {
		case "squash", "inline":
			opts.squash = true
		case "omitempty":
			opts.omitEmpty = true
		case "omit":
			opts.omit = true
		}
	}
	return strings.TrimSpace(s[0]), opts
}
