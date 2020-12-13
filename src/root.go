package godo

// StartApp configures the starting and running of the application
func StartApp() error {
	if err := InitializeLogging(); err != nil {
		return err
	}

	conn, err := InitializeNotesBackend()
	if err != nil {
		return err
	}

	config, err := GetConfig()
	if err != nil {
		return err
	}

	_, err = ConfigureUI(conn, &config)
	if err != nil {
		return err
	}

	return nil
}

func makeStringChannel() chan string {
	return make(chan string)
}

func makeIntChannel() chan int {
	return make(chan int)
}
