package str

func ParseBool(s string) bool {
	switch s {
	case "true", "True", "TRUE", "1":
		return true
	default:
		return false
	}
}
