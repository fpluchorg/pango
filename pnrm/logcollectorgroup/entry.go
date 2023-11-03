package logcollectorgroup

import (
	"encoding/xml"

	"github.com/fpluchorg/pango/util"
	"github.com/fpluchorg/pango/version"
)

// Entry is a normalized, version independent representation of a device group.
//
// Devices is a map where the key is the serial number of the target device and
// the value is a list of specific vsys on that device.  The list of vsys is
// nil if all vsys on that device should be included or if the device is a
// virtual firewall (and thus only has vsys1).
type Entry struct {
	Name               string
	MinRetentionPeriod int
	ForwardToAll       bool
	Collectors         []string

	raw map[string]string
}

// Copy copies the information from source's Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Name = s.Name
	o.MinRetentionPeriod = s.MinRetentionPeriod
	o.ForwardToAll = s.ForwardToAll
	o.Collectors = s.Collectors
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
		Name: o.Name,
	}

	ans.MinRetentionPeriod = o.GeneralSetting.Management.MinRetentionPeriod
	ans.ForwardToAll = util.AsBool(o.GeneralSetting.Management.ForwardToAll)
	ans.Collectors = o.LogForwardSetting.Collectors

	raw := make(map[string]string)

	if o.LogForwardSetting.Devices != nil {
		raw["devices"] = util.CleanRawXml(o.LogForwardSetting.Devices.Text)
	}

	if len(raw) > 0 {
		ans.raw = raw
	}

	return ans
}

type entry_v1 struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	MonitoringSetting *monitoringsetting_v1 `xml:"monitoring-setting"`
	GeneralSetting    *generalsetting_v1    `xml:"general-setting"`
	LogForwardSetting *logfwdsetting_v1     `xml:"logfwd-setting"`
}

type monitoringsetting_v1 struct {
	SNMPSetting *snmpsetting_v1 `xml:"snmp-setting"`
}

type snmpsetting_v1 struct {
	AccessSetting *accesssetting_v1 `xml:"access-setting"`
}

type accesssetting_v1 struct {
	Version *version_v1 `xml:"version"`
}

type version_v1 struct {
	V2C *util.RawXml `xml:"v2c"`
}

type generalsetting_v1 struct {
	Management *management_v1 `xml:"management"`
}

type management_v1 struct {
	MinRetentionPeriod int          `xml:"min-retention-period"`
	ForwardToAll       string       `xml:"forward-to-all"`
	QuotaSettings      *util.RawXml `xml:"quota-settings"`
}

type logfwdsetting_v1 struct {
	Collectors []string     `xml:"collectors>member"`
	Devices    *util.RawXml `xml:"devices"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name: e.Name,
	}

	ans.GeneralSetting.Management.MinRetentionPeriod = e.MinRetentionPeriod
	ans.GeneralSetting.Management.ForwardToAll = util.YesNo(e.ForwardToAll)

	ans.LogForwardSetting.Collectors = e.Collectors

	if text, present := e.raw["devices"]; present {
		ans.LogForwardSetting.Devices = &util.RawXml{text}
	}

	return ans
}
