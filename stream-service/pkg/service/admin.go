package service

func (s *streamSvc) StreamCount() (int, error) {
	count, err := s.repo.StreamCount()

	if err != nil {
		return 0, err
	}

	return count, nil
}
