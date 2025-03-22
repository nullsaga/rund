package cli

import (
	"flag"
	"fmt"
)

const usage = `Usage: rund [options...]:
  -a, --addr <ip:port>  start the api server on the specified ip and port.
  -c, --config <file>   path to the configuration file to use
  -v, --verbose         increase output verbosity
  -V, --version         display version information and exit
  -h, --help            display this help text and exit
`

type Options struct {
	Addr    string
	Config  string
	Verbose bool
	Version bool
	Help    bool
}

func NewWithDefaultOptions() *Options {
	return &Options{
		Addr:    ":8080",
		Config:  "",
		Verbose: false,
		Version: false,
		Help:    false,
	}
}

func (c *Options) Parse() {
	flag.StringVar(&c.Addr, "addr", c.Addr, "")
	flag.StringVar(&c.Addr, "a", c.Addr, "")
	flag.StringVar(&c.Config, "config", c.Config, "")
	flag.StringVar(&c.Config, "c", c.Config, "")
	flag.BoolVar(&c.Help, "help", c.Help, "")
	flag.BoolVar(&c.Help, "h", c.Help, "")
	flag.BoolVar(&c.Version, "version", c.Version, "")
	flag.BoolVar(&c.Version, "V", c.Version, "")
	flag.BoolVar(&c.Verbose, "verbose", c.Verbose, "")
	flag.BoolVar(&c.Verbose, "v", c.Verbose, "")

	flag.Usage = func() {
		fmt.Print(usage)
	}

	flag.Parse()
}
