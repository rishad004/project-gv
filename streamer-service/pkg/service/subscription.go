package service

func (s *streamerService) SubscriptionSetting(userId int, amount int) error {
	if err := s.repo.SubscriptionSetting(userId, amount); err != nil {
		return err
	}
	return nil
}

func (s *streamerService) SubscriptionCheck(userid int, id int) (int, int, error) {
	amount, sid, err := s.repo.SubscriptionCheck(userid, id)
	if err != nil {
		return 0, 0, err
	}
	return amount, sid, nil
}

func (s *streamerService) SubscriptionList() (map[int]map[string]any, error) {
	streamers, err := s.repo.StreamerList()
	if err != nil {
		return nil, err
	}

	res, er := s.repo.SubscriptionList(streamers)
	if er != nil {
		return nil, er
	}

	return res, nil
}
