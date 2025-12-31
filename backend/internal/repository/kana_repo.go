package repository

import (
	"database/sql"
	"fmt"

	"github.com/cakeru/kanago-backend/internal/db"
	"github.com/cakeru/kanago-backend/internal/models"
)

// KanaRepository handles database operations for Kana
type KanaRepository struct{}

// NewKanaRepository creates a new kana repository
func NewKanaRepository() *KanaRepository {
	return &KanaRepository{}
}

//GetKanaByID retrieves a kana character by its ID
func (r *KanaRepository) GetKanaByID(id int) (*models.Kana, error) {
	kana := &models.Kana{}

	query := `SELECT id, character, romaji, type, created_at FROM kana WHERE id = $1`

	err := db.DB.QueryRow(query, id).Scan(&kana.ID, &kana.Character, &kana.Romaji, &kana.Type, &kana.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("kana not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get kana: %w", err)
	}
	return kana, nil
}

// GetKanaBytype retrieves all kana characters of a specific type (hiragana or katakana)
func (r *KanaRepository) GetKanaByType(kanaType string) ([]models.Kana, error){
	query := `
	     SELECT id, character, romaji, type, created_at
		 FROM kana
		 WHERE type =$1
		 ORDER BY id ASC
		 `

	rows, err := db.DB.Query(query, kanaType)
	if err != nil {
		return nil, fmt.Errorf("failed to query kana by type: %w", err)
	}
	defer rows.Close()

	var kanas []models.Kana
	for rows.Next() {
		var kana models.Kana
		if err := rows.Scan(&kana.ID, &kana.Character, &kana.Romaji, &kana.Type, &kana.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan kana: %w", err)
		}
		kanas = append(kanas, kana)
	}
	
	if err = rows.Err(); err != nil{
		return nil, fmt.Errorf("error iteration kana: %w", err)
	}

	return kanas, nil
}

// GetAllKana retrives all Kana
func (r *KanaRepository) GetAllKana() ([]models.Kana, error){
	query := `
	     SELECT id, character, romaji, type, created_at
		 FROM kana
		 ORDER by id ASC
		 `
	
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query all kana: %w", err)
	}
	defer rows.Close()

	var kanas []models.Kana
	for rows.Next() {
		var kana models.Kana
		err := rows.Scan(&kana.ID, &kana.Character, &kana.Romaji, &kana.Type, &kana.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to get all kana: %w", err)
		}
		kanas = append(kanas, kana)
	}

	if err = rows.Err(); err != nil{
		return nil, fmt.Errorf("error iteration kana: %w", err)
	}
	return kanas, nil
}