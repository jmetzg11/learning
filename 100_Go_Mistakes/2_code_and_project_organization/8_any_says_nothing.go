// Avoid overgeneralizing the code we write at all costs. Perhaps a little bit of duplicated code might occasionally be better if it improves other aspects such as code expressiveness. Any can be helpful if there is a genuine need for accepting or returning any possible type (for instance, when it comes to marshaling or formatting).

// Before
package store

type Customer struct {
	// Some fields
}

type Contract struct {
	// Some fields
}

type Store struct{}

func (s *Store) Get(id string) (any, error) {
	return nil, nil
}

func (s *Store) Set(id string, v any) error {
	return nil
}

// After
package store
type Customer struct {
	// Some fields
}

type Contract struct {
	// Some fields
}

type Store struct{}

func (s *Store) GetContract(id string) (Contract, error) {
	return Contract{}, nil
}

func (s *Store) SetContract(id string, contract Contract) error {
	return nil
}

func (s *Store) GetCustomer(id string) (Customer, error) {
	return Customer{}, nil
}

func (s *Store) SetCustomer(id string, customer Customer) error {
	return nil
}
