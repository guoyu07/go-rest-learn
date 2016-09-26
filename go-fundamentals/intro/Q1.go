package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	i := 0
Tag:
	println(i)
	if i < 10 {
		i++
		goto Tag
	}
}
