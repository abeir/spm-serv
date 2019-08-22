package po

import (
	"spm-serv/core"
	"strconv"
	"time"
)

type UpgradeVersion struct {
	Id string			`json:"id"`

	Version string 		`json:"version"`

	VersionSort string	`json:"versionSort"`

	Description string 	`json:"description"`

	Path string 		`json:"path"`

	CreatedAt time.Time	`json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *UpgradeVersion) IsEmpty() bool{
	return u.Id=="" && u.Version==""
}

//
func (u *UpgradeVersion) ToVersionSort() string{
	if u.Version=="" {
		return ""
	}
	ver, _ := core.NewVersion(u.Version)

	major := strconv.Itoa(ver.Major)
	minor := strconv.Itoa(ver.Minor)
	revision := strconv.Itoa(ver.Revision)

	major, _ = core.PadLeft(major, "0", 6)
	minor, _ = core.PadLeft(minor, "0", 6)
	revision, _ = core.PadLeft(revision, "0", 6)
	return major + "-" + minor + "-" + revision
}
