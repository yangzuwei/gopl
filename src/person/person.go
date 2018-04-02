// person.go
package person

type Man struct {
	Name     string
	Age      int
	HasBeard bool
}

func (m *Man) Run() string {
	return m.Name + " is running!"
}
