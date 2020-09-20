package configs

type config struct {
	base string
	URL string
	PORT string
}

// MongoConfig returns the config option
func MongoConfig() config  {
	return config{
		base: "mongodb",
		URL: "localhost",
		PORT: "27017",
	}
}

// GetConnectionString returns the connection string
func (c *config) GetConnectionString() string {
	return c.base + "://" +c.URL + ":" + c.PORT
}