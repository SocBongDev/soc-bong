package database

type Migrator interface {
	Up() error
	Down(steps int) error
}
