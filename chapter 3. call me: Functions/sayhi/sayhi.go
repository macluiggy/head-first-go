package main

func main() {
	SayHi()
	repeatLine("Hello", 5)
}

func SayHi() {
	println("Hello")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		println(line)
	}

}
