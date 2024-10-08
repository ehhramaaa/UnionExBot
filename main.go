package main

import (
	"fmt"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"

	"agent301/core"

)

func main() {
	fmt.Println(`

	/$$   /$$           /$$                     /$$$$$$$$          
	| $$  | $$          |__/                    | $$_____/          
	| $$  | $$ /$$$$$$$  /$$  /$$$$$$  /$$$$$$$ | $$       /$$   /$$
	| $$  | $$| $$__  $$| $$ /$$__  $$| $$__  $$| $$$$$   |  $$ /$$/
	| $$  | $$| $$  \ $$| $$| $$  \ $$| $$  \ $$| $$__/    \  $$$$/ 
	| $$  | $$| $$  | $$| $$| $$  | $$| $$  | $$| $$        >$$  $$ 
	|  $$$$$$/| $$  | $$| $$|  $$$$$$/| $$  | $$| $$$$$$$$ /$$/\  $$
	 \______/ |__/  |__/|__/ \______/ |__/  |__/|________/|__/  \__/
																	
																	
																	
	
	`)
	fmt.Println(`ρσωєяє∂ ву : нσℓу¢αη`)

	// add driver for support yaml content
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config.yml")
	if err != nil {
		panic(err)
	}

	core.ProcessBot(config.Default())
}
