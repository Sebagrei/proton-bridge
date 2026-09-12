// Copyright (c) 2026 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package syncservice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSyncLimitsForSystem_ARM64MediumMemoryTune(t *testing.T) {
	limits := newSyncLimitsForSystem(2*Gigabyte, 3796*Megabyte, "arm64")

	require.Equal(t, uint64(1898*Megabyte), limits.MaxSyncMemory)
	require.Equal(t, uint64(64*Megabyte), limits.DownloadRequestMem)
	require.Equal(t, uint64(96*Megabyte), limits.MessageBuildMem)
	require.Equal(t, 16, limits.MaxParallelDownloads)
}

func TestNewSyncLimitsForSystem_NonARM64KeepsConservativePath(t *testing.T) {
	limits := newSyncLimitsForSystem(2*Gigabyte, 3796*Megabyte, "amd64")

	require.Equal(t, uint64(1898*Megabyte), limits.MaxSyncMemory)
	require.Equal(t, uint64(40*Megabyte), limits.DownloadRequestMem)
	require.Equal(t, uint64(64*Megabyte), limits.MessageBuildMem)
	require.Equal(t, 32, limits.MaxParallelDownloads)
}
