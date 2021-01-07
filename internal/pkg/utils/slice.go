package utils

func SliceDiff(s1, s2 []string) []string {
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
