package transfer

// The rule behind the largest file a macOS or Linux volume holds: FAT and
// FAT32 only, and anything unrecognized has no limit, as diskfree_windows.go
// answers on Windows. Pure, with no build constraint, so it runs on every OS.

// fsMaxFileSize answers the largest file the named file system holds, or 0,
// no known limit, for any other name. "msdos" is the name macOS gives FAT and
// FAT32 and the one linuxFSName gives Linux's vfat and msdos mounts.
func fsMaxFileSize(name string) int64 {
	if name == "msdos" {
		return fat32MaxFileSize
	}
	return 0
}

// linuxFSName names a Linux file system by the magic number statfs reports
// in f_type: "msdos" for vfat and msdos mounts alike, "" for any other
// number.
func linuxFSName(fsType int64) string {
	if fsType == 0x4d44 {
		return "msdos"
	}
	return ""
}
