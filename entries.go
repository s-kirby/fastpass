package fastpass

import (
	"math"
	"sort"

	"strings"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

//Entries is a set of entries
type Entries []*Entry

//FuzzyMatch returns all entries with names fuzzy matching search
func (es Entries) FuzzyMatch(search string) (ret Entries) {
	for _, entry := range es {
		if fuzzy.Match(search, entry.Name) {
			ret = append(ret, entry)
		}
	}
	return
}

//DeleteByName deletes an entry from es and returns the new slice
func (es Entries) DeleteByName(name string) (cleaned Entries) {
	for i, e := range es {
		if e.Name == name {
			cleaned = append(cleaned, es[:i]...)
			if i < (len(es) - 1) {
				cleaned = append(cleaned, es[i+1:]...)
			}
			return
		}
	}
	return es
}

//FindByName finds an entry by it's name.
//It returns nil if no entry was found.
func (es Entries) FindByName(name string) *Entry {
	for _, e := range es {
		if e.Name == name {
			return e
		}
	}
	return nil
}

//SortByName sorts es by name
func (es Entries) SortByName() Entries {
	sort.Slice(es, func(i, j int) bool {
		return es[i].Name < es[j].Name
	})
	return es
}

//SortByHits sorts es by hits
func (es Entries) SortByHits() Entries {
	sort.Slice(es, func(i, j int) bool {
		return es[i].Stats.Activity > es[j].Stats.Activity
	})
	return es
}

//SortByBestMatch tries to sort entries by best match
func (es Entries) SortByBestMatch(search string) Entries {
	distances := make([]int, len(es))
	for i, e := range es {
		distances[i] = fuzzy.LevenshteinDistance(search, e.Name)
	}
	sort.Slice(es, func(i, j int) bool {
		//if i contains the substring but j doesn't i is much more likely to be correct.
		//and vice versa
		{
			inI := strings.Contains(es[i].Name, search)
			inJ := strings.Contains(es[j].Name, search)
			if inI && !inJ {
				return true
			}
			if inJ && !inI {
				return false
			}
		}

		// //if j has everything in common, j's better
		// if jDistance == 0 {
		// 	return false
		// }

		//i is no where close to being close
		if distances[i] < 0 {
			// fmt.Printf("%v:%v dist < 0\n", es[i].Name, distances[i])
			return false
		}

		//j is no where close to being close
		if distances[j] < 0 {
			// fmt.Printf("%v:%v dist < 0\n", es[j].Name, distances[j])
			return true
		}

		score := func(entry *Entry) float64 {
			//as name becomes closer to query, this goes up
			return (math.Log2(float64(entry.Stats.Activity)+1) / (float64(distances[i])))
		}
		// fmt.Printf("scoring %v:%v, %v:%v...\n", es[i].Name, score(es[i]), es[j].Name, score(es[j]))
		return score(es[i]) > score(es[j])
	})
	return es
}

// //FilterByTag returns all entries with a certain tag
// func (entries Entries) FilterByTag(tag string) Entries {
// 	var matches Entries
// 	for _, e := range entries {
// 		if e.HasTag(tag) {
// 			matches = append(matches, e)
// 		}
// 	}
// 	return matches
// }
