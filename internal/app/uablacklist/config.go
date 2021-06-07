package uablacklist

type Config struct {
	url,
	tplFileName,
	outFileName,
	dropURL string
}

// NewConfig is a Config create helper
func NewConfig(url, tplFileName, outFileName, dropURL string) *Config {
	return &Config{
		url:         url,
		outFileName: outFileName,
		tplFileName: tplFileName,
		dropURL:     dropURL,
	}
}
