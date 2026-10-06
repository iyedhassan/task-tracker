
package task

type Repository interface {
	Load() ([]Task, error)
	Save(tasks []Task) error
}