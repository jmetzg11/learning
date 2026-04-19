package main 

const (
	StatusSuccess = "success"
	StatusErrorFoo = "error_foo"
	StatusErrorBar = "error_bar"
)

// Error: notify and incrementCounter are always called with an empty string!
func f() error { 
	var status string 
	defer notify(status) // argument evaluated right away 
	defer incrementCounter(status) // argument evaluated right away 
	
	if err := foo(); err != nil {
		status = StatusErrorFoo 
		return err 
	}
	
	if err := bar(); err!= nil {
		status = StatusErrorBar 
		return err
	}
	
	status = StatusSuccess 
	return nil
 }

// Solutions
func f1()  error {
	var status string 
	defer notify(&status)
	defer incrementCounter(&status)
	
	if err := foo(); err != nil {
		status = StatusErrorFoo 
		return err 
	}
	
	if err := bar(); err!= nil {
		status = StatusErrorBar 
		return err
	}
	
	status = StatusSuccess 
	return nil
}

// closures 
func main() {
	i := 0 
	j := 0 
	defer func(i int) {
		fmt.Println(i, j)
		// 0, 1 
	}(i) // evaluated right away
	i++
	j++
}

func f2()  error {
	var status string 
	// values from outside its body 
	defer func() {
		notify(status) 
		incrementCounter(status)
	}()
	
	if err := foo(); err != nil {
		status = StatusErrorFoo 
		return err 
	}
	
	if err := bar(); err!= nil {
		status = StatusErrorBar 
		return err
	}
	
	status = StatusSuccess 
	return nil
}


// pointer and value receivers 

func main() {
	s := Struct{ id: "foo" }
	defer s.print() // s evaluated immediately 
	// foo
	s.id = "bar"
}

type Struct struct {
	id string 
}

func (s Struct) print() {
	fmt.Println(s.id)
}

func main() {
	s := &Struct{ id: "foo" }
	defer s.print() // s evaluated immediately 
	// bar
	s.id = "bar"
}

type Struct struct {
	id string 
}

func (s *Struct) print() {
	fmt.Println(s.id)
}