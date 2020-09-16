package triplets

func CompareTriplets(alice []int, bob []int) []int {
	var (
		aliceScore = 0
		bobScore   = 0
	)

	for i := 0; i < 3; i++ {
		switch {
		case alice[i] == bob[i]:
			continue
		case alice[i] > bob[i]:
			aliceScore++
		case alice[i] < bob[i]:
			bobScore++
		default:
			continue
		}
	}

	return []int{aliceScore, bobScore}
}
