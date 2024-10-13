run:
	go build
	./main
test:
	go test -run TestTokenizerAndStateNames
	go test -run TestParser
	go test -run TestTruthTable
	go test -run TestLogicalEquivalenceCalculator
clean:
	rm main
