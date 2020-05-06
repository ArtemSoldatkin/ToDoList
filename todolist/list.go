package todolist

import "strings"

// ToDoList - list of deals
type ToDoList struct {
	Deals []*Deal
}

func (t ToDoList) String() string {
	var result []string
	for _, d := range t.Deals {
		result = append(result, d.String())
	}
	return strings.Join(result, "\n")
}

// Add - add deal
func (t *ToDoList) Add(deal *Deal) {
	t.Deals = append(t.Deals, deal)
}

// Edit - edit deal
func (t *ToDoList) Edit(id string, params map[string]string) {
	deal := findDealByID(t.Deals, id)
	if deal == nil {
		return
	}
	setValueByParam(deal, params)
}

// Complete - complete deal
func (t *ToDoList) Complete(id string) {
	deal := findDealByID(t.Deals, id)
	if deal == nil {
		return
	}
	deal.Complete()
}

// Remove - remove deal from deal list
func (t *ToDoList) Remove(id string) {
	idx := findDealIndex(t.Deals, id)
	if idx == -1 {
		panic("Deal is not found")
	}
	t.Deals[idx] = t.Deals[len(t.Deals)-1]
	t.Deals[len(t.Deals)-1] = nil
	t.Deals = t.Deals[:len(t.Deals)-1]
}

// Find - find deal by id
func (t ToDoList) Find(id string) *Deal {
	return findDealByID(t.Deals, id)
}

// Filter - filter deals by parameters
func (t ToDoList) Filter(params map[string]string) (deals []Deal) {
	for _, d := range t.Deals {
		if checkValueByParam(d, params) {
			deals = append(deals, *d)
		}
	}
	return
}

// Sort - sort deals list
func (t *ToDoList) Sort(sortBy, sortType string) []Deal {
	return sortByParam(t.Deals, sortBy, sortType)
}
