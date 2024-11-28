package repository

import (
	"fmt"
	"time"

	"github.com/rishad004/project-gv/stream-service/internal/domain"
)

func (r *streamRepo) EndStream(id int32) error {
	var stream domain.Stream

	if err := r.Db.First(&stream, "streamer_id=? AND status=?", id, true).Error; err != nil {
		return nil
	}

	stream.Status = false
	stream.StreamEnd = time.Now()

	r.Db.Save(&stream)

	return nil
}

func (r *streamRepo) StartStream(id int32) error {
	if err := r.Db.Create(&domain.Stream{
		StreamerId:  int(id),
		StreamStart: time.Now(),
		Status:      true,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *streamRepo) StreamDetailing(streamerId int32, title, description string) error {

	var stream domain.StreamData

	if err := r.Db.First(&stream, "streamer_id=?", streamerId).Error; err != nil {

		if er := r.Db.Create(&domain.StreamData{
			StreamerId:  int(streamerId),
			Title:       title,
			Description: description,
		}).Error; er != nil {
			return er
		}
		return nil

	}

	fmt.Println(stream)

	stream.Title = title
	stream.Description = description

	if err := r.Db.Save(&stream).Error; err != nil {
		return err
	}

	return nil
}

func (r *streamRepo) StreamDetails(streamerId int32) (domain.StreamData, error) {
	var stream domain.StreamData

	if err := r.Db.First(&stream, "streamer_id=?", streamerId).Error; err != nil {
		return domain.StreamData{}, err
	}

	return stream, nil
}
