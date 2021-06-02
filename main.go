package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

//var host []map[interface{}]interface{}
var host []Hosts

func main() {
	config := config()
	var server = http.NewServeMux()
	for i := 0; i < len(config)-1; i++ {
		webContent, err := ioutil.ReadFile(config[i].file)
		if err != nil {
			log.Fatalln(err)
		}
		server = http.NewServeMux()
		server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write(webContent)
		})
		if config[i].tls {
			go http.ListenAndServeTLS(config[i].ipaddress+":"+fmt.Sprint(config[i].port), config[i].cert, config[i].key, server)
		} else {
			go http.ListenAndServe(config[i].ipaddress+":"+fmt.Sprint(config[i].port), server)
		}
	}
	lastConfigItem := len(config) - 1
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		webContent, err := ioutil.ReadFile(config[lastConfigItem].file)
		if err != nil {
			log.Fatalln(err)
		}
		rw.Write(webContent)
	})
	if config[lastConfigItem].tls {
		http.ListenAndServeTLS(config[lastConfigItem].ipaddress+":"+fmt.Sprint(config[lastConfigItem].port), config[lastConfigItem].cert, config[lastConfigItem].key, nil)
	} else {
		http.ListenAndServe(config[lastConfigItem].ipaddress+":"+fmt.Sprint(config[lastConfigItem].port), nil)
	}

}

func config() []Hosts {

	if info, err := os.Stat("config.yaml"); err != nil || info.Size() == 0 {
		log.Fatalln("Error, config file clould not be found")
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	hostsInterface := viper.Get("hosts").([]interface{})

	host = make([]Hosts, len(hostsInterface))
	for i, item := range hostsInterface {
		hostMap := item.(map[interface{}]interface{})
		host[i].file = hostMap["file"].(string)
		host[i].ipaddress = hostMap["ipaddress"].(string)
		host[i].port = hostMap["port"].(int)
		host[i].tls = hostMap["tls"].(bool)
		host[i].cert = hostMap["cert"].(string)
		host[i].key = hostMap["key"].(string)
	}
	return host
}

type Hosts struct {
	file      string
	ipaddress string
	port      int
	tls       bool
	cert      string
	key       string
}
