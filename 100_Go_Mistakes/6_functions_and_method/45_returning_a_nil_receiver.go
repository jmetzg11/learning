package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	var foo *Foo
	fmt.Println(foo.Bar())

	customer := Customer{Age: 33, Name: "John"}
	if err := customer.Validate(); err != nil {
		log.Printf("customer is invalid: %v", err)
	}
	
	if err := customer.Validate1(); err != nil {
		log.Printf("customer is invalid: %v", err)
	}
}

type Foo struct{}

func (foo *Foo) Bar() string {
	return "bar"
}

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("Name is nil"))
	}

	return m
}

func (c Customer) Validate1() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("Name is nil"))
	}
	
	if m != nil {
		return m
	}

	return nil
}
