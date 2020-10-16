package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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

	targetFile := os.Args[1]

	fmt.Printf("Read from %s", targetFile)

	c := defs.Config{}

	file, err := ioutil.ReadFile(targetFile)

	if err != nil {
		fmt.Printf("Failed to read file, err: %v", err)
		return
	}

	err = yaml.Unmarshal(file, &c)
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

	d, _ := yaml.Marshal(&c)
	err = ioutil.WriteFile(targetFile, d, 0644)
	if err != nil {
		fmt.Printf("Failed to write file [%v], error:[%v", targetFile, err)
	}
}
