package repository

import (
	"database/sql"
	"fmt"

	"github.com/cakeru/kanago-backend/internal/db"
	"github.com/cakeru/kanago-backend/internal/models"
)

// KanjiRepository handles database operations for Kanji
type KanjiRepository struct{}

// NewKanjiRepository creates a new kanji repository
func NewKanjiRepository() *KanjiRepository {
	return &KanjiRepository{}
}

// GetKanjiByID retrieves a kanji by its ID
func (r *KanjiRepository) GetKanjiByID(id int) (*models.Kanji, error) {
	kanji := &models.Kanji{}

	query := `
		SELECT id, character, meaning, on_yomi, kun_yomi, strokes, created_at
		FROM kanji
		WHERE id = $1
	`

	err := db.DB.QueryRow(query, id).
		Scan(&kanji.ID, &kanji.Character, &kanji.Meaning, &kanji.OnYomi, &kanji.KunYomi, &kanji.Strokes, &kanji.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("kanji not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get kanji: %w", err)
	}
	return kanji, nil
}

func (r *KanjiRepository) GetKanjiByStrokes(strokes int) ([]models.Kanji, error){
	query := `
		SELECT id, character, meaning, on_yomi, kun_yomi, strokes, created_at
		FROM kanji
		WHERE strokes = $1
		ORDER BY id ASC
		`

	rows, err := db.DB.Query(query, strokes)
	if err != nil {
		return nil, fmt.Errorf("failed to query kanji by strokes: %w", err)
	}
	defer rows.Close()

	var kanjis []models.Kanji
	for rows.Next() {
		var kanji models.Kanji
		err := rows.Scan(&kanji.ID, &kanji.Character, &kanji.Meaning, &kanji.OnYomi, &kanji.KunYomi, &kanji.Strokes, &kanji.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan kanji: %w", err)
		}
		kanjis = append(kanjis, kanji)
	}
	if err = rows.Err(); err != nil{
		return nil, fmt.Errorf("error iterating kanji: %w", err)
	}
	return kanjis, nil
}

// GetAllKanji retrieves all Kanji
func (r *KanjiRepository) GetAllKanji(limit, offset int) ([]models.Kanji, error) {
	query := `
		SELECT id, character, meaning, on_yomi, kun_yomi, strokes, created_at
		FROM kanji
		ORDER BY strokes ASC, id ASC
		LIMIT $1 OFFSET $2
		`

	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query all kanji: %w", err)
	}
	defer rows.Close()

	var kanjis []models.Kanji
	for rows.Next() {
		var kanji models.Kanji
		err := rows.Scan(&kanji.ID, &kanji.Character, &kanji.Meaning, &kanji.OnYomi, &kanji.KunYomi, &kanji.Strokes, &kanji.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan kanji: %w", err)
		}
		kanjis = append(kanjis, kanji)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating kanji: %w", err)
	}
	return kanjis, nil
}