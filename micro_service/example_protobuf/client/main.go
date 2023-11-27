package main

import "google.golang.org/grpc"

func main() {
	grpc.Dial("127.0.0.1:8089", grpc.WithInsecure())
}
