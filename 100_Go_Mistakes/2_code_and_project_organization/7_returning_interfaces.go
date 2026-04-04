// returning interface is locking the client in a box, is there a reason to do that?
// Be conservative in what you do, be liberal in what you accept from others
// If a client needs to abstract an implmentation for whatever reason, it can still do that on the client's side
// avoids circula dependencies

// Bad - restricts what the client do with the return object
package store

type CustomerStore interface {
	Get(id string) Customer
}

type mysqlStore struct{} // private!

func (m *mysqlStore) Get(id string) Customer { return Customer{} }

// ❌ MISTAKE: The function returns the INTERFACE
func NewMySQLStore() CustomerStore {
	return &mysqlStore{}
}

// Good
package store

// MySQLStore is public/exported
type MySQLStore struct{}

func (m *MySQLStore) Get(id string) Customer { return Customer{} }
func (m *MySQLStore) Close() error { return nil } // A bonus method!

// ✅ CORRECT: The function returns the concrete STRUCT
func NewMySQLStore() *MySQLStore {
    return &MySQLStore{}
}
