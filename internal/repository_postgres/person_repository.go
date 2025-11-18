package repositorypostgres

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/DmitriyChubarov/enkod/internal/app"
	"github.com/gocraft/dbr/v2"
)

type PersonRepository interface {
	Create(ctx context.Context,p *app.Person) error
	GetByID(ctx context.Context, id int64) (*app.Person, error)
	Update(ctx context.Context, id int64, p *app.Person) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int, search string) ([]*app.Person, error)
}

type personRepository struct {
	session *dbr.Session
}

func NewPersonRepository(session *dbr.Session) PersonRepository {
	return &personRepository{session: session}
}

func (r *personRepository) Create(ctx context.Context, p *app.Person) error {
	log.WithFields(log.Fields{
		"email":     p.Email,
		"firstName": p.FirstName,
		"lastName":  p.LastName,
	}).Debug("Выполнение SQL: INSERT INTO person")

	_, err := r.session.InsertInto("person").
		Columns("email", "phone", "first_name", "last_name").
		Record(p).
		ExecContext(ctx)
	
	if err != nil {
		log.WithFields(log.Fields{
			"email": p.Email,
			"error": err.Error(),
		}).Error("SQL ошибка при создании пользователя")
		return err
	}
	
	log.WithField("email", p.Email).Debug("SQL: пользователь успешно создан в базе данных")
	return err
}

func (r *personRepository) GetByID(ctx context.Context, id int64) (*app.Person, error) {
	log.WithField("id", id).Debug("Выполнение SQL: SELECT FROM person WHERE id = ?")
	
	var p app.Person
	err := r.session.Select("*").From("person").Where("id = ?", id).LoadOneContext(ctx, &p)
	if err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err.Error(),
		}).Debug("SQL: пользователь не найден")
		return nil, errors.New("person not found")
	}
	
	log.WithFields(log.Fields{
		"id":    p.Id,
		"email": p.Email,
	}).Debug("SQL: пользователь успешно получен из базы данных")
	return &p, nil
}

func (r *personRepository) Update(ctx context.Context, id int64, p *app.Person) error {
	log.WithFields(log.Fields{
		"id":        id,
		"email":     p.Email,
		"firstName": p.FirstName,
		"lastName":  p.LastName,
	}).Debug("Выполнение SQL: UPDATE person WHERE id = ?")

	_, err := r.session.Update("person").
		SetMap(map[string]interface{}{
			"email":      p.Email,
			"phone":      p.Phone,
			"first_name": p.FirstName,
			"last_name":  p.LastName,
		}).
		Where("id = ?", id).
		ExecContext(ctx)
	
	if err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("SQL ошибка при обновлении пользователя")
		return err
	}
	
	log.WithField("id", id).Debug("SQL: пользователь успешно обновлен в базе данных")
	return err
}

func (r *personRepository) Delete(ctx context.Context, id int64) error {
	log.WithField("id", id).Debug("Выполнение SQL: DELETE FROM person WHERE id = ?")
	
	_, err := r.session.DeleteFrom("person").Where("id = ?", id).ExecContext(ctx)
	
	if err != nil {
		log.WithFields(log.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("SQL ошибка при удалении пользователя")
		return err
	}
	
	log.WithField("id", id).Debug("SQL: пользователь успешно удален из базы данных")
	return err
}

func (r *personRepository) List(ctx context.Context, limit, offset int, search string) ([]*app.Person, error) {
	log.WithFields(log.Fields{
		"limit":  limit,
		"offset": offset,
		"search": search,
	}).Debug("Выполнение SQL: SELECT FROM person с фильтрами")

	var people []*app.Person
	q := r.session.Select("*").From("person").Limit(uint64(limit)).Offset(uint64(offset))
	if search != "" {
		q = q.Where("email ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	_, err := q.LoadContext(ctx, &people)
	
	if err != nil {
		log.WithFields(log.Fields{
			"limit":  limit,
			"offset": offset,
			"error":  err.Error(),
		}).Error("SQL ошибка при получении списка пользователей")
		return people, err
	}
	
	log.WithField("count", len(people)).Debug("SQL: список пользователей успешно получен из базы данных")
	return people, err
}
