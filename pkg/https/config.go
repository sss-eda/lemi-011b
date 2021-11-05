package https

import (
	"flag"
	"strings"
)

// Config TODO
type Config struct {
	Hosts hostsVar
}

func configure() Config {
	config := Config{}
	flag.StringVar(&config.Hosts, "addr", "")
}

type hostsVar []string

// String TODO
func (hosts *hostsVar) String() string {
	if len(hosts) > 0 && hosts != nil {
		return strings.Join(hosts, ", ")
	}

	return ""
}

// Set TODO
func (hosts *hostsVar) Set(s string) error {

}
