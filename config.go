package heijunka

// Config is structure for toml
type Config struct {
	Heijunka HeijunkaConfig `toml:"heijunka"`
	Solvers  []SolverConfig `toml:"solver"`
}

// HeijunkaConfig section is the configurations for overall.
type HeijunkaConfig struct {
	ID      string `toml:"id"`
	Name    string `toml:"name"`
	Verbose bool   `toml:"verbose"`
}

// SolverConfig section is the configurations for the specific solver.
type SolverConfig struct {
	ID   string `toml:"id"`
	Name string `toml:"name"`
}
