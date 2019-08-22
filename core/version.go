package core

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//Version 版本号
type Version struct {
	//主版本号
	Major int
	//次要版本号
	Minor int
	//修订版本号
	Revision int
}

//ParseVersion 解析版本号字符串
func (v *Version)ParseVersion(ver string) error{
	if ver == "" {
		return errors.New("version is blank")
	}
	verNums := strings.Split(ver, ".")
	if verNums==nil || len(verNums) != 3 {
		return errors.New("version number format error")
	}
	major, err := strconv.Atoi(verNums[0])
	if err!=nil {
		return err
	}
	minor, err := strconv.Atoi(verNums[1])
	if err!=nil {
		return err
	}
	revision, err := strconv.Atoi(verNums[2])
	if err!=nil {
		return err
	}
	v.Major = major
	v.Minor = minor
	v.Revision = revision
	return nil
}

func (v *Version) String() string{
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Revision)
}


//CheckVersion 比较两个版本号大小，大于参数返回1，等于参数返回0，小于参数返回-1
func (v *Version) Compare(ver *Version) int8{
	if v.Major < ver.Major {
		return -1
	}
	if v.Major > ver.Major {
		return 1
	}
	if v.Minor < ver.Minor {
		return -1
	}
	if v.Minor > ver.Minor {
		return 1
	}
	if v.Revision < ver.Revision {
		return -1
	}
	if v.Revision > ver.Revision {
		return 1
	}
	return 0
}

func NewVersion(ver string) (*Version, error){
	version := &Version{}
	err := version.ParseVersion(ver)
	return version, err
}
