//go:build darwin || linux

// SPDX-License-Identifier: AGPL-3.0-or-later

package fswatch

import (
	"io/fs"
	"syscall"
)

func sameFilesystem(root, candidate fs.FileInfo) bool {
	rootStat, rootOK := root.Sys().(*syscall.Stat_t)
	candidateStat, candidateOK := candidate.Sys().(*syscall.Stat_t)
	return !rootOK || !candidateOK || rootStat.Dev == candidateStat.Dev
}
