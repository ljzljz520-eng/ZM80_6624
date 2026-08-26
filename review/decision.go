package review

import "museumlogistics/model"

type Decision struct {
	Approved     bool
	Rule, Reason string
}

func Decide(r model.Record) Decision {
	if err := Evaluate(r); err != nil {
		return Decision{Rule: PolicyName(r.Origin), Reason: err.Error()}
	}
	return Decision{Approved: true, Rule: PolicyName(r.Origin), Reason: "all rules passed"}
}
func Severity(d Decision) string {
	if d.Approved {
		return "low"
	}
	if d.Rule == "export-control" {
		return "critical"
	}
	return "major"
}
func RequiresCurator(d Decision) bool { return !d.Approved && Severity(d) == "critical" }
