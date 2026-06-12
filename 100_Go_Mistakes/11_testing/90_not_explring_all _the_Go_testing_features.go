// Code Coverage 
// go test -coverprofile=coverage.out ./...
// go tool cover -html=coverage.out 


// Testing from a different package 

// focus on what is exposed to the client and not implementation 

// all files files in a folder should be belong to the same package,
// with only one exception: a test file can belong to a _test package 
// this makes it easier to test exposed behavior 


// Utility functions 
func TestCustomer(t *testing.T) {
	customer, err := createCustomer("foo")
	if err != nil {
		t.Fatal(err)
	}
	// ...
}

func CreateCustomer(someArg string) (Customer, error) {
	// create customer
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}

// Better, less code and better error management 
// raise the error immediately 
func TestCustomer(t *testing.T) {
	customer := createCustomer(t, "foo")
}

fund createCustomer(t *testing.T, someArg string) Customer {
	// create customer 
	if err != nil {
		t.Fatal(err)
	}
	return customer
}


// Setup and teardown
// setup per test
func TestMySQLIntegration(t *testing.T) {
	setupMySQL()
	defer teardownMySQL
}

// function to be called at the end of a test
func TestMySQLIntegration(t *testing.T) {
	// ...
	db := createConnection(t, "tcp(localhost:3306)/db")
	// ...
}

func createConncetion(t *testing.T, dns string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.FailNow()
	}
	t.Clanup ( // register to be used at the end of the test
		func() {
			_ = db.Close()
		})
	return db
}

// Setup and teardown per package 
func TestMain(m *testing.M) {
	setupMySQL()
	code := m.Run()
	teardownMySQL()
	os.Exit(m.Run())
}