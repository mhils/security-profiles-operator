/*
Copyright 2024 The Kubernetes Authors.

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

package installer

import (
	"errors"

	ucli "github.com/urfave/cli/v2"
)

// Options define all possible options for the puller.
type Options struct {
	profilePath    string
	executablePath string
}

// Default returns a default options instance.
func Default() *Options {
	return &Options{
		profilePath: DefaultProfileFile,
	}
}

// FromContext can be used to create Options from an CLI context.
func FromContext(ctx *ucli.Context) (*Options, error) {
	options := Default()

	args := ctx.Args().Slice()
	if len(args) >= 1 {
		options.profilePath = args[0]
	}
	if len(args) >= 2 {
		options.executablePath = args[1]
	}
	if len(args) >= 3 {
		return nil, errors.New("too many arguments")
	}

	return options, nil
}
