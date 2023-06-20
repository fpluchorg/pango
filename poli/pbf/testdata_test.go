package pbf

import (
	"github.com/fpluchorg/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v1 from interface nopbf no symmetric", version.Number{7, 1, 0, ""}, Entry{
			Name:                      "v1",
			Description:               "my desc",
			Tags:                      []string{"tag1", "tag2"},
			FromType:                  FromTypeInterface,
			FromValues:                []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:           []string{"any"},
			SourceUsers:               []string{"user1", "user2"},
			NegateSource:              true,
			DestinationAddresses:      []string{"dest1", "dest2"},
			NegateDestination:         true,
			Applications:              []string{"app1", "app2"},
			Services:                  []string{"s1", "s2"},
			Schedule:                  "my schedule",
			Disabled:                  true,
			Action:                    ActionNoPbf,
			ActiveActiveDeviceBinding: "1",
			Targets: map[string][]string{
				"serial1": {"vsys1", "vsys2"},
				"serial2": nil,
			},
		}},
		{"v1 from interface discard with symmetric", version.Number{7, 1, 0, ""}, Entry{
			Name:                         "v1",
			Description:                  "my desc",
			Tags:                         []string{"tag1", "tag2"},
			FromType:                     FromTypeInterface,
			FromValues:                   []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:              []string{"any"},
			SourceUsers:                  []string{"user1", "user2"},
			NegateSource:                 true,
			DestinationAddresses:         []string{"dest1", "dest2"},
			NegateDestination:            true,
			Applications:                 []string{"app1", "app2"},
			Services:                     []string{"s1", "s2"},
			Schedule:                     "my schedule",
			Disabled:                     true,
			Action:                       ActionDiscard,
			EnableEnforceSymmetricReturn: true,
			SymmetricReturnAddresses:     []string{"10.1.1.1", "10.2.2.2"},
			ActiveActiveDeviceBinding:    "both",
		}},
		{"v1 from interface to vsys no symmetric", version.Number{7, 1, 0, ""}, Entry{
			Name:                 "v1",
			Description:          "my desc",
			Tags:                 []string{"tag1", "tag2"},
			FromType:             FromTypeInterface,
			FromValues:           []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:      []string{"any"},
			SourceUsers:          []string{"any"},
			DestinationAddresses: []string{"any"},
			Applications:         []string{"app"},
			Services:             []string{"s1", "s2"},
			Schedule:             "my schedule",
			Disabled:             true,
			Action:               ActionVsysForward,
			ForwardVsys:          "vsys2",
		}},
		{"v1 forward", version.Number{7, 1, 0, ""}, Entry{
			Name:                   "v1",
			Description:            "my desc",
			FromType:               FromTypeInterface,
			FromValues:             []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:        []string{"any"},
			SourceUsers:            []string{"any"},
			DestinationAddresses:   []string{"any"},
			Applications:           []string{"app"},
			Services:               []string{"application-default"},
			Action:                 ActionForward,
			ForwardEgressInterface: "ethernet1/5",
		}},
		{"v1 forward no nexthop with monitor", version.Number{7, 1, 0, ""}, Entry{
			Name:                               "v1",
			Description:                        "my desc",
			FromType:                           FromTypeInterface,
			FromValues:                         []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:                    []string{"any"},
			SourceUsers:                        []string{"any"},
			DestinationAddresses:               []string{"any"},
			Applications:                       []string{"app"},
			Services:                           []string{"application-default"},
			Action:                             ActionForward,
			ForwardEgressInterface:             "ethernet1/5",
			ForwardMonitorProfile:              "mon-prof",
			ForwardMonitorIpAddress:            "10.5.15.155",
			ForwardMonitorDisableIfUnreachable: true,
		}},
		{"v1 forward ip nexthop", version.Number{7, 1, 0, ""}, Entry{
			Name:                   "v1",
			Description:            "my desc",
			FromType:               FromTypeInterface,
			FromValues:             []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:        []string{"any"},
			SourceUsers:            []string{"any"},
			DestinationAddresses:   []string{"any"},
			Applications:           []string{"app"},
			Services:               []string{"application-default"},
			Action:                 ActionForward,
			ForwardEgressInterface: "ethernet1/5",
			ForwardNextHopType:     ForwardNextHopTypeIpAddress,
			ForwardNextHopValue:    "10.6.16.166",
		}},
		{"v2 from interface nopbf no symmetric", version.Number{9, 0, 0, ""}, Entry{
			Name:                      "v2",
			Description:               "my desc",
			Tags:                      []string{"tag1", "tag2"},
			FromType:                  FromTypeInterface,
			FromValues:                []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:           []string{"any"},
			SourceUsers:               []string{"user1", "user2"},
			NegateSource:              true,
			DestinationAddresses:      []string{"dest1", "dest2"},
			NegateDestination:         true,
			Applications:              []string{"app1", "app2"},
			Services:                  []string{"s1", "s2"},
			Schedule:                  "my schedule",
			Disabled:                  true,
			Action:                    ActionNoPbf,
			ActiveActiveDeviceBinding: "1",
			Targets: map[string][]string{
				"serial1": {"vsys1", "vsys2"},
				"serial2": nil,
			},
		}},
		{"v2 from interface discard with symmetric", version.Number{9, 0, 0, ""}, Entry{
			Name:                         "v2",
			Description:                  "my desc",
			Tags:                         []string{"tag1", "tag2"},
			FromType:                     FromTypeInterface,
			FromValues:                   []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:              []string{"any"},
			SourceUsers:                  []string{"user1", "user2"},
			NegateSource:                 true,
			DestinationAddresses:         []string{"dest1", "dest2"},
			NegateDestination:            true,
			Applications:                 []string{"app1", "app2"},
			Services:                     []string{"s1", "s2"},
			Schedule:                     "my schedule",
			Disabled:                     true,
			Action:                       ActionDiscard,
			EnableEnforceSymmetricReturn: true,
			SymmetricReturnAddresses:     []string{"10.1.1.1", "10.2.2.2"},
			ActiveActiveDeviceBinding:    "both",
		}},
		{"v2 from interface to vsys no symmetric", version.Number{9, 0, 0, ""}, Entry{
			Name:                 "v2",
			Description:          "my desc",
			Tags:                 []string{"tag1", "tag2"},
			FromType:             FromTypeInterface,
			FromValues:           []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:      []string{"any"},
			SourceUsers:          []string{"any"},
			DestinationAddresses: []string{"any"},
			Applications:         []string{"app"},
			Services:             []string{"s1", "s2"},
			Schedule:             "my schedule",
			Disabled:             true,
			Action:               ActionVsysForward,
			ForwardVsys:          "vsys2",
		}},
		{"v2 forward", version.Number{9, 0, 0, ""}, Entry{
			Name:                   "v2",
			Description:            "my desc",
			FromType:               FromTypeInterface,
			FromValues:             []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:        []string{"any"},
			SourceUsers:            []string{"any"},
			DestinationAddresses:   []string{"any"},
			Applications:           []string{"app"},
			Services:               []string{"application-default"},
			Action:                 ActionForward,
			ForwardEgressInterface: "ethernet1/5",
		}},
		{"v2 forward no nexthop with monitor", version.Number{9, 0, 0, ""}, Entry{
			Name:                               "v2",
			Description:                        "my desc",
			FromType:                           FromTypeInterface,
			FromValues:                         []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:                    []string{"any"},
			SourceUsers:                        []string{"any"},
			DestinationAddresses:               []string{"any"},
			Applications:                       []string{"app"},
			Services:                           []string{"application-default"},
			Action:                             ActionForward,
			ForwardEgressInterface:             "ethernet1/5",
			ForwardMonitorProfile:              "mon-prof",
			ForwardMonitorIpAddress:            "10.5.15.155",
			ForwardMonitorDisableIfUnreachable: true,
		}},
		{"v2 forward ip nexthop", version.Number{9, 0, 0, ""}, Entry{
			Name:                   "v2",
			Description:            "my desc",
			FromType:               FromTypeInterface,
			FromValues:             []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:        []string{"any"},
			SourceUsers:            []string{"any"},
			DestinationAddresses:   []string{"any"},
			Applications:           []string{"app"},
			Services:               []string{"application-default"},
			Action:                 ActionForward,
			ForwardEgressInterface: "ethernet1/5",
			ForwardNextHopType:     ForwardNextHopTypeIpAddress,
			ForwardNextHopValue:    "10.6.16.166",
		}},
		{"v2 forward fqdn nexthop", version.Number{9, 0, 0, ""}, Entry{
			Name:                   "v2",
			Uuid:                   "aaaa-bb-ccccc",
			GroupTag:               "myGroupTag",
			Description:            "my desc",
			FromType:               FromTypeInterface,
			FromValues:             []string{"ethernet1/1", "ethernet1/2"},
			SourceAddresses:        []string{"any"},
			SourceUsers:            []string{"any"},
			DestinationAddresses:   []string{"any"},
			Applications:           []string{"app"},
			Services:               []string{"application-default"},
			Action:                 ActionForward,
			ForwardEgressInterface: "ethernet1/5",
			ForwardNextHopType:     ForwardNextHopTypeFqdn,
			ForwardNextHopValue:    "host.example.com",
		}},
	}
}
