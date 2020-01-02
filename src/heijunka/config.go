package heijunka

// Config is structure for toml
type Config struct {
	Heijunka Heijunka `toml:"heijunka"`
	Solvers  []Solver `toml:"solver"`
}

// Heijunka section is the configurations for overall.
type Heijunka struct {
	ID   string `toml:"id"`
	Name int    `toml:"name"`
}

// Solver section is the configurations for the specific solver.
type Solver struct {
	ID   string `toml:"id"`
	Name string `toml:"name"`
}
