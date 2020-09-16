package rot

func RotateLeft(array []int, rotations int) []int {
	// Create a local copy of array
	result := array

	// mutate the copied array
	for i := 0; i < rotations; i++ {
		targetRot := result[0]
		chunk := result[1:]
		result = append(chunk, targetRot)
	}

	return result
}
