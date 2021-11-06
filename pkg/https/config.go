package https

// Config TODO
type Config struct {
	Hosts   []string
	CertDir string
}

// Config TODO
// type Config struct {
// 	Hosts hostsVar
// 	Certs string
// }

// func configure() Config {
// 	config := Config{}
// 	flag.StringVar(&config.Hosts, "addr", "")
// 	flag.StringVar(&config.Hosts, "a", "")
// 	flag.StringVar(&config.Certs, "certs", "")
// }

// type hostsVar []string

// // String TODO
// func (hosts *hostsVar) String() string {
// 	if len(*hosts) > 0 && hosts != nil {
// 		return strings.Join(*hosts, ", ")
// 	}

// 	return ""
// }

// // Set TODO
// func (hosts *hostsVar) Set(s string) error {
// 	return nil
// }
