package model

import "errors"

type Status int

const (
	NOTPROCESS Status = iota
	ONPROCESS
	FINISH
	OTHERS
)

func (s Status) String() string {
	return [...]string{"Not Process", "On Process", "Finish"}[s]
}

func (s Status) EnumIndex() int {
	return int(s)
}

func StatusFromString(s string) (Status, error) {
	switch s {
	case "Not Process":
		return NOTPROCESS, nil
	case "On Process":
		return ONPROCESS, nil
	case "Finish":
		return FINISH, nil
	default:
		return OTHERS, errors.New("No status")
	}
}
