package uablacklist

type Config struct {
	url,
	tplFileName,
	outFileName string
}

// NewConfig is a Config create helper
func NewConfig(url, tplFileName, outFileName string) *Config {
	return &Config{
		url:         url,
		outFileName: outFileName,
		tplFileName: tplFileName,
	}
}
