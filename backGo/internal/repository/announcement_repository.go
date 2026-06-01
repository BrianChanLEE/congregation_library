package repository

import (
	"boock/backGo/internal/models"
	"database/sql"
)

// AnnouncementRepositoryInterface는 공지사항 저장소의 동작을 정의합니다.
type AnnouncementRepositoryInterface interface {
	Create(a *models.Announcement) error
	GetAll() ([]models.Announcement, error)
	Delete(id int64) error
	Update(a *models.Announcement) error
}

type AnnouncementRepository struct {
	db *sql.DB
}

func NewAnnouncementRepository(db *sql.DB) *AnnouncementRepository {
	return &AnnouncementRepository{db: db}
}

func (r *AnnouncementRepository) Create(a *models.Announcement) error {
	query := "INSERT INTO announcements (title, content, author_id) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, a.Title, a.Content, a.AuthorID)
	return err
}

func (r *AnnouncementRepository) GetAll() ([]models.Announcement, error) {
	query := "SELECT id, title, content, author_id, created_at FROM announcements ORDER BY created_at DESC"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var anns []models.Announcement
	for rows.Next() {
		var a models.Announcement
		if err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.AuthorID, &a.CreatedAt); err != nil {
			return nil, err
		}
		anns = append(anns, a)
	}
	return anns, nil
}

func (r *AnnouncementRepository) Delete(id int64) error {
	query := "DELETE FROM announcements WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *AnnouncementRepository) Update(a *models.Announcement) error {
	query := "UPDATE announcements SET title = ?, content = ? WHERE id = ?"
	_, err := r.db.Exec(query, a.Title, a.Content, a.ID)
	return err
}
