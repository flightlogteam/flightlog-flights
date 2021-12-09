package start

type Writer interface {
	Create(start *Start) (uint, error)
	Update(start *Start) error
	Delete(id uint) error
}

type Repository interface {
	Writer
}

type Service interface {
	Writer
}
