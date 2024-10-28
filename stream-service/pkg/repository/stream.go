package repository

import (
	"github.com/rishad004/project-gv/stream-service/internal/domain"
	"gorm.io/gorm"
)

type streamRepo struct {
	Db *gorm.DB
}

func NewStreamRepo(Db *gorm.DB) StreamRepo {
	return &streamRepo{Db: Db}
}

// func (r *streamRepo) EndStream(id int32) error {
// 	var stream domain.Stream

// 	if err := r.Db.First(&stream, "streamer_id=? AND status=?", id, true).Error; err != nil {
// 		return err
// 	}

// 	stream.Status = false
// 	stream.StreamEnd = time.Now()

// 	if err := r.Db.Save(&stream).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *streamRepo) StartStream(id int32) error {
// 	if err := r.Db.Create(&domain.Stream{
// 		StreamerId:  int(id),
// 		Title:       "Live streaming....",
// 		Description: "Welcome to live stream....",
// 		StreamStart: time.Now(),
// 		Status:      true,
// 	}).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

func (r *streamRepo) StreamDetailing(streamerId int32, title, description string) error {
	var stream domain.Stream
	if err := r.Db.First(&stream, "streamer_id=?", streamerId).Error; err != nil {

		if er := r.Db.Create(&domain.Stream{
			StreamerId:  int(streamerId),
			Title:       title,
			Description: description,
		}).Error; er != nil {
			return er
		}

	} else {

		stream.Title = title
		stream.Description = description

		if err := r.Db.Save(&stream).Error; err != nil {
			return err
		}

	}

	return nil
}

func (r *streamRepo) StreamDetails(streamerId int32) (domain.Stream, error) {
	var stream domain.Stream

	if err := r.Db.First(&stream, "streamer_id=?", streamerId).Error; err != nil {
		return domain.Stream{}, err
	}

	return stream, nil
}
