package main

type (
	Binary uint8 // i.e. 0 or 1
	State  any
	States map[State][]Binary // 'a' or "a'" -> [0,1,0,0,1]
)

func StateManager() {
}
