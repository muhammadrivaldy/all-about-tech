package main

func main() {
	reqTest := requestTesting{
		a: 1,
		b: "school",
		c: "street of school",
		d: 5,
		e: 123,
	}

	testing(reqTest)
}

type requestTesting struct {
	a int         // required
	b string      // required
	c string      // required
	d int         // required
	e int         // required
	f interface{} // optional (new parameter)
}

func testing(req requestTesting) {
	// code for testing purpose
}
