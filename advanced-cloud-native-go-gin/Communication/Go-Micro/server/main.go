package main

import (
	"fmt"

	proto "learning/advanced-cloud-native-go-gin/Communication/Go-Micro/proto"

	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// The Greeter API.
type Greeter struct{}

// Hello is a Greeter API method.
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name

	fmt.Printf("Responding with %s\n", rsp.Greeting)
	return nil
}

func main() {
	//  Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
