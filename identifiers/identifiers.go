package identifiers

// Converting to map for 0(1) searching
func GetKeywords() map[string]bool {
	return map[string]bool{
		"if":   true,
		"else": true,
	}
}

func GetOperators() map[string]bool {
	return map[string]bool{
		"+": true,
		"-": true,
	}
}

func GetSeperators() map[string]bool {
	return map[string]bool{
		"{": true,
		"}": true,
		";": true,
	}
}
