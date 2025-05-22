package main

import (
	"flag"

	"github.com/san-lab/sx402/facilitator"
)

func main() {
	withDemoStore := flag.Bool("demoStore", false, "starts the demo store under /store")
	flag.Parse()
	//*withDemoStore = true
	facilitator.Start(*withDemoStore)

}
