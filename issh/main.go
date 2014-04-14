package main

import (
	"github.com/jamesandariese/issh"
	"flag"
	"bytes"
	"regexp"
	"strconv"
	"fmt"
	"os"
)

var user_host_port_re, _ = regexp.Compile("(([^@]+)@)?" + 
	"([a-zA-Z0-9.-]+)" +
	"(:(6553[0-5]|" +
	"655[0-2][0-9]|" +
	"65[0-4][0-9][0-9]|" +
	"6[0-4][0-9][0-9][0-9]|" +
	"[1-5][0-9]{0,4}))?")

func parseUserHostPort(str string) (user, host string, port int) {
	matches := user_host_port_re.FindStringSubmatch(str)
	user = matches[2]
	host = matches[3]
	port, _ = strconv.Atoi(matches[5])
	return
}


var show_key = flag.Bool("K", false, "Display the public key associated with the provided arguments")

func main() {
	flag.Parse()
	if *show_key {
		if flag.NArg() < 1 {
			flag.Usage()
			os.Exit(1)
		}
		for _, seed := range flag.Args() {
			key, err := issh.GetAuthorizedKey(seed)
			if err != nil {
				panic("Failed to create a key: " + err.Error())
			}
			fmt.Printf("public key for \"%s\"\n" +
				"command=\"exit 1\",no-port-forwarding,no-X11-forwarding,no-pty %s " +
				"issh generated from seed: '%s'\n",
				seed, bytes.TrimRight(key, "\r\n \t"), seed);
		}
		return
	}

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	user, host, port := parseUserHostPort(flag.Arg(1))
	stdout, exitcode, err := issh.Run(user, host, uint16(port), flag.Arg(0))
	if err != nil {
		panic("Failed to execute remote command: " + err.Error())
	}
	
	os.Stdout.Write(stdout)
	os.Exit(exitcode)
}
