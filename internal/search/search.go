package search

import (
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Search struct {
	Cities []string
	Map    map[string][]City
}

func New() (*Search, error) {
	cityItems, err := GetCities()
	if err != nil {
		return nil, err
	}
	var cities []string
	m := make(map[string][]City)
	for _, cityItem := range cityItems {
		city := strings.ToLower(cityItem.City)
		if len(m[city]) == 0 {
			cities = append(cities, city)
		}
		m[city] = append(m[city], cityItem)
	}
	search := &Search{
		Cities: cities,
		Map:    m,
	}
	return search, nil
}

func (s *Search) GetTopTen(input string) Cities {
	cities := Cities{}
	if input == "" {
		return cities
	}
	input = strings.ToLower(input)
	matches := fuzzy.RankFind(input, s.Cities)
	sort.Sort(matches)
	for _, match := range matches {
		if len(cities) >= 10 {
			break
		}
		city := s.Map[match.Target]
		cities = append(cities, city...)
	}
	if len(cities) > 10 {
		cities = cities[:10]
	}
	return cities
}
