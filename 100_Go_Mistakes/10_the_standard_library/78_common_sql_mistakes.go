// sql.Open doesn't necessarily establish connection to a database
db, err := sql.Open("mysql", dsn)
if err != nil {
	return err
}
// depends on the dql driver being used
// sometimes we want to make sure all the services are ready before sending traffic
// but sometimes dbs open lazily
// send a ping
db, err := sql.Open("mysql", dsn)
if err != nil {
	return err
}
if err := db.Ping(); err != nil { // establishes a connection
	return err
}


// Fogetting about connection pooling
// sql.Open returns *sql.DB, it represents a pool of connections
// a connection pool can have two states: already used, idele
// 4 config parameters

// SetMaxOpensConns - maximum number of open connections
// 	make sure it fits what the underlying database can handle
// SetMaxIdelConns - Maximum numer of idle connections
// 	should be increased if the app receives a lot of concurrent requests
// SetConnMaxIdelTimer - Maximum amount of time a connection can be idle before closing
// 	if the app receives bursts of requests, so that idle connections will close
// SetConnMaxLifetime - Maximum amount of time a connection can be held open before it's closed
// 	can be helpful if we connect to load balancing db, so no connection is held too long

// we can also use multiple connection pools if an application faces different use cases


// Not using prepared statements
// Efficiency - the statement doesn't have to be recompiled
// Security - This approach reduces the risks of SQL injection attackes

// any statement that is repeated we should use prepared statements.

stmt, err := db.Prepare("SELECT i FROM ORDER WHERE ID = ?")
if err != nil {
	return err
}
rows, err := stmt.Query(id)
stmt.Close()


// Mishandling null values
rows, err := db.Query("SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
if err != nil {
	return err
}
defer rows.Close()

var (
	department string
	age int
)
for rows.Next() {
	err := rows.Scan(&department, &age)
	if err != nil {
		return err
	}
}

// if DEP is nullable we can get errors

// solution 1, provide the address of a string type direction, if DEP is NULL then it will be nill
var (
	department *string
	age int
)
for rows.Next() {
	err := rows.Scan(&department, &age)
}

// solution 2, use sql.NullXXX
var (
	department sql.NullString
	age int
)
for rows.Next() {
	err := rows.Scan(&department, &age)
}

// there is also Bool, Int32, Int64, Float64 and NullTime for sql.NullXXX


// Not handling row iteration errors
func get(ctx context.Context, db *sql.DB id string) (string, int, error) {
	rows, err := db.QueryContext(ctx,
		"SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return "", 0, err
	}
	// anonymous closure function
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	var (
		department string
		age int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return "", 0, err
		}
	}

	// Errors can happend during iteration
	// Either no more rows or error while preparing the next row
	if err := rows.Err(); err != nil {
		return "", 0, err
	}
	return department, age, nil
}
