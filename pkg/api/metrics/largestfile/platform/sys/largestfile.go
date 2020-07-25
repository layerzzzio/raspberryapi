package sys

import (
	"strings"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
)

// LargestFile represents an empty LargestFile entity on the current system.
type LargestFile struct{}

// List returns a list of LargestFile stats
func (lf LargestFile) List(top100files []metrics.PathSize) ([]rpi.LargestFile, error) {
	var result []rpi.LargestFile

	for i := range top100files {
		fileCat := FileCategory(top100files[i].Path)
		result = append(result, rpi.LargestFile{
			Size:                top100files[i].Size,
			Path:                top100files[i].Path,
			Category:            strings.TrimSuffix(fileCat.category, "/"),
			CategoryDescription: fileCat.description,
		})
	}

	return result, nil
}

// FileCategory returns the category of the given file
func FileCategory(file string) struct{ category, description string } {
	var res struct{ category, description string }

	for _, cat := range lfc {
		if strings.HasPrefix(file, cat.Category) {
			for sb, sbDesc := range cat.SubCategory {
				if strings.HasPrefix(file, sb) {
					res.category = sb
					res.description = sbDesc
					return res
				}
			}

			res.category = cat.Category
			res.description = cat.DefaultDescription
			return res
		}
	}
	return res
}

// LinuxFilesCategories represents one or multiple categories linux file fall into
type LinuxFilesCategories struct {
	Category           string
	SubCategory        map[string]string
	DefaultDescription string
}

var lfc = []LinuxFilesCategories{
	{
		Category:           "/bin/",
		SubCategory:        nil,
		DefaultDescription: "represents some essential user command binaries",
	},
	{
		Category:           "/boot/",
		SubCategory:        nil,
		DefaultDescription: "represents the boot loader static files",
	},
	{
		Category:           "/dev/",
		SubCategory:        nil,
		DefaultDescription: "represents the system devices and special files",
	},
	{
		Category:           "/etc/",
		SubCategory:        nil,
		DefaultDescription: "contains some host-specific system configurations",
	},
	{
		Category:           "/proc/",
		SubCategory:        nil,
		DefaultDescription: "represents the kernel and process information virtual filesystem",
	},
	{
		Category:           "/home/",
		SubCategory:        nil,
		DefaultDescription: "represents the users home directories",
	},
	{
		Category:           "/lib/",
		SubCategory:        nil,
		DefaultDescription: "represents some essential shared libraries and kernel modules",
	},
	{
		Category:           "/media/",
		SubCategory:        nil,
		DefaultDescription: "is the mount point for removable media",
	},
	{
		Category:           "/mnt/",
		SubCategory:        nil,
		DefaultDescription: "is mount point for a temporarily mounted filesystem",
	},
	{
		Category:           "/opt/",
		SubCategory:        nil,
		DefaultDescription: "represents some add-on application software packages",
	},
	{
		Category:           "/root/",
		SubCategory:        nil,
		DefaultDescription: "is the root user home directory (optional)",
	},
	{
		Category:           "/run/",
		SubCategory:        nil,
		DefaultDescription: "represents the run-time variable data",
	},
	{
		Category:           "/sbin/",
		SubCategory:        nil,
		DefaultDescription: "represents some essential system binaries",
	},
	{
		Category:           "/sys/",
		SubCategory:        nil,
		DefaultDescription: "represents the kernel and system information virtual filesystem",
	},
	{
		Category:           "/srv/",
		SubCategory:        nil,
		DefaultDescription: "represents data for services provided by the system",
	},
	{
		Category:           "/tmp/",
		SubCategory:        nil,
		DefaultDescription: "contains temporary files",
	},
	{
		Category: "/usr/",
		SubCategory: map[string]string{
			"/usr/bin/":     "contains most of the executable files that are not needed for booting or repairing the system",
			"/usr/include/": "contains system general-use include files for the C programming language",
			"/usr/lib/":     "contains libraries for programming and packages",
			"/usr/libexec/": "contains binaries run by other programs",
			"/usr/local/":   "contains local binaries",
			"/usr/sbin/":    "contains non-essential standard system binaries",
			"/usr/share/":   "contains read-only architecture independent data files",
			"/usr/src/":     "contains system source code",
		},
		DefaultDescription: "contains shareable and read-only data",
	},
	{
		Category: "/var/",
		SubCategory: map[string]string{
			"/var/cache/": "contains application cache data",
			"/var/crash/": "contains system crash dumps",
			"/var/games/": "contains variable game data",
			"/var/lib/":   "contains variable state information",
			"/var/lock/":  "contains lock files",
			"/var/log/":   "contains log files and directories",
			"/var/mail/":  "contains user mailbox files",
			"/var/opt/":   "contains variable data for /opt",
			"/var/run/":   "contains run-time variable data",
			"/var/spool/": "contains application spool data",
			"/var/tmp/":   "contains temporary files preserved between system reboots",
		},
		DefaultDescription: "contains variable data files. This includes spool directories and files, administrative and logging data, and transient and temporary files.",
	},
}
