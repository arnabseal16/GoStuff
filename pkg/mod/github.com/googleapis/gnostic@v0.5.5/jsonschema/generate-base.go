// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func write(f *os.File, s string) {
	_, err := f.WriteString(s)
	check(err)
}

func main() {
	f, err := os.Create("base.go")
	check(err)

	defer f.Close()

	write(f, `
// THIS FILE IS AUTOMATICALLY GENERATED.

package jsonschema

import (
	"encoding/base64"
)

func baseSchemaBytes() ([]byte, error){
	return base64.StdEncoding.DecodeString(
`)
	write(f, "`")

	b, err := ioutil.ReadFile("schema.json")
	check(err)

	s := base64.StdEncoding.EncodeToString(b)
	limit := len(s)
	width := 80
	for i := 0; i < limit; i += width {
		if i > 0 {
			write(f, "\n")
		}
		j := i + width
		if j > limit {
			j = limit
		}
		write(f, s[i:j])
	}
	write(f, "`)}")
}
