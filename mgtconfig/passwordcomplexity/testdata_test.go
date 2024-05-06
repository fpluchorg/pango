package passwordcomplexity

type testCase struct {
	desc string
	conf Entry
}

func getPasswordComplexityTests() []testCase {
	return []testCase{
		{"test1 Password Complexity", Entry{
			MinimumLength:                  8,
			Enabled:                        "true",
			MinimumUppercaseLetters:        1,
			MinimumLowercaseLetters:        1,
			MinimumNumericLetters:          1,
			MinimumSpecialCharacters:       1,
			BlockRepeatedCharacters:        1,
			BlockUsernameInclusion:         "true",
			NewPasswordDiffersByCharacters: 1,
			PasswordChangeOnFirstLogin:     "true",
			PasswordHistoryCount:           1,
			PasswordChangePeriodBlock:      1,
			ExpirationPeriod:               1,
			ExpirationWarningPeriod:        1,
			PostExpirationAdminLoginCount:  1,
			PostExpirationGracePeriod:      1,
		}},
	}
}
