package cli

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

const (
	version = "0.0.2"
)

type Options struct {
	Ip       string
	Port     int
	ConfPath string
	Verbose  bool
	Version  bool
	Help     bool
}

func NewWithDefaultOptions() *Options {
	return &Options{
		Ip:       "0.0.0.0",
		Port:     8080,
		ConfPath: "",
		Verbose:  false,
		Version:  false,
		Help:     false,
	}
}

func (c *Options) Parse() {
	flag.StringVarP(&c.Ip, "ip", "i", c.Ip, "`<ip>` address to listen on")
	flag.IntVarP(&c.Port, "port", "p", c.Port, "bind `<port>` to listen on")
	flag.StringVarP(&c.ConfPath, "conf", "c", c.ConfPath, "`<path>` to the configuration file")
	flag.BoolVarP(&c.Help, "help", "h", c.Help, "display this help text and exit")
	flag.BoolVarP(&c.Version, "version", "V", c.Version, "display version information and exit")
	flag.BoolVarP(&c.Verbose, "verbose", "v", c.Verbose, "increase output verbosity")
	flag.Parse()

	if c.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	if c.Help {
		flag.Usage()
		os.Exit(0)
	}
}
