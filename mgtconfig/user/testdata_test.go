package user

type testCase struct {
	desc string
	conf Entry
}

func getAdministratorsUserTests() []testCase {
	return []testCase{
		{"test1 superUser dynamic", Entry{
			Name:         "t1",
			PasswordHash: "123",
			Role:         SuperUser,
			Type:         Dynamic,
			PublicKey:    "Public-Key",
		}},
		{"test2 superReader dynamic", Entry{
			Name:         "t2",
			PasswordHash: "123",
			Role:         SuperReader,
			Type:         Dynamic,
			PublicKey:    "Public-Key",
		}},
		{"test3 deviceAdmin dynamic", Entry{
			Name:         "t3",
			PasswordHash: "123",
			Role:         DeviceAdmin,
			Type:         Dynamic,
			PublicKey:    "Public-Key",
		}},
		{"test4 deviceReader dynamic", Entry{
			Name:         "t4",
			PasswordHash: "123",
			Role:         DeviceReader,
			Type:         Dynamic,
			PublicKey:    "Public-Key",
		}},
		{"test5 auditadmin custom", Entry{
			Name:         "t5",
			PasswordHash: "123",
			Role:         "auditadmin",
			Type:         Custom,
			PublicKey:    "Public-Key",
		}},
		{"test6 cryptoadmin custom", Entry{
			Name:         "t6",
			PasswordHash: "123",
			Role:         "cryptoadmin",
			Type:         Custom,
			PublicKey:    "Public-Key",
		}},
		{"test7 securityadmin custom", Entry{
			Name:         "t7",
			PasswordHash: "123",
			Role:         "securityadmin",
			Type:         Custom,
			PublicKey:    "Public-Key",
		}},
	}
}

func getPanoramaAdministratorsUserTests() []testCase {
	return []testCase{
		{"test1 superUser dynamic", Entry{
			Name:         "t1",
			PasswordHash: "123",
			Role:         SuperUser,
			Type:         Dynamic,
			PublicKey:    "Public-Key",
		}},
		{"test2 superReader dynamic", Entry{
			Name:         "t2",
			PasswordHash: "123",
			Role:         SuperReader,
			Type:         Dynamic,
			PublicKey:    "Public-Key",
		}},
		{"test3 deviceAdmin dynamic", Entry{
			Name:         "t3",
			PasswordHash: "123",
			Role:         "api-vmauthkey-automation",
			Type:         Custom,
			PublicKey:    "Public-Key",
		}},
		{"test4 deviceReader dynamic", Entry{
			Name:         "t4",
			PasswordHash: "123",
			Role:         "api-object-automation",
			Type:         Custom,
			PublicKey:    "Public-Key",
		}},
		{"test5 auditadmin custom", Entry{
			Name:         "t5",
			PasswordHash: "123",
			Role:         "api-rules-automation",
			Type:         Custom,
			PublicKey:    "Public-Key",
		}},
		{"test6 cryptoadmin custom", Entry{
			Name:         "t6",
			PasswordHash: "123",
			Role:         "gpio-esi-ro",
			Type:         Custom,
			PublicKey:    "Public-Key",
		}},
		{"test7 securityadmin custom", Entry{
			Name:         "t7",
			PasswordHash: "123",
			Role:         "installeradmin",
			Type:         Custom,
			PublicKey:    "Public-Key",
		}},
		{"test8 panorama-admin dynamic", Entry{
			Name:         "t8",
			PasswordHash: "123",
			Role:         "panorama-admin",
			Type:         Dynamic,
			PublicKey:    "Public-Key",
		}},
	}
}
