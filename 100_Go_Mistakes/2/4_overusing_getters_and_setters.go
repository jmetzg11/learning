// If ther is no Interface, no reason to make setter/getters
type Person struct {
	Name    string
	Savings float64 // Exported field
}

func main() {
	p := Person{Name: "Alice", Savings: 1000}

	// Direct access is clean and readable
	p.Savings += 500
	fmt.Println(p.Name, "has", p.Savings)
}

// We have an Interface and we need the setter/getters
type Taxable interface {
	Value() float64
	SetValue(amt float64)
}

type Person struct {
	Name    string
	Savings float64
}

func (p *Person) Value() float64 {
	return p.Savings
}

func (p *Person) SetValue(amt float64) {
	if amt < 0 {
		p.Savings = 0
		return
	}
	p.Saving = amt
}

type Company struct {
	TaxID        string
	LiquidAssets float64
}

func (c *Company) Value() float64 {
	return c.LiquidAssets
}

func (c *Company) SetValue(amt float64) {
	if amt < 0 {
		c.LiquidAssets = 0
		return
	}
	c.LiquidAssets = amt
}

func ApplyTaxRelief(t Taxable, reliefAmount float64) {
	newBalance := t.Value() - reliefAmount
	if newBalance < 0 {
		newBalance = 0 // Business logic: don't let relief cause debt
	}
	t.SetValue(newBalance)
}

func main() {
	// Note: We use & (pointers) because the interface methods use * pointer receivers
	p := &Person{Name: "Alice", Savings: 1000}
	c := &Company{TaxID: "99-123", LiquidAssets: 50000}

	ApplyTaxRelief(p, 100)
	ApplyTaxRelief(c, 5000)

	fmt.Printf("%s new savings: %.2f\n", p.Name, p.Value())
	fmt.Printf("Company %s new assets: %.2f\n", c.TaxID, c.Value())
}
