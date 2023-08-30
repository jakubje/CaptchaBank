package captcha

import "sync"

type LocalStore struct {
	mu          sync.Mutex
	captchaList map[string][]string
}

func NewLocalStore() *LocalStore {
	return &LocalStore{
		captchaList: make(map[string][]string),
	}
}

func (s *LocalStore) AddCaptcha(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	slice := s.captchaList[key]
	slice = append(slice, value)
	s.captchaList[key] = slice
}

func (s *LocalStore) GetCaptcha(key string) ([]string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	slice, ok := s.captchaList[key]
	if !ok || len(slice) == 0 {
		return nil, false
	}
	item := slice[0]
	slice = slice[1:]
	s.captchaList[key] = slice
	return []string{item}, true
}
