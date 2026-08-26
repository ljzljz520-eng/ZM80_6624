package model

import "fmt"

func (r Record) Summary() string {
	return fmt.Sprintf("%s: %s (%s -> %s) [%s]", r.ID, r.Title, r.Origin, r.Destination, r.Status)
}
func (e Event) Label() string   { return e.Kind + " by " + e.Actor }
func (a Audit) IsFailure() bool { return a.Outcome != "success" }
