package service

import "gincrud/pkg/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoServoce {
	return &videoService{}
}
