package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/iimeru/go-guerrilla"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	configure()
	goguerrilla.Run(gConfig)

	for i := 0; i < 3; i++ {
		go saveMail()
	}
}

func saveMail() {
	mail := <-goguerrilla.ProcessMailChan
	fmt.Println(mail.Subject)
}

var gConfig = map[string]string{
	"GSMTP_MAX_SIZE":         "131072",
	"GSMTP_HOST_NAME":        "server.example.com", // This should also be set to reflect your RDNS
	"GSMTP_VERBOSE":          "Y",
	"GSMTP_LOG_FILE":         "",    // Eg. /var/log/goguerrilla.log or leave blank if no logging
	"GSMTP_TIMEOUT":          "100", // how many seconds before timeout.
	"GSTMP_LISTEN_INTERFACE": "0.0.0.0:25",
	"GSMTP_PUB_KEY":          "/etc/ssl/certs/ssl-cert-snakeoil.pem",
	"GSMTP_PRV_KEY":          "/etc/ssl/private/ssl-cert-snakeoil.key",
	"GM_ALLOWED_HOSTS":       "guerrillamail.de,guerrillamailblock.com",
	"GM_PRIMARY_MAIL_HOST":   "guerrillamail.com",
	"GM_MAX_CLIENTS":         "500",
	"NGINX_AUTH_ENABLED":     "N",              // Y or N
	"NGINX_AUTH":             "127.0.0.1:8025", // If using Nginx proxy, ip and port to serve Auth requsts
	"SGID":                   "1008",           // group id
	"SUID":                   "1008",           // user id, from /etc/passwd
}

func configure() {
	var configFile, verbose, iface string
	// parse command line arguments
	flag.StringVar(&configFile, "config", "goguerrilla.conf", "Path to the configuration file")
	flag.StringVar(&verbose, "v", "n", "Verbose, [y | n] ")
	flag.StringVar(&iface, "if", "", "Interface and port to listen on, eg. 127.0.0.1:2525 ")
	flag.Parse()
	// load in the config.
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalln("Could not read config file")
	}
	var myConfig map[string]string
	err = json.Unmarshal(b, &myConfig)
	if err != nil {
		log.Fatalln("Could not parse config file")
	}
	for k, v := range myConfig {
		gConfig[k] = v
	}
	gConfig["GSMTP_VERBOSE"] = strings.ToUpper(verbose)

	if len(iface) > 0 {
		gConfig["GSTMP_LISTEN_INTERFACE"] = iface
	}
}
