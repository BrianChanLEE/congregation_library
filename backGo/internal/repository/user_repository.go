package repository

import (
	"boock/backGo/internal/models"
	"database/sql"
)

// UserRepositoryInterface는 사용자 저장소의 동작을 정의합니다.
type UserRepositoryInterface interface {
	Create(user *models.User) error
	GetByID(id int64) (*models.User, error)
	GetByJwhubEmailAndCongID(email string, congID int64) (*models.User, error)
	UpdateStatus(userID int64, status string) error
	GetAllPending() ([]models.User, error)
	Delete(userID int64) error
	GetPasswordHash(userID int64) (string, error)
	UpdatePassword(userID int64, hashedNewPassword string) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := "INSERT INTO users (name, jwhub_email, password_hash, status) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.Name, user.JWhubEmail, user.PasswordHash, user.Status)
	return err
}

func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	query := "SELECT id, name, jwhub_email, role, status FROM users WHERE id = ? AND deleted_at IS NULL"
	var user models.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.JWhubEmail, &user.Role, &user.Status)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByJwhubEmailAndCongID(email string, congID int64) (*models.User, error) {
	query := "SELECT id, name, jwhub_email, password_hash, role FROM users WHERE jwhub_email = ? AND congregation_id = ? AND deleted_at IS NULL"
	var user models.User
	err := r.db.QueryRow(query, email, congID).Scan(&user.ID, &user.Name, &user.JWhubEmail, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateStatus(userID int64, status string) error {
	query := "UPDATE users SET status = ? WHERE id = ? AND deleted_at IS NULL"
	_, err := r.db.Exec(query, status, userID)
	return err
}

func (r *UserRepository) GetAllPending() ([]models.User, error) {
	query := "SELECT id, name, jwhub_email, status FROM users WHERE status = 'PENDING' AND deleted_at IS NULL"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.JWhubEmail, &user.Status); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Delete(userID int64) error {
	query := "UPDATE users SET deleted_at = NOW() WHERE id = ?"
	_, err := r.db.Exec(query, userID)
	return err
}

func (r *UserRepository) GetPasswordHash(userID int64) (string, error) {
	query := "SELECT password_hash FROM users WHERE id = ? AND deleted_at IS NULL"
	var hash string
	err := r.db.QueryRow(query, userID).Scan(&hash)
	return hash, err
}

func (r *UserRepository) UpdatePassword(userID int64, hashedNewPassword string) error {
	query := "UPDATE users SET password_hash = ? WHERE id = ? AND deleted_at IS NULL"
	_, err := r.db.Exec(query, hashedNewPassword, userID)
	return err
}
