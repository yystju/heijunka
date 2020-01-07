package heijunka

import (
	"log"
	"math"
	"sort"
)

// Plan is ...
type Plan struct {
	Orders map[string]int `json:"orders"`
	Categories map[string][]string `json:"categories"`
}

// Heijunka is ...
type Heijunka struct {
	Config *Config `json:"config"`
	Plan  Plan     `json:"plan"`
	Items []string `json:"items"`
}

// NewHeijunka is ...
func NewHeijunka(config *Config, orders map[string]int, categories map[string][]string) *Heijunka {
	h := new(Heijunka)

	h.Config = config

	n := 0

	h.Plan.Orders = make(map[string]int)

	for k, v := range orders {
		h.Plan.Orders[k] = v

		n += v
	}

	h.Items = make([]string, 0, n)

	h.Plan.Categories = make(map[string][]string)

	for k, v := range categories {
		h.Plan.Categories[k] = v
	}

	return h
}

// Process is ...
func (h *Heijunka) Process() {
	n := h.count()

	proportion := make(map[string]float64)

	for k, v := range h.Plan.Categories {
		c := 0

		for _, o := range v {
			c += (h.Plan.Orders[o])
		}

		proportion[k] = float64(c) / float64(n)
	}

	if h.Config.Heijunka.Verbose {
		log.Printf("proportion : %v", proportion)
	}

	for i := 0; i < n; i++ {
		if h.Config.Heijunka.Verbose {
			log.Printf("[ROUND %v]", i)
		}

		ratios := make(map[string]float64)

		for k, v := range h.Plan.Categories {
			existed := 0

			for _, o := range v {
				existed += h.existed(o)
			}

			ratios[k] = float64(i+1)*proportion[k] - float64(existed)
		}

		if h.Config.Heijunka.Verbose {
			log.Printf("ratios : %v", ratios)
		}

		values := make([]float64, 0, len(h.Plan.Orders))

		for k := range h.Plan.Categories {
			r := arrayIndexOf(values, ratios[k])

			if r == -1 {
				values = append(values, ratios[k])

				r = arrayIndexOf(values, ratios[k])
			}
		}

		sort.Sort(sort.Reverse(sort.Float64Slice(values)))

		if h.Config.Heijunka.Verbose {
			log.Printf("values : %v", values)
		}

		rank := make(map[int][]string)

		for k := range h.Plan.Categories {
			r := arrayIndexOf(values, ratios[k])

			arry, ok := rank[r]

			if !ok {
				arry = make([]string, 0)
			}

			arry = append(arry, k)

			rank[r] = arry
		}

		if h.Config.Heijunka.Verbose {
			log.Printf("rank : %v", rank)
		}

		keys := mapKeysi(rank)

		if h.Config.Heijunka.Verbose {
			log.Printf("keys : %v", keys)
		}

		for _, r := range keys {
			cates := rank[r]

			if h.Config.Heijunka.Verbose {
				log.Printf("cates : %v", cates)
			}

			selected := make([]string, 0)

			for _, o := range cates {
				for _, e := range h.Plan.Categories[o] {
					selected = append(selected, e)
				}
			}


			if h.Config.Heijunka.Verbose {
				log.Printf("selected : %v", selected)
			}

			var s string

			min := len(h.Items)

			m := make(map[string]int)

			for _, o := range selected {
				n := h.existed(o)

				m[o] = n

				if n < min {
					min = n
				}
			}

			selected = make([]string, 0, len(m))

			for k,v := range m {
				if min == v {
					selected = append(selected, k)
				}
			}
			
			if len(selected) > 1 {
				var last string

				if len(h.Items) > 0 {
					last = h.Items[len(h.Items) - 1]
				} else {
					last = ""
				}

				for _, o := range selected {
					if o != last {
						s = o
						break
					}
				}
			} else if len(selected) == 1 {
				s = selected[0]
			}

			if h.Config.Heijunka.Verbose {
				log.Printf("final selected : %v", s)
			}

			h.Items = append(h.Items, s)

			break
		}

		if h.Config.Heijunka.Verbose {
			log.Printf("items : %v", h.Items)
		} 
	}
}

func (h *Heijunka) count() int {
	n := 0

	for _, v := range h.Plan.Orders {
		n += v
	}

	return n
}

func (h *Heijunka) existed(item string) int {
	n := 0

	for i := range h.Items {
		if item == h.Items[i] {
			n++
		}
	}

	return n
}

func arrayIndexOf(arry []float64, item float64) int {
	index := -1

	for i := range arry {
		if math.Abs(arry[i]-item) < 1e-5 {
			index = i
			break
		}
	}

	return index
}

func mapKeysi(m map[int][]string) []int {
	keys := make([]int, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Sort(sort.IntSlice(keys))

	return keys
}

func mapKeyss(m map[string]int) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Sort(sort.StringSlice(keys))

	return keys
}
