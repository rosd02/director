package pkg

import "time"

type shapeFactory struct {
	*Director
	components struct {
		*basicComponents
	}
}

func (factory *shapeFactory) update(s *Scene, dt time.Duration) {}