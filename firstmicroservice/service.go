package firstmicroservice

import (
	"context"
	"time"
)

//Service provides some "data capabilites" to your application
type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

type dataService struct{}

// NewService makes a new Service.
func NewService() Service {
	return dataService{}
}

func (dataService) Status(ctx context.Context) (string, error) {
	return "ok", nil

}

func (dataService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("01/01/2019"), nil
}

func (dataService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse("01/01/2019", date)
	if err != nil {
		return false, err
	}
	return true, nil
}
