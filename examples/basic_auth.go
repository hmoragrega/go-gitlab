//
// Copyright 2021, Sander van Harmelen
//
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
//

package main

import (
	"log"

	"github.com/hmoragrega/go-gitlab"
)

// This example shows how to create a client with username and password.
func basicAuthExample() {
	git, err := gitlab.NewBasicAuthClient(
		"svanharmelen",
		"password",
		gitlab.WithBaseURL("https://gitlab.company.com"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// List all projects
	projects, _, err := git.Projects.ListProjects(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d projects", len(projects))
}
