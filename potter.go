package katapotter

func BuyHarryPotter(
	s1 int,
	s2 int,
	s3 int,
	s4 int,
	s5 int,
) float32 {

	var result float32 = 0
	// Initialize Data Structure
	requestBooks := []int{s1, s2, s3, s4, s5}

	// Scan all five books count
	for {
		var tmp float32
		totalBooks := 0
		for i := 0; i < len(requestBooks); i++ {
			if requestBooks[i] > 0 {
				totalBooks += 1
				requestBooks[i] -= 1
			}
		}

		switch totalBooks {
		case 5:
			tmp = 40 * 0.75
		case 4:
			tmp = 32 * 0.80
		case 3:
			tmp = 24 * 0.90
		case 2:
			tmp = 16 * 0.95
		case 1:
			tmp = 8
		default:
			tmp = 0
		}
		if tmp == 0 {
			break
		} else {
			result += tmp
		}
	}
	return result
}
