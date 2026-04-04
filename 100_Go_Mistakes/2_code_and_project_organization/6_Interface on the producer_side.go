// Don't make a god interface on the Producer side
package store

type CustomerStorage interface {
	StoreCustomer(customer Customer) error
	GetCustomer(id string) (Customer, error)
	UpdateCustomer(customer Customer) error
	GetAllCustomers() ([]Customer, error)
	GetCustomersWithoutContract() ([]Customer, error)
	GetCustomersWithNegativeBalance() ([]Customer, error)
}

type Customer struct{}

//
package client

import "store"

type customersGetter interface {
	GetAllCustomers() ([]store.Customer, error)
}


// Instead make an Interface on the client side that has just what you need
package store

type Customer struct {
    ID   string
    Name string
}

// MySQLStore is a "Concrete Implementation"
// It has 50 methods, but it doesn't define an interface!
type MySQLStore struct {}

func (m *MySQLStore) StoreCustomer(c Customer) error { /* ... */ return nil }
func (m *MySQLStore) GetAllCustomers() ([]Customer, error) { /* ... */ return nil, nil }
func (m *MySQLStore) GetCustomersWithNegativeBalance() ([]Customer, error) { /* ... */ return nil, nil }

//
package client

import "myproject/store"

type customerLister interface { // This logic ONLY needs to list customers. It doesn't care about 'Delete' or 'Update'.
    GetAllCustomers() ([]store.Customer, error)
}

func ListAll(l customerLister) {
    list, _ := l.GetAllCustomers()
    // ... logic here
}
