package models

import "fmt"

type Category int

const (
	Mammal Category = iota
	Bird
	Reptile
	Fish
	Amphibian
)

var CategoryNames = []string{
	"Mammal",
	"Bird",
	"Reptile",
	"Fish",
	"Amphibian",
}

func (c Category) String() string {
	if int(c) < len(CategoryNames) {
		return CategoryNames[c]
	}
	return "Unknown"
}

type Animal interface {
	Speak() string
	Move() string
	GetID() int
	GetName() string
	GetCategory() Category
	GetDetails() string
	SetID(int)
}

type Lion struct {
	ID	   int	   `json:"id"`
	Name   string  `json:"name"`
	Age	   int     `json:"age"`
	Category Category `json:"category"`
}

func (l *Lion) Speak() string {
	return "Roar"
}

func (l *Lion) Move() string {
	return "Run"
}

func (l *Lion) GetID() int {
	return l.ID
}

func (l *Lion) GetName() string {
	return l.Name
}

func (l *Lion) GetCategory() Category {
	return l.Category
}

func (l *Lion) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d, Category: %s", l.ID, l.Name, l.Age, l.Category)
}

func (l *Lion) SetID(id int) {
	l.ID = id
}

type Eagle struct {
	ID	   int	   `json:"id"`
	Name   string  `json:"name"`
	Age	   int     `json:"age"`
	Category Category `json:"category"`
}

func (e *Eagle) Speak() string {
	return "Screech"
}

func (e *Eagle) Move() string {
	return "Fly"
}

func (e *Eagle) GetID() int {
	return e.ID
}

func (e *Eagle) GetName() string {
	return e.Name
}

func (e *Eagle) GetCategory() Category {
	return e.Category
}

func (e *Eagle) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d, Category: %s", e.ID, e.Name, e.Age, e.Category)
}

func (e *Eagle) SetID(id int) {
	e.ID = id
}

type Snake struct {
	ID	   int	   `json:"id"`
	Name   string  `json:"name"`
	Age	   int     `json:"age"`
	Category Category `json:"category"`
}

func (s *Snake) Speak() string {
	return "Hiss"
}

func (s *Snake) Move() string {
	return "Slither"
}

func (s *Snake) GetID() int {
	return s.ID
}

func (s *Snake) GetName() string {
	return s.Name
}

func (s *Snake) GetCategory() Category {
	return s.Category
}

func (s *Snake) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d, Category: %s", s.ID, s.Name, s.Age, s.Category)
}

func (s *Snake) SetID(id int) {
	s.ID = id
}

type Shark struct {
	ID	   int	   `json:"id"`
	Name   string  `json:"name"`
	Age	   int     `json:"age"`
	Category Category `json:"category"`
}

func (s *Shark) Speak() string {
	return "Silent"
}

func (s *Shark) Move() string {
	return "Swim"
}

func (s *Shark) GetID() int {
	return s.ID
}

func (s *Shark) GetName() string {
	return s.Name
}

func (s *Shark) GetCategory() Category {
	return s.Category
}

func (s *Shark) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d, Category: %s", s.ID, s.Name, s.Age, s.Category)
}

func (s *Shark) SetID(id int) {
	s.ID = id
}

type Frog struct {
	ID	   int	   `json:"id"`
	Name   string  `json:"name"`
	Age	   int     `json:"age"`
	Category Category `json:"category"`
}

func (f *Frog) Speak() string {
	return "Croak"
}

func (f *Frog) Move() string {
	return "Jump"
}

func (f *Frog) GetID() int {
	return f.ID
}

func (f *Frog) GetName() string {
	return f.Name
}

func (f *Frog) GetCategory() Category {
	return f.Category
}

func (f *Frog) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d, Category: %s", f.ID, f.Name, f.Age, f.Category)
}

func (f *Frog) SetID(id int) {
	f.ID = id
}