package device

import (
	"encoding/xml"

	"github.com/fpluchorg/pango/version"
)

// Entry is a normalized, version independent representation of a device group.
type Entry struct {
	Name     string
	Hostname string
	Ip       string
}

// Copy copies the information from source's Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Name = s.Name
	o.Hostname = s.Hostname
	o.Ip = s.Ip
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
		Name:     o.Name,
		Hostname: o.Hostname,
		Ip:       o.Ip,
	}

	return ans
}

type entry_v1 struct {
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`
	Hostname string   `xml:"hostname,omitempty"`
	Ip       string   `xml:"ip,omitempty"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:     e.Name,
		Hostname: e.Hostname,
		Ip:       e.Ip,
	}

	return ans
}
