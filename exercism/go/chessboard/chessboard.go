package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	for fileDSG, fileValue  := range cb {	
		if file == fileDSG {
			var counter int

			for _, inFile := range fileValue {
				if inFile {
					counter += 1
				}
			} 
			return counter
		}
	} 
	return 0
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank > 8 || rank < 1 { return 0 }
	var counter int 
	for _, fileVal := range cb {
		if fileVal[rank-1] {
			counter += 1
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var squares int
	for _, file := range cb {
		squares += len(file)
	}
	return squares 
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	panic("Please implement CountOccupied()")
}
