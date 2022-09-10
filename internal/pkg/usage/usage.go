package usage

import "fmt"

// app usage menual
func Usage() {
	fmt.Println(`Usage: go-web [-i <alias>] [-n <title>] -[gcsdhv]

	<descriptive text>

		 ------- < Commands Arguments > -------
	argument:
	 -c, config        Create a new blog. (e.g go-web -c <external config file>)

	optional:
	  -h, help          Show this help message. 
	  -v, version       Show this app version. 

For more help, you can use 'go-web help' for the detailed information
or you can check the docs: https://github.com/ka1i/go-web`)
}
