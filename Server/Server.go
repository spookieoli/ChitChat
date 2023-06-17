package Server

// Server go will create a simple webserver that will listen on port 8080 - it supports http and https

type Server struct {
	// Server will be a struct that will hold the port number and the router
	ip   string
	port string
	ssl  bool
}

// New will read the configuration JSON and returns a new Server struct
func New() *Server {
	// create a new Server struct
	s := Server{}
	// return the struct
	return &s
}
