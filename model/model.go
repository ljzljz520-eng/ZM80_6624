package model

import "time"

type Record struct {
	ID, Title, Origin, Destination, Status, Rule, Notes string
	CreatedAt, UpdatedAt                                time.Time
}
type Profile struct {
	ID, Institution, Contact, Email string
	Approved                        bool
}
type Event struct {
	ID, RecordID, Kind, Actor, Detail string
	At                                time.Time
}
type Audit struct {
	ID, RecordID, Action, Outcome, Reason string
	At                                    time.Time
}

func NewRecord(id, title, origin, destination string) Record {
	now := time.Now().UTC()
	return Record{ID: id, Title: title, Origin: origin, Destination: destination, Status: "received", CreatedAt: now, UpdatedAt: now}
}
func (r Record) Validate() error {
	if r.ID == "" {
		return ErrInvalid("record id")
	}
	if r.Title == "" {
		return ErrInvalid("title")
	}
	if r.Origin == "" || r.Destination == "" {
		return ErrInvalid("route")
	}
	return nil
}
func (r *Record) Transition(status string) error {
	switch status {
	case "received", "reviewed", "approved", "archived", "rejected":
		r.Status = status
		r.UpdatedAt = time.Now().UTC()
		return nil
	default:
		return ErrInvalid("status")
	}
}

type invalidError string

func (e invalidError) Error() string { return "invalid " + string(e) }
func ErrInvalid(s string) error      { return invalidError(s) }
