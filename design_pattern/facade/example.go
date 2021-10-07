// Purpose:
// Provide a simple, front facing interface to a more complex system, library or API

// Scenarios:
// Improve usability of a more complex API
// Serve as a starting point of refactoring

package main

func main() {
	// 使用的時候，不需要知道 coffeeMachine 底層的邏輯，可以更簡潔的使用高階邏輯
	makeEspresso()
	makeAmericano()
}
