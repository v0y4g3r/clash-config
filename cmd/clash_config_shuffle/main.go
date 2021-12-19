package main

import (
	"fmt"
	"github.com/RayneHwang/clash-config/pkg/defs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

func RemoveIndex(s []defs.ProxyGroup, index int) []defs.ProxyGroup {
	return append(s[:index], s[index+1:]...)
}

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

	var afterRemove []defs.ProxyGroup
	for idx, _ := range c.ProxyGroups {
		if c.ProxyGroups[idx].Name != "其他站点" {
			afterRemove = append(afterRemove, c.ProxyGroups[idx])
		}
	}

	c.ProxyGroups = afterRemove

	for idx, _ := range c.ProxyGroups {

		if c.ProxyGroups[idx].Name == "PROXY" || c.ProxyGroups[idx].Name == "NF油管" {
			c.ProxyGroups[idx].Type = "url-test"
			c.ProxyGroups[idx].Url = "http://www.gstatic.com/generate_204"
			c.ProxyGroups[idx].Interval = 180
			rand.Shuffle(len(c.ProxyGroups[idx].Proxies), func(i, j int) {
				c.ProxyGroups[idx].Proxies[i], c.ProxyGroups[idx].Proxies[j] = c.ProxyGroups[idx].Proxies[j], c.ProxyGroups[idx].Proxies[i]
			})
			fmt.Println(c.ProxyGroups[idx])
		}
	}

	for idx := range c.Rules {
		if c.Rules[idx] == "MATCH,,其他站点" {
			c.Rules[idx] = "MATCH,,PROXY"
		}
	}
	d, _ := yaml.Marshal(&c)
	err = ioutil.WriteFile(targetFile, d, 0644)
	if err != nil {
		fmt.Printf("Failed to write file [%v], error:[%v", targetFile, err)
	}
}
