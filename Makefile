run:
	go build
	mv main simulator
	./simulator
test:
	go test -run TestTokenizerAndStateNames
	go test -run TestParser
	go test -run TestTruthTable
	go test -run TestLogicalEquivalenceCalculator
clean:
	rm simulator
