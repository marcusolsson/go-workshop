// +build OMIT

package defer

func loadConfig(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // HL

	var loaded bool

	// Try loading config

	if !loaded {
		return errors.New("unable to load config")
	}

	return nil
}