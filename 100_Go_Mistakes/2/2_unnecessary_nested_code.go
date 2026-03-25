// reduce the number o fnested blocks, aligning the happy path on the left, and returning as early as possible -> better readability

// Too nested - no mental model
func join1(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

// Easier to follow the logic, align the happy path to the left; you should quickly be able to scan down one column to see the xpected execution flow.
func join2(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return "", nil
}

// Bad
if foo() {
	// ...
	return true
} else {
	// ...
}

// Good, always omit the else
if foo() {
	return true
}
// ..

// Bad
if s != "" {
	// ...
} else {
	return errors.New("empty string")
}

// Good
if s == "" {
	return errors.New("empty string")
}
// ...
