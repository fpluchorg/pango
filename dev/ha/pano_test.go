package ha

import (
	"reflect"
	"testing"

	"github.com/fpluchorg/pango/testdata"
	"github.com/fpluchorg/pango/version"
)

func TestPanoNormalization(t *testing.T) {
	testCases := getTests()

	mc := &testdata.MockClient{}
	ns := PanoramaNamespace(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc.Version = tc.version
			mc.Reset()
			mc.AddResp("")
			err := ns.Set("my template", "", "", tc.conf)
			if err != nil {
				t.Errorf("Error in set: %s", err)
			} else {
				mc.AddResp(mc.Elm)
				r, err := ns.Get("my template", "", "")
				if err != nil {
					t.Errorf("Error in get: %s", err)
				}
				if !reflect.DeepEqual(tc.conf, r) {
					t.Errorf("%#v != %#v", tc.conf, r)
				}
			}
		})
	}
}

func TestPanoHa2StateSyncEnableIsTrueWithNoStateSyncConfig(t *testing.T) {
	mc := &testdata.MockClient{
		Version: version.Number{6, 1, 0, ""},
	}
	ns := FirewallNamespace(mc)
	mc.AddResp(`
<high-availability>
    <enabled>yes</enabled>
    <group>
        <group-id>2</group-id>
        <description>My description</description>
        <peer-ip>127.0.0.1</peer-ip>
        <configuration-synchronization>
            <enabled>yes</enabled>
        </configuration-synchronization>
    </group>
</high-availability>
`)
	ans, err := ns.Get()
	if err != nil {
		t.Errorf("Error in get: %s", err)
	}
	if ans.GroupId != 2 {
		t.Errorf("GroupId is %d, not 2", ans.GroupId)
	}
	if !ans.Ha2StateSyncEnable {
		t.Errorf("Ha2StateSyncEnable is incorrectly set to %t", ans.Ha2StateSyncEnable)
	}
}

func TestPanoHa2StateSyncEnableIsTrueWithStateSyncConfig(t *testing.T) {
	mc := &testdata.MockClient{
		Version: version.Number{6, 1, 0, ""},
	}
	ns := FirewallNamespace(mc)
	mc.AddResp(`
<high-availability>
    <enabled>yes</enabled>
    <group>
        <group-id>2</group-id>
        <description>My description</description>
        <peer-ip>127.0.0.1</peer-ip>
        <configuration-synchronization>
            <enabled>yes</enabled>
        </configuration-synchronization>
        <state-synchronization>
            <transport>mytransport</transport>
        </state-synchronization>
    </group>
</high-availability>
`)
	ans, err := ns.Get()
	if err != nil {
		t.Errorf("Error in get: %s", err)
	}
	if ans.GroupId != 2 {
		t.Errorf("GroupId is %d, not 2", ans.GroupId)
	}
	if !ans.Ha2StateSyncEnable {
		t.Errorf("Ha2StateSyncEnable is incorrectly set to %t", ans.Ha2StateSyncEnable)
	}
}

func TestPanoHa2StateSyncEnableCanBeFalse(t *testing.T) {
	mc := &testdata.MockClient{
		Version: version.Number{6, 1, 0, ""},
	}
	ns := FirewallNamespace(mc)
	mc.AddResp(`
<high-availability>
    <enabled>yes</enabled>
    <group>
        <group-id>2</group-id>
        <description>My description</description>
        <peer-ip>127.0.0.1</peer-ip>
        <configuration-synchronization>
            <enabled>yes</enabled>
        </configuration-synchronization>
        <state-synchronization>
            <enabled>no</enabled>
            <transport>mytransport</transport>
        </state-synchronization>
    </group>
</high-availability>
`)
	ans, err := ns.Get()
	if err != nil {
		t.Errorf("Error in get: %s", err)
	}
	if ans.GroupId != 2 {
		t.Errorf("GroupId is %d, not 2", ans.GroupId)
	}
	if ans.Ha2StateSyncEnable {
		t.Errorf("Ha2StateSyncEnable is incorrectly set to %t", ans.Ha2StateSyncEnable)
	}
}
