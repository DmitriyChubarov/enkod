package logic

import (
	"context"
	"errors"
	"regexp"

	log "github.com/sirupsen/logrus"
	"github.com/DmitriyChubarov/enkod/internal/app"
	repositorypostgres "github.com/DmitriyChubarov/enkod/internal/repository_postgres"
)

type PersonService struct {
	repo repositorypostgres.PersonRepository
}

func NewPersonService(repo repositorypostgres.PersonRepository) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) CreatePerson(ctx context.Context, p *app.Person) (*app.Person, error) {
	if p.Email == "" {
		log.Error("Email не может быть пустым")
		return nil, errors.New("email is required")
	}

	matched, _ := regexp.MatchString(`.+@.+\..+`, p.Email)
	if !matched {
		log.WithField("email", p.Email).Error("Email введен некорректно")
		return nil, errors.New("invalid email format")
	}
	if err := s.repo.Create(p); err != nil {
		log.WithFields(log.Fields{
			"email": p.Email,
			"error": err.Error(),
		}).Error("Ошибка при создании пользователя в базе данных")
		return nil, err
	}
	log.WithFields(log.Fields{
		"id":        p.Id,
		"email":     p.Email,
		"firstName": p.FirstName,
		"lastName":  p.LastName,
	}).Info("Пользователь успешно создан")
	return p, nil
}

func (s *PersonService) GetPerson(ctx context.Context, id int64) (*app.Person, error) {
	person, err := s.repo.GetByID(id)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("Ошибка при получении пользователя")
		return nil, err
	}
	log.WithFields(log.Fields{
		"id":    person.Id,
		"email": person.Email,
	}).Info("Пользователь успешно получен")
	return person, nil
}

func (s *PersonService) UpdatePerson(ctx context.Context, id int64, p *app.Person) (*app.Person, error) {
	if err := s.repo.Update(id, p); err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("Ошибка при обновлении пользователя в базе данных")
		return nil, err
	}
	log.WithField("id", id).Info("Пользователь успешно обновлен")
	return s.repo.GetByID(id)
}

func (s *PersonService) DeletePerson(ctx context.Context, id int64) error {
	if err := s.repo.Delete(id); err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("Ошибка при удалении пользователя из базы данных")
		return err
	}
	log.WithField("id", id).Info("Пользователь успешно удален")
	return nil
}

func (s *PersonService) ListPersons(ctx context.Context, limit, offset int, search string) ([]*app.Person, error) {
	people, err := s.repo.List(limit, offset, search)
	if err != nil {
		log.WithFields(log.Fields{
			"limit":  limit,
			"offset": offset,
			"error":  err.Error(),
		}).Error("Ошибка при получении списка пользователей из базы данных")
		return nil, err
	}
	log.WithField("count", len(people)).Info("Список пользователей успешно получен")
	return people, nil
}
