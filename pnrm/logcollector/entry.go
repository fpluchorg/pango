package logcollector

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
	Name                    string
	PanoramaServer          string
	Service                 Service
	SpeedDuplex             string
	MTU                     int
	NTPPrimaryServerAddress string
	DNSPrimaryServerAddress string
	Hostname                string
	Timezone                string
	//IPType                  string
	IPAddress       string
	Netmask         string
	DefaultGateway  string
	PublicIPAddress string

	raw map[string]string
}

type Service struct {
	DisableDeviceLogCollection         bool
	DisableCollectorGroupCommunication bool
	DisableSSH                         bool
	DisableICMP                        bool
	DisableSNMP                        bool
	DisableSyslogForwarding            bool
	DisableDeviceTelemetryForwarding   bool
}

// Copy copies the information from source's Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Name = s.Name
	o.PanoramaServer = s.PanoramaServer
	o.Service = s.Service
	o.SpeedDuplex = s.SpeedDuplex
	o.MTU = s.MTU
	o.NTPPrimaryServerAddress = s.NTPPrimaryServerAddress
	o.DNSPrimaryServerAddress = s.DNSPrimaryServerAddress
	o.Hostname = s.Hostname
	o.Timezone = s.Timezone
	//o.IPType = s.IPType
	o.IPAddress = s.IPAddress
	o.Netmask = s.Netmask
	o.DefaultGateway = s.DefaultGateway
	o.PublicIPAddress = s.PublicIPAddress
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

	if o.DeviceConfig.System != nil {
		if o.DeviceConfig.System.PanoramaServer != "" {
			ans.PanoramaServer = o.DeviceConfig.System.PanoramaServer
		}
		ans.Service = Service{
			DisableDeviceLogCollection:         util.AsBool(o.DeviceConfig.System.Service.DisableDeviceLogCollection),
			DisableCollectorGroupCommunication: util.AsBool(o.DeviceConfig.System.Service.DisableCollectorGroupCommunication),
			DisableSSH:                         util.AsBool(o.DeviceConfig.System.Service.DisableSSH),
			DisableICMP:                        util.AsBool(o.DeviceConfig.System.Service.DisableICMP),
			DisableSNMP:                        util.AsBool(o.DeviceConfig.System.Service.DisableSNMP),
			DisableSyslogForwarding:            util.AsBool(o.DeviceConfig.System.Service.DisableSyslogForwarding),
			DisableDeviceTelemetryForwarding:   util.AsBool(o.DeviceConfig.System.Service.DisableDeviceTelemetryForwarding),
		}
		ans.SpeedDuplex = o.DeviceConfig.System.SpeedDuplex
		ans.MTU = o.DeviceConfig.System.MTU
		ans.NTPPrimaryServerAddress = o.DeviceConfig.System.NTPPrimaryServerAddress
		ans.DNSPrimaryServerAddress = o.DeviceConfig.System.DNSPrimaryServerAddress
		ans.Hostname = o.DeviceConfig.System.Hostname
		ans.Timezone = o.DeviceConfig.System.Timezone
		// IPType
		ans.IPAddress = o.DeviceConfig.System.IPAddress
		ans.Netmask = o.DeviceConfig.System.Netmask
		ans.DefaultGateway = o.DeviceConfig.System.DefaultGateway
		ans.PublicIPAddress = o.DeviceConfig.System.PublicIPAddress
	}

	raw := make(map[string]string)

	if o.DiskSettings != nil {
		raw["disk-settings"] = util.CleanRawXml(o.DiskSettings.Text)
	}

	if len(raw) > 0 {
		ans.raw = raw
	}

	return ans
}

type entry_v1 struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	DeviceConfig *deviceconfig_v1 `xml:"deviceconfig"`
	DiskSettings *util.RawXml     `xml:"disk-settings"`
}

type deviceconfig_v1 struct {
	System *system_v1 `xml:"system"`
}

type system_v1 struct {
	PanoramaServer          string      `xml:"panorama-server,omitempty"`
	Service                 *service_v1 `xml:"service"`
	SpeedDuplex             string      `xml:"speed-duplex"`
	MTU                     int         `xml:"mtu"`
	NTPPrimaryServerAddress string      `xml:"ntp-servers>primary-ntp-server>ntp-server-address,omitempty"`
	DNSPrimaryServerAddress string      `xml:"dns-settings>servers>primary,omitempty"`
	Hostname                string      `xml:"hostname,omitempty"`
	Timezone                string      `xml:"timezone,omitempty"`
	//IPType                  string
	IPAddress       string `xml:"ip-address,omitempty"`
	Netmask         string `xml:"netmask,omitempty"`
	DefaultGateway  string `xml:"default-gateway,omitempty"`
	PublicIPAddress string `xml:"public-ip-address,omitempty"`
}

type service_v1 struct {
	DisableDeviceLogCollection         string `xml:"disable-device-log-collection"`
	DisableCollectorGroupCommunication string `xml:"disable-collector-group-communication"`
	DisableSSH                         string `xml:"disable-ssh"`
	DisableICMP                        string `xml:"disable-icmp"`
	DisableSNMP                        string `xml:"disable-snmp"`
	DisableSyslogForwarding            string `xml:"disable-syslog-forwarding"`
	DisableDeviceTelemetryForwarding   string `xml:"disable-device-telemetry-forwarding"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name: e.Name,
	}

	ans.DeviceConfig.System = &system_v1{
		PanoramaServer:          e.PanoramaServer,
		SpeedDuplex:             e.SpeedDuplex,
		MTU:                     e.MTU,
		NTPPrimaryServerAddress: e.NTPPrimaryServerAddress,
		DNSPrimaryServerAddress: e.DNSPrimaryServerAddress,
		Hostname:                e.Hostname,
		Timezone:                e.Timezone,
		// IPType: e.IPType,
		IPAddress:       e.IPAddress,
		Netmask:         e.Netmask,
		DefaultGateway:  e.DefaultGateway,
		PublicIPAddress: e.PublicIPAddress,
	}
	ans.DeviceConfig.System.Service = &service_v1{
		DisableDeviceLogCollection:         util.YesNo(e.Service.DisableDeviceLogCollection),
		DisableCollectorGroupCommunication: util.YesNo(e.Service.DisableCollectorGroupCommunication),
		DisableSSH:                         util.YesNo(e.Service.DisableSSH),
		DisableICMP:                        util.YesNo(e.Service.DisableICMP),
		DisableSNMP:                        util.YesNo(e.Service.DisableSNMP),
		DisableSyslogForwarding:            util.YesNo(e.Service.DisableSyslogForwarding),
		DisableDeviceTelemetryForwarding:   util.YesNo(e.Service.DisableDeviceTelemetryForwarding),
	}

	if text, present := e.raw["disk-settings"]; present {
		ans.DiskSettings = &util.RawXml{text}
	}

	return ans
}
