package main

import (
	"log"

	"github.com/matsuev/klsh-email-sender/internal/config"
	"github.com/matsuev/klsh-email-sender/internal/parser"
	"github.com/matsuev/klsh-email-sender/internal/template"
)

func main() {
	println("KLSH Email sender")

	// Create application config
	cfg, err := config.Create()
	if err != nil {
		log.Fatalln(err)
	}

	// Create CSV file parser
	csv, err := parser.Create(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer csv.Close()

	// Create template file parser
	tpl, err := template.Create(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	// Check CSV and template keys
	if !tpl.CheckKeys(csv.Keys) {
		log.Fatalln("Keys error")
	}

	// for csv.Scan() {
	// 	m, err := csv.GetLine()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println(m)
	// }

	// fmt.Printf("%#v\n", cfg)
	// fmt.Printf("%#v\n", csv)

	// s := "abcd"
	// ss := strings.TrimPrefix(strings.TrimSuffix(s, "\""), "\"")
	// fmt.Println(ss)
}
