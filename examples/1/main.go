package main

import ni "github.com/dwdcth/netif"

func main() {
	is := ni.Parse(
		ni.Path("input"),
	)

	is.Write(
		ni.Path("output"),
	)
}
