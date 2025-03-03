package secClient

type boundary struct {
	start int
	end   int
}

func computeBoundaries(line string) []boundary {
	var boundaries []boundary
	inDash := false
	var start int
	for i, r := range line {
		if r == '-' {
			if !inDash {
				inDash = true
				start = i
			}
		} else {
			if inDash {
				inDash = false
				boundaries = append(boundaries, boundary{start: start, end: i})
			}
		}
	}
	if inDash {
		boundaries = append(boundaries, boundary{start: start, end: len(line)})
	}
	return boundaries
}
