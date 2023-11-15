package user

import (
	"encoding/xml"

	"github.com/fpluchorg/pango/util"
	"github.com/fpluchorg/pango/version"
)

// Entry is a normalized, version independent representation of a device group.
type Entry struct {
	Name         string
	PasswordHash string
	Type         string
	Role         string
}

// Copy copies the information from source's Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Name = s.Name
	o.PasswordHash = s.PasswordHash
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

	if o.SuperUser == "yes" || o.SuperReader == "yes" || o.PanoramaAdmin == "yes" {
		ans.Type = "dynamic"
		if o.SuperUser == "yes" {
			ans.Role = "superuser"
		} else if o.SuperReader == "yes" {
			ans.Role = "superreader"
		} else if o.PanoramaAdmin == "yes" {
			ans.Role = "panorama-admin"
		} else {
			ans.Role = "ERROR"
		}
	} else if o.CustomProfile != "" {
		ans.Type = "custom"
		ans.Role = o.CustomProfile
	}

	return ans
}

type entry_v1 struct {
	XMLName       xml.Name `xml:"entry"`
	Name          string   `xml:"name,attr"`
	PasswordHash  string   `xml:"phash"`
	SuperUser     string   `xml:"permissions>role-based>superuser,omitempty"`
	SuperReader   string   `xml:"permissions>role-based>superreader,omitempty"`
	PanoramaAdmin string   `xml:"permissions>role-based>panorama-admin,omitempty"`
	CustomProfile string   `xml:"permissions>role-based>custom>profile,omitempty"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:         e.Name,
		PasswordHash: e.PasswordHash,
	}

	if e.Type == "dynamic" {
		if e.Role == "superuser" {
			ans.SuperUser = util.YesNo(true)
		} else if e.Role == "superreader" {
			ans.SuperReader = util.YesNo(true)
		} else if e.Role == "panorama-admin" {
			ans.PanoramaAdmin = util.YesNo(true)
		}
	} else if e.Type == "custom" {
		ans.CustomProfile = e.Role
	}

	return ans
}
