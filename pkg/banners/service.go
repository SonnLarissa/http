package banners

import (
	"context"
	"errors"
	"sync"
)

type Service struct {
	mu    sync.RWMutex
	items []*Banner
}

func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}
}

type Banner struct {
	ID      int64
	Title   string
	Content string
	Button  string
	Link    string
}

func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items, nil
}

func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if item.ID == 0 {
		ind := int64(len(s.items) + 1)
		item.ID = ind
		s.items = append(s.items, item)
	} else {
		if ind, _ := s.getBannerById(item.ID); ind != -1 {
			s.items[ind] = item
		}
	}
	return item, nil
}

func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ind, banner := s.getBannerById(id)
	if ind == -1 {
		return nil, errors.New("banner not found")
	}
	s.items = remove(s.items, ind)
	return banner, nil
}

func (s *Service) getBannerById(id int64) (index int, banner *Banner) {
	index = -1
	for i, b := range s.items {
		if b.ID == id {
			index = i
			banner = b
			break
		}
	}
	return index, banner
}

func remove(items []*Banner, ind int) []*Banner {
	return append(items[:ind], items[ind+1:]...)
}
