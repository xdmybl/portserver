package main

var (
	GlobalBigJson = &BigJson{}
)

const (
	httpOk        = "200"
	httpNok       = "600"
	httpServerNok = "500"
)

var (
	TotalPoolRange = &PortRange{
		PortMin: 20000,
		PortMax: 29999,
	}
)
