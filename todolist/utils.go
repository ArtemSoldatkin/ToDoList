package todolist

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

func findDealByID(deals []*Deal, id string) *Deal {
	for _, d := range deals {
		if d.ID == id {
			return d
		}
	}
	return nil
}

func findDealIndex(deals []*Deal, id string) (index int) {
	for i, d := range deals {
		if d.ID == id {
			return i
		}
	}
	return -1
}

func checkValueByParam(deal *Deal, params map[string]string) bool {
	for k, v := range params {
		var result bool
		switch k {
		case "name":
			result = strings.Contains(deal.Name, v)
		case "description":
			result = strings.Contains(deal.Description, v)
		case "start_date":
			result = deal.StartDate.Format("02.01.2006") == v
		case "end_date":
			result = deal.EndDate.Format("02.01.2006") == v
		case "complete":
			b, err := strconv.ParseBool("true")
			if err != nil {
				result = true
			} else {
				result = deal.IsComplete == b
			}
		default:
			result = true
		}
		if !result {
			return false
		}
	}
	return true
}

func setValueByParam(deal *Deal, params map[string]string) {
	for k, v := range params {
		switch k {
		case "name":
			deal.Name = v
		case "description":
			deal.Description = v
		case "end_date":
			t, err := time.Parse("02.01.2006 15:4", v)
			if err != nil {
				continue
			}
			deal.EndDate = t
		default:
			continue
		}
	}
}

func sortByParam(deals []*Deal, sortBy, sortType string) (sortedDeals []Deal) {
	for _, d := range deals {
		sortedDeals = append(sortedDeals, *d)
	}
	switch sortBy {
	case "name":
		if sortType == "up" {
			sort.Sort(byName(sortedDeals))
		} else {
			sort.Reverse(byName(sortedDeals))
		}
	case "description":
		if sortType == "up" {
			sort.Sort(byDescription(sortedDeals))
		} else {
			sort.Reverse(byDescription(sortedDeals))
		}
	case "start_date":
		if sortType == "up" {
			sort.Sort(byStartDate(sortedDeals))
		} else {
			sort.Reverse(byStartDate(sortedDeals))
		}
	case "end_date":
		if sortType == "up" {
			sort.Sort(byEndDate(sortedDeals))
		} else {
			sort.Reverse(byEndDate(sortedDeals))
		}
	default:
		sort.Sort(byName(sortedDeals))
	}
	return
}
