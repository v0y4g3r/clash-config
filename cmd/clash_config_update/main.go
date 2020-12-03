package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"clash-config/pkg/defs"

	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "url cannot be empty. Params: clash_config_update <url> [<targetFile>]")
		return
	}

	url := os.Args[1]

	var targetFile string
	if len(os.Args) == 2 {
		targetFile = "./config.yaml"
	} else {
		targetFile = os.Args[2]
	}

	fmt.Printf("Read from %s and output to %s", url, targetFile)

	c := defs.Config{}
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Failed to send request to url %v", url)
		return
	}

	defer resp.Body.Close()

	if err != nil {
		panic(fmt.Sprintf("failed to request config file %v", url))
	}

	body, _ := ioutil.ReadAll(resp.Body)

	err = yaml.Unmarshal(body, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	rand.Seed(time.Now().UnixNano())

	for idx, _ := range c.ProxyGroups {
		if c.ProxyGroups[idx].Name == "PROXY" {
			c.ProxyGroups[idx].Type = "url-test"
			c.ProxyGroups[idx].Url = "http://www.gstatic.com/generate_204"
			c.ProxyGroups[idx].Interval = 180
			rand.Shuffle(len(c.ProxyGroups[idx].Proxies), func(i, j int) {
				c.ProxyGroups[idx].Proxies[i], c.ProxyGroups[idx].Proxies[j] = c.ProxyGroups[idx].Proxies[j], c.ProxyGroups[idx].Proxies[i]
			})
			fmt.Println(c.ProxyGroups[idx])
		}
	}

	// fmt.Println(d)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	c.Rules = append([]string{"DOMAIN-SUFFIX,alipay.com,DIRECT"}, c.Rules...)
	c.Rules = append([]string{"DOMAIN-SUFFIX,antfin-inc.com,DIRECT"}, c.Rules...)
	c.Rules = append([]string{"DOMAIN-SUFFIX,alipay.net,DIRECT"}, c.Rules...)
	c.Rules = append([]string{"DOMAIN-SUFFIX,alibaba-inc.com,DIRECT"}, c.Rules...)
	c.Rules = append([]string{"DOMAIN-SUFFIX,antfin.com,DIRECT"}, c.Rules...)
	c.Rules = append([]string{"DOMAIN-SUFFIX,alipay-inc.com,DIRECT"}, c.Rules...)
	c.Rules = append([]string{"DOMAIN-SUFFIX,atatech.org,DIRECT"}, c.Rules...)
	d, _ := yaml.Marshal(&c)
	// fmt.Println(string(d))
	ioutil.WriteFile(targetFile, d, 0644)
}
