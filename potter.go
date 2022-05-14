package katapotter

func BuyHarryPotter(
	s1 int,
	s2 int,
	s3 int,
	s4 int,
	s5 int,
) float32 {

	totalBooks := s1 + s2 + s3 + s4 + s5
	fullPrice := float32(totalBooks) * 8
	if totalBooks >= 5 {
		return fullPrice * 0.75
	} else if totalBooks >= 4 {
		return fullPrice * 0.80
	} else if totalBooks >= 3 {
		return fullPrice * 0.90
	} else if totalBooks >= 2 {
		return fullPrice * 0.95
	} else {
		return fullPrice
	}
}
