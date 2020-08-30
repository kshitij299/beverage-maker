package config

//Configuration handles all the configuration parameters
type Configuration struct {
	Machine machine `json:"machine"`
}

type machine struct {
	Outlets     outlets                   `json:"outlets"`
	Ingredients map[string]int            `json:"total_items_quantity"`
	Beverages   map[string]map[string]int `json:"beverages"`
}

type outlets struct {
	Count int `json:"count_n"`
}

var (
	config Configuration
)

func Config() Configuration {
	return config
}

//ReadConfigFromFile reads program configuration from a configuration file
func ReadConfigFromFile(configFileName string) error {
	return readConfigFromFile(configFileName, &config)
}
