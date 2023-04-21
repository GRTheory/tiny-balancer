package balancer

import "errors"

var (
	NoHostError                = errors.New("no host")
	AlgorithmNotSupportedError = errors.New("algorithm not supported")
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

var fatories = make(map[string]Factory)

// Build generates teh corresponding Balancer according to the algorithm
func Build(algorithm string, hosts []string) (Balancer, error) {
	factory, ok := fatories[algorithm]
	if !ok {
		return nil, AlgorithmNotSupportedError
	}
	return factory(hosts), nil
}