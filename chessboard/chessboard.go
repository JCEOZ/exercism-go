package chessboard

type File []bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0

	for key, files := range cb {
		for _, occupied := range files {
			if key == file && occupied {
				count++
			}
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank < 1 || rank > 8 {
		return count
	}

	for _, files := range cb {
		for i, occupied := range files {
			if i == rank-1 && occupied {
				count++
			}
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0

	for _, files := range cb {
		count += len(files)
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

	for _, files := range cb {
		for _, occupied := range files {
			if occupied {
				count++
			}
		}
	}

	return count
}
