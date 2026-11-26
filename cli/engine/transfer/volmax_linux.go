package transfer

// Linux only on purpose (the _linux suffix is the build constraint): the
// magic number is how Linux names a file system.

import "golang.org/x/sys/unix"

// statfsMagic reports the magic number statfs gives for the file system
// under p: f_type of a unix.Statfs_t.
func statfsMagic(p string) (int64, error) {
	var st unix.Statfs_t
	if err := unix.Statfs(p, &st); err != nil {
		return 0, err
	}
	return st.Type, nil
}

// platformFSName names the file system under p, through the magic number
// statfs reports in f_type.
func platformFSName(p string) (string, error) {
	magic, err := statfsMagic(p)
	if err != nil {
		return "", err
	}
	return linuxFSName(magic), nil
}
