package main
// Server defines the functions and listens for requests 

import (
	"fmt"
	"net/rpc"
	"errors"
	"net/http"
	"./types"
)

// Let us assign a type/name to our rpc
type WebService struct {}

// NOTE: This allows us to define various functionalities for a single server

func(t *WebService) Multiply(args *types.Ints, res *int) error {
	*res = args.A * args.B
	return nil
}

func(t* WebService) Divide(args *types.Ints,quo *types.Quotient) error {
	if args.B == 0 {
		return errors.New("division by 0")
	}
	quo.Q = args.A / args.B
	quo.R = args.A % args.B
	return nil
}

func(t* WebService) SayHello(args *types.Name,res *string) error {
	*res = args.Fname + " " + args.Lname + " says hello!"
	return nil
}

func main() {
	server := new(WebService)
	// server.Multiply, WebService.Multiply()
	rpc.Register(server)
	// Server Procedure Stubs
	// m := map("WebService" -> server)
	// client is making a request for "WebService.Multiply" --> "WebService" "Multiply"
	// m["WebService"].Multiply(args,res) 
	rpc.HandleHTTP()

	fmt.Println("Service has started, waiting for request ....")
	err := http.ListenAndServe(":1234",nil)
	if err != nil {
		fmt.Println(err.Error)
	}
	// if client makes Multiply request the rpc package will call server.Multiply()
}




