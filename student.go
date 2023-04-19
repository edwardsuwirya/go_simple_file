package main

import "fmt"

type Student struct {
	Id      string
	Name    string
	Address string
}

func (s *Student) ToCsv(delimiter string) string {
	return fmt.Sprintln("%s%s%s%s%s", s.Id, delimiter, s.Name, delimiter, s.Address)
}
