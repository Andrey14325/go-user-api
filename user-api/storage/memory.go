package storage

import (
	"errors"
	"maps"
	"user-api/models"

	"github.com/google/uuid"
)

type MemoryStorage struct {
	users map[uuid.UUID]models.User
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users: make(map[uuid.UUID]models.User),
	}
}

func (m *MemoryStorage) CreateUser(name string, email string) error {
	if name == "" {
		return errors.New("empty name")
	}
	if email == "" {
		return errors.New("empty email")
	}
	id := uuid.New()
	m.users[id] = *models.MakeUser(id, name, email)
	return nil
}

func (m MemoryStorage) GetUsers() []models.User {
	allUser := make([]models.User, 0)
	for _, val := range m.users {
		allUser = append(allUser, val)
	}
	return allUser
}

func (m MemoryStorage) GetUser(id uuid.UUID) (models.User, error) {
	val, exis := m.users[id]
	if !exis {
		return models.User{}, errors.New("no user with this ID")
	}
	return val, nil
}

func (m *MemoryStorage) DeleteUser(id uuid.UUID) error {
	_, exis := m.users[id]
	if !exis {
		return errors.New("no user with this ID")
	}
	maps.DeleteFunc(m.users, func(k uuid.UUID, v models.User) bool {
		return k == id
	})
	return nil
}
