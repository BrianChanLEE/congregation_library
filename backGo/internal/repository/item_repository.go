package repository

import (
	"boock/backGo/internal/models"
	"database/sql"
)

// ItemRepositoryInterface는 품목 저장소의 동작을 정의합니다.
type ItemRepositoryInterface interface {
	Create(item *models.Item) error
	GetAll() ([]models.Item, error)
	GetCatalog() ([]models.Item, error)
	GetInventory(congID, itemID int64) (models.Inventory, error)
	Update(item *models.Item) error
	Delete(id int64) error
	GetCategories() ([]string, error)
}

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(item *models.Item) error {
	query := "INSERT INTO items (code, name) VALUES (?, ?)"
	_, err := r.db.Exec(query, item.Code, item.Name)
	return err
}

func (r *ItemRepository) Update(item *models.Item) error {
	query := "UPDATE items SET name = ?, code = ? WHERE id = ?"
	_, err := r.db.Exec(query, item.Name, item.Code, item.ID)
	return err
}

func (r *ItemRepository) GetAll() ([]models.Item, error) {
	query := "SELECT id, code, name FROM items WHERE deleted_at IS NULL"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Code, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ItemRepository) GetCatalog() ([]models.Item, error) {
	query := "SELECT id, code, name FROM items WHERE deleted_at IS NULL"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Code, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ItemRepository) GetInventory(congID, itemID int64) (models.Inventory, error) {
	query := `
		SELECT i.id, i.code, i.name, 
		       COALESCE(SUM(CASE WHEN a.type='IN' THEN a.quantity ELSE -a.quantity END), 0) as stock
		FROM items i
		LEFT JOIN activity_logs a ON i.id = a.item_id
	WHERE i.id = ? AND i.deleted_at IS NULL
	GROUP BY i.id`

	row := r.db.QueryRow(query, itemID)
	var inv models.Inventory
	err := row.Scan(&inv.ID, &inv.Code, &inv.Name, &inv.Stock)
	if err != nil {
		return models.Inventory{}, err
	}
	return inv, nil
}

func (r *ItemRepository) Delete(id int64) error {
	query := "UPDATE items SET deleted_at = NOW() WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ItemRepository) GetCategories() ([]string, error) {
	rows, err := r.db.Query("SELECT DISTINCT category FROM items WHERE category IS NOT NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var cat string
		if err := rows.Scan(&cat); err != nil {
			continue
		}
		categories = append(categories, cat)
	}
	return categories, nil
}
