package katapotter

func BuyHarryPotter(
	s1 int,
	s2 int,
	s3 int,
	s4 int,
	s5 int,
) float32 {

	var result float32 = 0

	// Scan all five books count
	for {
		var tmp float32
		totalBooks := 0
		// Count Books
		if s1 > 0 {
			totalBooks += 1
			s1 -= 1
		}

		if s2 > 0 {
			totalBooks += 1
			s2 -= 1
		}

		if s3 > 0 {
			totalBooks += 1
			s3 -= 1
		}

		if s4 > 0 {
			totalBooks += 1
			s4 -= 1
		}

		if s5 > 0 {
			totalBooks += 1
			s5 -= 1
		}

		if totalBooks >= 5 {
			tmp = 40 * 0.75
		} else if totalBooks >= 4 {
			tmp = 32 * 0.80
		} else if totalBooks >= 3 {
			tmp = 24 * 0.90
		} else if totalBooks >= 2 {
			tmp = 16 * 0.95
		} else {
			tmp = float32(totalBooks) * 8
		}
		if tmp == 0 {
			break
		} else {
			result += tmp
		}
	}

	return result
}
