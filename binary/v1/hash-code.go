package ignite

// HashCode calculates Java hash code for string
func HashCode(s string) int32 {
	return HashCodeForSlice([]byte(s))
}

// HashCodeForSlice calculates Java hash code for byte array
func HashCodeForSlice(b []byte) int32 {
	if len(b) == 0 {
		// HashCode of empty string return 0 so it can be used in QuerySQLFields with no TableName but a Schema
		return 0
	}

	h := uint32(0)
	for i := 0; i < len(b); i++ {
		h = 31*h + uint32(b[i])
	}
	return int32(h)
}
