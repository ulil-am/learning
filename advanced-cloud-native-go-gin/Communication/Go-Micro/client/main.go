package main

import (
	"context"
	"fmt"
	"time"

	proto "learning/advanced-cloud-native-go-gin/Communication/Go-Micro/proto"

	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings
	service.Init()

	// Create new greeter client and call hello
	greeter := proto.NewGreeterClient("greeter", service.Client())
	callEvery(3*time.Second, greeter, hello)

}

func hello(t time.Time, greeter proto.GreeterClient) {
	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Ulil, calling at " + t.String()})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print response
	fmt.Printf("%s\n", rsp.Greeting)

}

func callEvery(d time.Duration, greeter proto.GreeterClient, f func(time.Time, proto.GreeterClient)) {
	for x := range time.Tick(d) {
		f(x, greeter)
	}
}
