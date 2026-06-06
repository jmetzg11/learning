
ticker := time.NewTicker(1000) // nanoseconds
for {
	select {
	case <-ticker.C:
		// Do something
	}
}

// Better
ticker := time.NewTicker(time.MicroSeconds)
ticker := time.NewTicer(1000 * time.Nanosecond)
