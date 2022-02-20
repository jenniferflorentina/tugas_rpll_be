package model

import "errors"

type Status int

const (
	NotProcess Status = iota
	OnProcess
	Finish
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
		return NotProcess, nil
	case "On Process":
		return On, nil
	case "Finish":
		return Finish, nil
	default:
		return errors.New("No Status")
	}
}
