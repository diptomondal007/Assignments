package worker_pool

type Request struct {
	Data interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})