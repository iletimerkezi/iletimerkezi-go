package iletimerkezi

const (
    VersionMajor = "1"
    VersionMinor = "0"
    VersionPatch = "0"
)

func Version() string {
    return "iletimerkezi-go/v" + VersionMajor + "." + VersionMinor + "." + VersionPatch
} 