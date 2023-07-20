package util

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ovh/go-ovh/ovh"
)

const (
	TITLE_SECTION = "==="
	END_SECTION   = "======"
)

// PrintIfNotEmpty prints if not empty, an interface array
func PrintIfNotEmpty(name string, objects []interface{}) {
	if len(objects) == 0 {
		return
	}
	var b strings.Builder
	b.WriteString(name)
	b.WriteString("\t")
	for _, obj := range objects {
		fmt.Fprintf(&b, "%s\t", obj)
	}
	fmt.Println(b.String())
}

// PrintIfNotEmpty prints if not empty, a string array
func PrintIfNotEmptyS(name string, s string) {
	if len(s) == 0 {
		return
	}
	fmt.Println(s)
}

// PrintAssets prints all available services of a given API URL
func PrintAssets(url, name string, client *ovh.Client) {
	var assets []interface{}
	err := client.Get(url, &assets)
	if err != nil {
		log.Fatal("client.Get("+url+"):", err)
	}
	PrintIfNotEmpty(name+":", assets)
	time.Sleep(100 * time.Millisecond)
}

func PrintTitle(name string) {
	fmt.Println(TITLE_SECTION + " " + name + " " + TITLE_SECTION)
}
func PrintEndSection() {
	fmt.Println(END_SECTION)
}
