//go:build !windows && !linux && !darwin

package transfer

// volumeMaxFileSize on every system but Windows, macOS and Linux answers 0,
// no known limit, as today.
func volumeMaxFileSize(dir string) (int64, error) { return 0, nil }



