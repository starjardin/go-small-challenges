package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	if file, ok := cb[file]; ok {
		count := 0
		for _, occupied := range file {
			if occupied {
				count++
			}
		}
		return count
	}

	return 0
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	count := 0
	for _, fileData := range cb {
		if rank <= len(fileData) {
			if fileData[rank-1] {
				count++
			}
		} else {
			return 0
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, fileData := range cb {
		count += len(fileData)
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, fileData := range cb {
		for _, occupied := range fileData {
			if occupied {
				count++
			}
		}
	}
	return count
}
