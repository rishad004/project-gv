package repository

import "github.com/rishad004/project-gv/streamer-service/internal/domain"


func (r *streamerRepo) FindByStreamKey(channel string) (int32, string, error) {

	var streamer domain.Streamer
	if err := r.Db.First(&streamer, "name=?", channel).Error; err != nil {
		return 0, "", err
	}

	return int32(streamer.ID), streamer.StreamKey, nil
}

func (r *streamerRepo) StreamStart(key string) (int32, error) {

	var streamer domain.Streamer
	if err := r.Db.First(&streamer, "stream_key=?", key).Error; err != nil {
		return 0, err
	}

	return int32(streamer.ID), nil
}

func (r *streamerRepo) StreamEnd(key string) (int32, error) {

	var streamer domain.Streamer
	if err := r.Db.First(&streamer, "stream_key=?", key).Error; err != nil {
		return 0, err
	}

	return int32(streamer.ID), nil
}
