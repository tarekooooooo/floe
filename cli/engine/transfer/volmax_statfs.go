//go:build linux || darwin

package transfer

// The largest file a macOS or Linux volume holds, asked through statfs. The
// //go:build line above is the build constraint: statfs is not a GOOS suffix,
// so without it this file would build on Windows too, where volumeMaxFileSize
// already lives in diskfree_windows.go.

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// statfsName is the seam the tests replace: it names the file system of the
// folder it is asked about.
var statfsName = platformFSName

// volumeMaxFileSize reports the largest single file the volume under dir can
// hold, or 0 when the file system sets no limit this code knows of. Only FAT
// and FAT32 have one worth refusing on, as on Windows; anything unrecognized
// is treated as unlimited rather than guessed at.
//
// dir may not be made yet (the receive makes it after Decide), so a missing
// dir is asked through the nearest folder above it that exists.
func volumeMaxFileSize(dir string) (int64, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return 0, err
	}
	for {
		_, err := os.Stat(abs)
		if err == nil {
			break
		}
		if !errors.Is(err, fs.ErrNotExist) {
			return 0, err
		}
		parent := filepath.Dir(abs)
		if parent == abs {
			return 0, err
		}
		abs = parent
	}
	name, err := statfsName(abs)
	if err != nil {
		return 0, err
	}
	return fsMaxFileSize(name), nil
}
