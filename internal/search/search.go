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
	states, err := GetStatesMap()
	if err != nil {
		return nil, err
	}
	var cities []string
	m := make(map[string][]City)
	for _, cityItem := range cityItems {
		city := strings.ToLower(cityItem.City)
		stateCode := strings.ToLower(cityItem.State)
		stateFull := strings.ToLower(states[cityItem.State])
		cityState := city + ", " + stateCode
		cityStateFull := city + ", " + stateFull
		if len(m[city]) == 0 {
			cities = append(cities, city)
		}
		cities = append(cities, cityState)
		cities = append(cities, cityStateFull)
		m[city] = append(m[city], cityItem)
		m[cityState] = append(m[cityState], cityItem)
		m[cityStateFull] = append(m[cityStateFull], cityItem)
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
		matchedCities := s.Map[match.Target]
		for _, match := range matchedCities {
			if cities.Contains(match) {
				continue
			}
			cities = append(cities, match)
		}
	}
	if len(cities) > 10 {
		cities = cities[:10]
	}
	return cities
}
