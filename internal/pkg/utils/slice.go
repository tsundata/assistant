package utils

func StringSliceDiff(s1, s2 []string) []string {
	if len(s1) == 0 {
		return s2
	}
	mb := make(map[string]struct{}, len(s2))
	for _, x := range s2 {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range s1 {
		if _, ok := mb[x]; !ok {
			diff = append(diff, x)
		}
	}
	return diff
}
