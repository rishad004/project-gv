package repository

import "github.com/rishad004/project-gv/stream-service/internal/domain"

func (r *streamRepo) StreamCount() (int, error) {
	var streams []domain.Stream

	if err := r.Db.Find(&streams, "status=?", true).Error; err != nil {
		return 0, err
	}

	return len(streams), nil
}
