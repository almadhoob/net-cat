package chat



func StrToBy(args []string) []byte {
	
	totalLen := 0
	for _, arg := range args {
		totalLen += len(arg)
	}

	result := make([]byte, totalLen)

	offset := 0
	for _, arg := range args {
		n := copy(result[offset:], arg)
		offset += n
	}

	return result
}
