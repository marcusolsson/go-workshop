// +build OMIT

package defer

func loadConfig(filename string) error {
	f, err := os.Open(filename) // HL
	if err != nil {
		return err
	}

	var loaded bool

	// Try loading config

	if !loaded {
		return errors.New("unable to load config")
		// Oh no, forgot to close file :( // HL
	}

	f.Close() // HL

	return nil
}