// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"github.com/clivern/beetle/app/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	// Expose build info to cmd subpackage to avoid custom ldflags
	cmd.Version = version
	cmd.Commit = commit
	cmd.Date = date
	cmd.BuiltBy = builtBy

	cmd.Execute()
}
