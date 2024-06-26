// Package util contains various shared structs and functions used across
// the pango package.
package util

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

// VsysEntryType defines an entry config node with vsys entries underneath.
type VsysEntryType struct {
	Entries []VsysEntry `xml:"entry"`
}

// VsysEntry defines the "vsys" xpath node under a VsysEntryType config node.
type VsysEntry struct {
	XMLName xml.Name   `xml:"entry"`
	Serial  string     `xml:"name,attr"`
	Vsys    *EntryType `xml:"vsys"`
}

// VsysEntToMap normalizes a VsysEntryType pointer into a map.
func VsysEntToMap(ve *VsysEntryType) map[string][]string {
	if ve == nil {
		return nil
	}

	ans := make(map[string][]string)
	for i := range ve.Entries {
		ans[ve.Entries[i].Serial] = EntToStr(ve.Entries[i].Vsys)
	}

	return ans
}

// MapToVsysEnt converts a map into a VsysEntryType pointer.
//
// This struct is used for "Target" information on Panorama when dealing with
// various policies.  Maps are unordered, but FWICT Panorama doesn't seem to
// order anything anyways when doing things in the GUI, so hopefully this is
// ok...?
func MapToVsysEnt(e map[string][]string) *VsysEntryType {
	if len(e) == 0 {
		return nil
	}

	i := 0
	ve := make([]VsysEntry, len(e))
	for key := range e {
		ve[i].Serial = key
		ve[i].Vsys = StrToEnt(e[key])
		i++
	}

	return &VsysEntryType{ve}
}

// YesNo returns "yes" on true, "no" on false.
func YesNo(v bool) string {
	if v {
		return "yes"
	}
	return "no"
}

// YesNoEmpty returns "yes" on true, "no" on false and empty string if null
func YesNoEmpty(v *bool) string {
	if v == nil {
		return EmptyString
	}
	if *v {
		return "yes"
	}
	return "no"
}

// AsBool returns true on yes, else false.
func AsBool(val string) bool {
	if val == "yes" {
		return true
	}
	return false
}

// AsBoolEmpty returns true on yes, false on no and nil on empty string
func AsBoolEmpty(val string) *bool {
	if val == EmptyString {
		return nil
	}
	if val == "yes" {
		trueVal := true
		return &trueVal
	}
	falseVal := true
	return &falseVal
}

// AsXpath makes an xpath out of the given interface.
func AsXpath(i interface{}) string {
	switch val := i.(type) {
	case string:
		return val
	case []string:
		return fmt.Sprintf("/%s", strings.Join(val, "/"))
	default:
		return ""
	}
}

// AsEntryXpath returns the given values as an entry xpath segment.
func AsEntryXpath(vals []string) string {
	if len(vals) == 0 || (len(vals) == 1 && vals[0] == "") {
		return "entry"
	}

	var buf bytes.Buffer

	buf.WriteString("entry[")
	for i := range vals {
		if i != 0 {
			buf.WriteString(" or ")
		}
		buf.WriteString("@name='")
		buf.WriteString(vals[i])
		buf.WriteString("'")
	}
	buf.WriteString("]")

	return buf.String()
}

// AsMemberXpath returns the given values as a member xpath segment.
func AsMemberXpath(vals []string) string {
	var buf bytes.Buffer

	buf.WriteString("member[")
	for i := range vals {
		if i != 0 {
			buf.WriteString(" or ")
		}
		buf.WriteString("text()='")
		buf.WriteString(vals[i])
		buf.WriteString("'")
	}

	buf.WriteString("]")

	return buf.String()
}

// LogCollectorGroupDeviceXpathPrefix returns the logcollector xpath prefix of the given logcollector name.
func LogCollectorGroupDeviceXpathPrefix(logcollectorgroup, device string) []string {
	return []string{
		"config",
		"devices",
		AsEntryXpath([]string{"localhost.localdomain"}),
		"log-collector-group",
		AsEntryXpath([]string{logcollectorgroup}),
		"logfwd-setting",
		"devices",
		AsEntryXpath([]string{device}),
	}
}

// LogCollectorGroupXpathPrefix returns the logcollector xpath prefix of the given logcollector name.
func LogCollectorGroupXpathPrefix(logcollectorgroup string) []string {
	return []string{
		"config",
		"devices",
		AsEntryXpath([]string{"localhost.localdomain"}),
		"log-collector-group",
		AsEntryXpath([]string{logcollectorgroup}),
	}
}

// LogCollectorXpathPrefix returns the logcollector xpath prefix of the given logcollector name.
func LogCollectorXpathPrefix(logcollector string) []string {
	return []string{
		"config",
		"devices",
		AsEntryXpath([]string{"localhost.localdomain"}),
		"log-collector",
		AsEntryXpath([]string{logcollector}),
	}
}

