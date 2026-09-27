package main

type (
	Binary uint8 // i.e. 0 or 1
	State  string
	States map[State][]Binary // 'a' or "a'" -> [0,1,0,0,1]
	Term   []Token
)

func PopulatesStateBins(tokens []Token, stateNames []State) (states States) {
	numberOfStates := len(stateNames)
	if numberOfStates == 0 {
		return make(States)
	}
	numberOfRows := 1 << (numberOfStates / 2)
	states = make(States, numberOfStates)
	binaries := getAllBinaryRows(numberOfRows, numberOfStates)
	for row, bins := range binaries {
		for col, bin := range bins {
			if _, exists := states[stateNames[col]]; !exists {
				states[stateNames[col]] = make([]Binary, numberOfRows)
			}
			states[stateNames[col]][row] = bin
		}
	}
	return states
}

func getAllBinaryRows(numberOfRows int, numberOfStates int) [][]Binary {
	numVars := numberOfStates / 2
	binaryRows := make([][]Binary, numberOfRows)
	for i := 0; i < numberOfRows; i++ {
		bins := make([]Binary, numberOfStates)
		for k := 0; k < numVars; k++ {
			bit := Binary((i >> (numVars - 1 - k)) & 1)
			bins[k] = bit
			bins[k+numVars] = bit ^ 1
		}
		binaryRows[i] = bins
	}
	return binaryRows
}
