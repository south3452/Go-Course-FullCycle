package main


const a = "Hello World!"

var (
	b bool
	c int
	d string
	e float64
) 

func main() {
	var f string = "Olá Mundo"
	println("bool", b)	
	println("int", c)	
	println("string", d)	
	println("float64", e)	
	println("escopo de var", f)	
	x()
}

func x (){
	println("escopo de var func x", a)	
}
