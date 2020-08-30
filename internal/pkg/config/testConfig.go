package config

//TestConfiguration handles all the test configuration parameters
type TestConfiguration struct {
	Test test `json:"test"`
}

type test struct {
	ParallelThreads int      `json:"parallel-threads"`
	Beverages       []string `json:"beverages"`
}

var (
	testConfig TestConfiguration
)

func TestConfig() TestConfiguration {
	return testConfig
}

//ReadTestConfigFromFile reads program's test configuration from a configuration file
func ReadTestConfigFromFile(configFileName string) error {
	return readConfigFromFile(configFileName, &testConfig)
}
