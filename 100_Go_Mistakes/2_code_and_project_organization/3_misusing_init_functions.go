// Used to initialize the state of an application.
// No arguments or results


// Redis init() -> Main init() -> main()
// will be called in alphabetical order; a.go -> b.go
// Main package
package main

import (
	"fmt"

	"github.com/teivah/100-go-mistakes/src/02-code-project-organization/3-init-functions/redis"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	err := redis.Store("foo", "bar")
	_ = err
}
// Redit package
package redis

import "fmt"

func init() {
	fmt.Println("redis")
}

func Store(key, value string) error {
	return nil
}

// side effects
import (
	"fmt"
	_ "foo" // not used but init() for foo will be called
)


// BAD EXAMPLE

// error messages in init() panic and stop app, no chance for client to handle error
// complicates unit testing. init() will always be called
// shouldn't assign to a global variable, variable can reassigned and isn't truely isolated
package main

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func init() {
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}
	db = d
}

// GOOD EXAMPLE
// error handling left to the caller
// possible to create integration tests
// connection pool is encapsulated within the function

func createClient(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// populating global routing table before server starts, DefaultServeMux
func init() {
	// Redirect "/blog/" to "/", because the menu bar link is to "/blog/"
	// but we're serving from the root.
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	http.HandleFunc("/blog", redirect)
	http.HandleFunc("/blog/", redirect)

	// Keep these static file handlers in sync with app.yaml.
	static := http.FileServer(http.Dir("static"))
	http.Handle("/favicon.ico", static)
	http.Handle("/fonts.css", static)
	http.Handle("/fonts/", static)

	http.Handle("/lib/godoc/", http.StripPrefix("/lib/godoc/", http.HandlerFunc(staticHandler)))
}

