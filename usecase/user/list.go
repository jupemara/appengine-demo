package user

import (
	"github.com/jupemara/appengine-demo/domain/model/user"
	"golang.org/x/sync/errgroup"
)

type ListUserUsecase struct {
	repositories []user.UserRepository
}

type Dto struct {
	Repository string
	Users      []*user.User
}

func NewListUserUsecase(repositories []user.UserRepository) *ListUserUsecase {
	return &ListUserUsecase{repositories}
}

func (u *ListUserUsecase) Execute() ([]*Dto, error) {
	rc := make(chan *Dto, len(u.repositories))
	eg := new(errgroup.Group)
	for _, r := range u.repositories {
		r := r
		eg.Go(func() error {
			users, err := r.List()
			if err != nil {
				return err
			}
			rc <- &Dto{
				Repository: r.VisibleName(),
				Users:      users,
			}
			return nil
		})
	}
	err := eg.Wait()
	close(rc)
	if err != nil {
		return nil, err
	}
	results := make([]*Dto, 0, len(u.repositories))
	for v := range rc {
		results = append(results, v)
	}
	return results, nil
}
