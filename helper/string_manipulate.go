package helper

func ConvertToSnakeCase(s string) string {
	// ConvertToSnakeCase converts a string to snake case
	var snake string
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			snake += "_"
		}
		snake += string(r)
	}
	return snake
}
