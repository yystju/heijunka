package heijunka

import (
	"log"
	"math"
	"sort"
)

// type Item interface {
// 	Type() string
// 	Name() string
// 	Gategory() *Category
// }

// type Category interface {
// 	Name() string
// // }

// type Category struct {
// 	name string
// }

// func (c *Category) Name() string {
// 	return c.name;
// }

// func NewCategory(name string) *Category {
// 	c := new(Category)

// 	c.name = name

// 	return c
// }

// type VC struct {
// 	name string
// 	count int
// 	category *Category
// 	// Item;
// }

// func (vc *VC) Type() string {
// 	return "VC"
// }

// func (vc *VC) Name() string {
// 	return vc.name
// }

// func (vc *VC) Gategory() *Category {
// 	return vc.category
// }

// func NewVC(name string, category * Category) *VC {
// 	vc := new(VC)

// 	vc.name = name
// 	vc.category = category

// 	return vc
// }

// Plan is ...
type Plan struct {
	Orders map[string]int `json:"map"`
}

// Heijunka is ...
type Heijunka struct {
	Plan  Plan     `json:"plan"`
	Items []string `json:"items"`
}

// NewHeijunka is ...
func NewHeijunka(orders map[string]int) *Heijunka {
	h := new(Heijunka)

	n := 0

	h.Plan.Orders = make(map[string]int)

	for k, v := range orders {
		h.Plan.Orders[k] = v

		n += v
	}

	h.Items = make([]string, 0, n)

	return h
}

// Process is ...
func (h *Heijunka) Process() {
	n := h.count()

	// showup := make(map[string]float64)
	proportion := make(map[string]float64)

	for k, v := range h.Plan.Orders {
		proportion[k] = float64(v) / float64(n)
	}

	for i := 0; i < n; i++ {
		ratios := make(map[string]float64)

		for k := range h.Plan.Orders {
			ratios[k] = float64(i+1)*proportion[k] - float64(h.existed(k))
		}

		values := make([]float64, 0, len(h.Plan.Orders))

		for k := range h.Plan.Orders {
			r := arrayIndexOf(values, ratios[k])

			if r == -1 {
				values = append(values, ratios[k])

				r = arrayIndexOf(values, ratios[k])
			}
		}

		sort.Sort(sort.Reverse(sort.Float64Slice(values)))

		// log.Printf("values : %v", values)

		rank := make(map[int][]string)

		for k := range h.Plan.Orders {
			r := arrayIndexOf(values, ratios[k])

			arry, ok := rank[r]

			if !ok {
				arry = make([]string, 0)
			}

			arry = append(arry, k)

			rank[r] = arry
		}

		log.Printf("i : %v, ratios : %v", i, ratios)

		log.Printf("rank : %v", rank)

		log.Printf("items : %v", h.Items)

		keys := mapKeysi(rank)

		for _, r := range keys {
			selected := rank[r]

			var s string

			if len(selected) > 1 {
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
			} else {
				s = selected[0]
			}

			log.Printf("selected : %v", s)

			h.Items = append(h.Items, s)

			break
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
