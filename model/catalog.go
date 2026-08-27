package model

import "sort"

type Catalog struct{ Records []Record }

func (c *Catalog) Add(r Record) { c.Records = append(c.Records, r) }
func (c Catalog) Sorted() []Record {
	out := append([]Record(nil), c.Records...)
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.Before(out[j].UpdatedAt) })
	return out
}
func (c Catalog) CountByStatus() map[string]int {
	m := map[string]int{}
	for _, r := range c.Records {
		m[r.Status]++
	}
	return m
}
func (c Catalog) Routes() []string {
	seen := map[string]bool{}
	out := []string{}
	for _, r := range c.Records {
		k := r.Origin + "->" + r.Destination
		if !seen[k] {
			seen[k] = true
			out = append(out, k)
		}
	}
	return out
}
