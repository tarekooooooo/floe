package transfer

// macOS only on purpose (the _darwin suffix is the build constraint): the
// name in f_fstypename is how macOS names a file system.

import "golang.org/x/sys/unix"

// platformFSName names the file system under p, through the name statfs
// reports in f_fstypename.
func platformFSName(p string) (string, error) {
	var st unix.Statfs_t
	if err := unix.Statfs(p, &st); err != nil {
		return "", err
	}
	return unix.ByteSliceToString(st.Fstypename[:]), nil
}
