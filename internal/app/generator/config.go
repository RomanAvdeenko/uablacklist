package generator

type Config struct {
	uri,
	blockedFileName,
	tplFileName,
	outFileName,
	dropIP string
}

// NewConfig is a Config create helper
func NewConfig(uri, blockedFileName, tplFileName, outFileName, dropIP string) *Config {
	return &Config{
		uri:             uri,
		blockedFileName: blockedFileName,
		outFileName:     outFileName,
		tplFileName:     tplFileName,
		dropIP:          dropIP,
	}
}
