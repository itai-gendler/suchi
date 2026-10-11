//go:build !darwin && !linux

// SPDX-License-Identifier: AGPL-3.0-or-later

package fswatch

import "io/fs"

func sameFilesystem(_, _ fs.FileInfo) bool { return true }
