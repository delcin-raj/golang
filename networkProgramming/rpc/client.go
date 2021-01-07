package main

import (
	"net/rpc"
	"fmt"
	"./types"
	"check"
)

func main() {
	client,err := rpc.DialHTTP("tcp","localhost:1234")
	check.Err(err)

	// the results for the three calls are stored here
	var mul int
	var div types.Quotient
	var hello string

	check.Err(
		client.Call(
			"WebService.Multiply",
			types.Ints{3,5},
			&mul,
		),
	)
	check.Err(
		client.Call(
			"WebService.Divide",
			types.Ints{99,43},
			&div,
		),
	)
	check.Err(
		client.Call(
			"WebService.SayHello",
			types.Name{"Delcin","Raj"},
			&hello,
		),
	)

	fmt.Println("Mul: ",mul,", Div: ",div,", hello: ",hello)
}
