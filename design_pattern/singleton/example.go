package main

func main() {
	log := getLoggerInstance()
	log.Log("My first log")

	log = getLoggerInstance()
	log.setLoggerLevel("error")
	log.Log("My second log")
}