// TemplateXpathPrefix returns the template xpath prefix of the given template name.
func TemplateXpathPrefix(tmpl, ts string) []string {
	if tmpl != EmptyString {
		return []string{
			"config",
			"devices",
			AsEntryXpath([]string{"localhost.localdomain"}),
			"template",
			AsEntryXpath([]string{tmpl}),
		}
	}

	return []string{
		"config",
		"devices",
		AsEntryXpath([]string{"localhost.localdomain"}),
		"template-stack",
		AsEntryXpath([]string{ts}),
	}
}

// DeviceGroupXpathPrefix returns a device group xpath prefix.
// If the device group is empty, then the default is "shared".
func DeviceGroupXpathPrefix(dg string) []string {
	if dg == "" || dg == "shared" {
		return []string{"config", "shared"}
	}

	return []string{
		"config",
		"devices",
		AsEntryXpath([]string{"localhost.localdomain"}),
		"device-group",
		AsEntryXpath([]string{dg}),
	}
}

// VsysXpathPrefix returns a vsys xpath prefix.
func VsysXpathPrefix(vsys string) []string {
	if vsys == "" {
		vsys = "vsys1"
	} else if vsys == "shared" {
		return []string{"config", "shared"}
	}

	return []string{
		"config",
		"devices",
		AsEntryXpath([]string{"localhost.localdomain"}),
		"vsys",
		AsEntryXpath([]string{vsys}),
	}
}

// PanoramaXpathPrefix returns the panorama xpath prefix.
func PanoramaXpathPrefix() []string {
	return []string{
		"config",
		"panorama",
	}
}

// StripPanosPackaging removes the response / result and an optional third
// containing XML tag from the given byte slice.
func StripPanosPackaging(input []byte, tag string) []byte {
	var index int
	gt := []byte(">")
	lt := []byte("<")

	// Remove response.
	index = bytes.Index(input, gt)
	ans := input[index+1:]
	index = bytes.LastIndex(ans, lt)
	ans = ans[:index]

	// Remove result.
	index = bytes.Index(ans, gt)
	ans = ans[index+1:]
	index = bytes.LastIndex(ans, lt)
	ans = ans[:index]

	ans = bytes.TrimSpace(ans)

	if tag != "" {
		if bytes.HasPrefix(ans, []byte("<"+tag+" ")) || bytes.HasPrefix(ans, []byte("<"+tag+">")) {
			index = bytes.Index(ans, gt)
			ans = ans[index+1:]
			if len(ans) > 0 {
				index = bytes.LastIndex(ans, lt)
				ans = ans[:index]
				ans = bytes.TrimSpace(ans)
			}
		}
	}

	return ans
}

// CdataText is for getting CDATA contents of XML docs.
type CdataText struct {
	Text string `xml:",cdata"`
}

// RawXml is what allows the use of Edit commands on a XPATH without
// truncating any other child objects that may be attached to it.
type RawXml struct {
	Text string `xml:",innerxml"`
}

// CleanRawXml removes extra XML attributes from RawXml objects without
// requiring us to have to parse everything.
func CleanRawXml(v string) string {
	re := regexp.MustCompile(` admin="\S+" dirtyId="\d+" time="\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}"`)
	return re.ReplaceAllString(v, "")
}

// ValidMovement returns if the movement constant is valid or not.
func ValidMovement(v int) bool {
	switch v {
	case MoveSkip, MoveBefore, MoveDirectlyBefore, MoveAfter, MoveDirectlyAfter, MoveTop, MoveBottom:
		return true
	}

	return false
}

// RelativeMovement returns if the movement constant is a relative movement.
func RelativeMovement(v int) bool {
	switch v {
	case MoveBefore, MoveDirectlyBefore, MoveAfter, MoveDirectlyAfter:
		return true
	}

	return false
}

// ValidateRulebase validates the device group and rulebase pairing for
// Panorama policies.
func ValidateRulebase(dg, base string) error {
	switch base {
	case "":
		return fmt.Errorf("rulebase must be specified")
	case Rulebase:
		if dg != "shared" {
			return fmt.Errorf("rulebase %q requires \"shared\" device group", base)
		}
	case PreRulebase, PostRulebase:
	default:
		return fmt.Errorf("unknown rulebase %q", base)
	}

	return nil
}

// RefactorElement for removing default struct tag name in case it was auto append
func RefactorElement(str string) string {
	if str[1:len(EntryV1)+1] == EntryV1 {
		str = str[len(EntryV1)+2:]
		str = str[0 : len(str)-len(EntryV1)-3]
	}
	return str
}
