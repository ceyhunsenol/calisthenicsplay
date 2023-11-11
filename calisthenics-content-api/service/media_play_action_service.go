package service

import (
	"calisthenics-content-api/data/repository"
)

type IMediaPlayActionService interface {
	GetPlayAction()
}

type mediaPlayActionService struct {
}

func NewMediaPlayActionService(mediaRepo repository.IMediaRepository) IMediaPlayActionService {
	return &mediaPlayActionService{}
}

func (s *mediaPlayActionService) GetPlayAction() {
	
}
