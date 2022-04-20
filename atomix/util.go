package atomix

// StringsToBytes converts []string to [][]byte - 'a'.
func StringsToBytes(ss []string) [][]byte {
	bb := make([][]byte, len(ss))
	for i, s := range ss {
		bb[i] = toIndex(s)
	}
	return bb
}

func toIndex(ss string) []byte {
	xb := make([]byte, len(ss))
	for k, s := range ss {
		xb[k] = byte(s - 'a')
	}
	return xb
}
