package user

import (
	"encoding/xml"
	"github.com/fpluchorg/pango/util"
	"github.com/fpluchorg/pango/version"
)

const (
	Custom        = "custom"
	DeviceAdmin   = "deviceadmin"
	DeviceReader  = "devicereader"
	Dynamic       = "dynamic"
	Error         = "ERROR"
	PanoramaAdmin = "panorama-admin"
	SuperReader   = "superreader"
	SuperUser     = "superuser"
)

// Entry is a normalized, version independent representation of a device group.
type Entry struct {
	Name         string
	PasswordHash string
	PublicKey    string
	Type         string
	Role         string
}

// Copy copies the information from source's Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Name = s.Name
	o.PasswordHash = s.PasswordHash
	o.PublicKey = s.PublicKey
	o.Type = s.Type
	o.Role = s.Role
}

/** Structs / functions for normalization. **/

func (o Entry) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return o.Name, fn(o)
}

type normalizer interface {
	Normalize() []Entry
	Names() []string
}

type container_v1 struct {
	Answer []entry_v1 `xml:"entry"`
}

func (o *container_v1) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v1) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:         o.Name,
		PasswordHash: o.PasswordHash,
	}
	if o.PublicKey != nil {
		ans.PublicKey = *o.PublicKey
	}

	var (
		superUser     *string
		superReader   *string
		panoramaAdmin *string
		deviceAdmin   *string
		deviceReader  *string
	)
	if o.SuperUser != nil && *o.SuperUser == util.YesNo(true) {
		superUser = o.SuperUser
	}
	if o.SuperReader != nil && *o.SuperReader == util.YesNo(true) {
		superReader = o.SuperReader
	}
	if o.DeviceAdmin != nil && *o.DeviceAdmin == util.EmptyString {
		deviceAdmin = o.DeviceAdmin
	}
	if o.DeviceReader != nil && *o.DeviceReader == util.EmptyString {
		deviceReader = o.DeviceReader
	}
	if o.PanoramaAdmin != nil && *o.PanoramaAdmin == util.YesNo(true) {
		panoramaAdmin = o.PanoramaAdmin
	}
	if (superUser != nil && *superUser == util.YesNo(true)) ||
		(superReader != nil && *superReader == util.YesNo(true)) ||
		(deviceAdmin != nil && *deviceAdmin == util.EmptyString) ||
		(deviceReader != nil && *deviceReader == util.EmptyString) ||
		(panoramaAdmin != nil && *panoramaAdmin == util.YesNo(true)) {
		ans.Type = Dynamic
		if superUser != nil && *superUser == util.YesNo(true) {
			ans.Role = SuperUser
		} else if superReader != nil && *superReader == util.YesNo(true) {
			ans.Role = SuperReader
		} else if deviceAdmin != nil && *deviceAdmin == util.EmptyString {
			ans.Role = DeviceAdmin
		} else if deviceReader != nil && *deviceReader == util.EmptyString {
			ans.Role = DeviceReader
		} else if panoramaAdmin != nil && *panoramaAdmin == util.YesNo(true) {
			ans.Role = PanoramaAdmin
		} else {
			ans.Role = Error
		}
	} else if o.CustomProfile != nil {
		ans.Type = Custom
		ans.Role = *o.CustomProfile
	}

	return ans
}

type entry_v1 struct {
	XMLName       xml.Name `xml:"entry"`
	Name          string   `xml:"name,attr"`
	PasswordHash  string   `xml:"phash"`
	PublicKey     *string  `xml:"public-key"`
	SuperUser     *string  `xml:"permissions>role-based>superuser,omitempty"`
	SuperReader   *string  `xml:"permissions>role-based>superreader,omitempty"`
	DeviceReader  *string  `xml:"permissions>role-based>devicereader,omitempty"`
	DeviceAdmin   *string  `xml:"permissions>role-based>deviceadmin,omitempty"`
	PanoramaAdmin *string  `xml:"permissions>role-based>panorama-admin,omitempty"`
	CustomProfile *string  `xml:"permissions>role-based>custom>profile,omitempty"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:         e.Name,
		PasswordHash: e.PasswordHash,
	}

	if e.PublicKey != util.EmptyString {
		ans.PublicKey = &e.PublicKey
	}

	yes := util.YesNo(true)
	emptyString := util.EmptyString

	switch e.Type {
	case Dynamic:
		switch e.Role {
		case SuperUser:
			ans.SuperUser = &yes
		case SuperReader:
			ans.SuperReader = &yes
		case DeviceReader:
			ans.DeviceReader = &emptyString
		case DeviceAdmin:
			ans.DeviceAdmin = &emptyString
		case PanoramaAdmin:
			ans.PanoramaAdmin = &yes
		default:
			break
		}
	case Custom:
		ans.CustomProfile = &e.Role
	default:
		break
	}

	return ans
}
