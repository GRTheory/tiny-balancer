package balancer

import "errors"

var (
	ErrorNoHost                = errors.New("no host")
	ErrorAlgorithmNotSupported = errors.New("algorithm not supported")
)

// Balancer iterface is the load balancer for the reverse proxy
type Balancer interface {
	Add(string)
	Remove(string)
	Balance(string) (string, error)
	Inc(string)
	Done(string)
}

//Factory is the factory that generates Balancer,
// and the factory design pattern is used here
type Factory func([]string) Balancer

var factories = make(map[string]Factory)

// Build generates teh corresponding Balancer according to the algorithm
func Build(algorithm string, hosts []string) (Balancer, error) {
	factory, ok := factories[algorithm]
	if !ok {
		return nil, ErrorAlgorithmNotSupported
	}
	return factory(hosts), nil
}