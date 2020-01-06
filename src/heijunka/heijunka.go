package heijunka

import "fmt"

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

type Plan struct {
	Orders map[string]int `json:"map"`
}

type Vehicle struct {
	vc string
}

type Heijunka struct {
	Plan Plan `json:"plan"`
	Items []*Vehicle `json:"items"`
}

func NewHeijunka(orders map[string]int) *Heijunka {
	h := new(Heijunka)

	n := 0

	for k, v := range(orders) {
		h.Plan.Orders[k] = v

		n += v
	}

	h.Items = make([]*Vehicle, 0, n)

	return h
}

func (h *Heijunka) process() {
	categoryMap := make(map[string][]int)

	for vc, count := range(h.Plan.Orders) {
		cate := vc

		vcList, isOK := categoryMap[cate]

		if isOK {
			vcList = make([]int, 0, 10)
		}

		vcList = append(vcList, count)

		categoryMap[cate] = vcList
	}

	fmt.Println(categoryMap)
}