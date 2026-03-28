// Common behavior
type Notifier interface {
	Send(message string) error
}

type Email struct {
	Address string
}

func (e Email) Send(msg string) error {
	fmt.Printf("Sending Email to %s: %s\n", e.Address, msg)
	return nil
}

type SMS struct {
	PhoneNumber string
}

func (s SMS) Send(msg string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, msg)
	return nil
}

func AlertUser(n Notifier, message string) {
	err := n.Send(message)
	if err != nil {
		fmt.Println("Failed to send alert")
	}
}

func main() {
	myEmail := Email{Address: "alice@example.com"}
	myPhone := SMS{PhoneNumber: "555-0101"}

	AlertUser(myEmail, "Your package has arrived!")
	AlertUser(myPhone, "Your package has arrived!")

	alerts := []Notifier{myEmail, myPhone}
	for _, a := range alerts {
		a.Send("System Maintenance at Midnight")
	}
}

// Decoupling
type Storer interface {
	Save(data string)
}

type Disk struct {
	Path string
}

func (d Disk) Save(data string) {
	fmt.Printf(" [DISK] Writing to %s: %s\n", d.Path, data)
}

type Cloud struct {
	Region string
}

func (c Cloud) Save(data string) {
	fmt.Printf(" [CLOUD] Uploading to %s: %s\n", c.Region, data)
}

func PerformArchive(s Storer, content string) {
	fmt.Println("Starting archive process...")
	s.Save(content) // Go dynamically calls the correct method here!
}

func getStoreType(useCloud bool) Storer {
	if useCloud {
		return Cloud{Region: "us-east-1"}
	}
	// Otherwise, return a local disk
	return Disk{Path: "/tmp/backup"}
}

func main() {
	// Imagine this comes from an environment variable like os.Getenv("USE_CLOUD")
	envSetting := true

	// 'storer' is of TYPE Storer (the interface).
	// It doesn't know it's a 'Cloud' struct yet.
	storer := getStoreType(envSetting)

	// Because storer is a 'Storer', we can call Save() on it.
	// Go "dynamically" dispatches this to Cloud.Save() at runtime.
	storer.Save("System Backup")
}

// Restricting behavior
type Reader interface {
	Read() string
}

type Document struct {
	Content string
}

func (d *Document) Read() string {
	return d.Content
}

func (d *Document) Write(newText string) {
	d.Content = newText
	fmt.Println("[DOC] Content updated.")
}

func (d *Document) Delete() {
	d.Content = ""
	fmt.Println("[DOC] Content wiped!")
}

func DisplayStatus(r Reader) { // impossible to do .Write() or .Delete()
	fmt.Println("Current Content:", r.Read())

	// ❌ THIS WILL NOT COMPILE:
	// r.Write("Hacked!")
	// Error: r.Write undefined (type Reader has no field or method Write)
}

func main() {
	doc := &Document{Content: "Top Secret Data"}

	// Inside main, we can do anything
	doc.Delete()
	doc.Write("Updated Secret Data")

	// When we pass it here, we "downgrade" its permissions
	DisplayStatus(doc)
}

// Interface Pollution - interfaces create abstractions. Abstractions should be discovered, not created.
