func viewCheckup() checkup.Checkup {
	configBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var c checkup.Checkup
	err = json.Unmarshal(configBytes, &c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// Perform adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Perform() {
	if err := RootCmd.Perform(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
