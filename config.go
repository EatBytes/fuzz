package razboy

type Config struct {
	Url         string
	Method      string
	Parameter   string
	Key         string
	Proxy       string
	Raw         bool
	Shellmethod string
}

func NewConfig() *Config {
	c := new(Config)
	c.Method = "GET"
	c.Parameter = "razboynik"
	c.Shellmethod = "system"
	c.Key = ""

	return c
}
