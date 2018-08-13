package grpc

type Request struct {
	Data     string
	Response interface{}
}

type server struct {
	apiDBRequest chan Request
}
