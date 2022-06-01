package main

func main() {
	sayHi()
	repeatLine("Hello", 5)
}

func sayHi() {
	println("Hello")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		println(line)
	}

}
