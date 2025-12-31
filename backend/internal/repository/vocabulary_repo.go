package repository

import (
	"database/sql"
	"fmt"

	"github.com/cakeru/kanago-backend/internal/db"
	"github.com/cakeru/kanago-backend/internal/models"
)

// VocabularyRepository handles database operations for Vocabulary
type VocabularyRepository struct{}

// NewVocabularyRepository creates a new vocabulary repository
func NewVocabularyRepository() *VocabularyRepository {
	return &VocabularyRepository{}
}

// GetVocabularyByID retrieves a vocabulary word by its ID
func (r *VocabularyRepository) GetVocabularyByID(id int) (*models.Vocabulary, error) {
	vocab := &models.Vocabulary{}

	query := `
		SELECT id, word, reading, meaning, part_of_speech, example, created_at
		FROM vocabulary
		WHERE id = $1
		`

	err := db.DB.QueryRow(query, id).
		Scan(&vocab.ID, &vocab.Word, &vocab.Reading, &vocab.Meaning, &vocab.PartOfSpeech, &vocab.Example, &vocab.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Vocabulary not found")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get vocabulary: %w", err)
	}

	return vocab, nil
}

// GetVocabularyByPartOfSpeech retrieves vocabulary words by their part of speech
func (r *VocabularyRepository) GetVocabularyByPartOfSpeech(partOfSpeech string) ([]models.Vocabulary, error) {
	query := `
		SELECT id, word, reading, meaning, part_of_speech, example, created_at
		FROM vocabulary
		WHERE part_of_speech = $1
		ORDER BY id ASC
		`

	rows, err := db.DB.Query(query, partOfSpeech)
	if err != nil {
		return nil, fmt.Errorf("failed to query vocabulary by part of speech: %w", err)
	}
	defer rows.Close()

	var vocabs []models.Vocabulary
	for rows.Next() {
		var vocab models.Vocabulary
		err := rows.Scan(&vocab.ID, &vocab.Word, &vocab.Reading, &vocab.Meaning, &vocab.PartOfSpeech, &vocab.Example, &vocab.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan vocabulary: %w", err)
		}
		vocabs = append(vocabs, vocab)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating vocabulary: %w", err)
	}

	return vocabs, nil
}

// GetAllVocabulary retrieves all vocabulary with pagination
func (r *VocabularyRepository) GetAllVocabulary(limit, offset int) ([]models.Vocabulary, error) {
	query := `
		SELECT id, word, reading, meaning, part_of_speech, example, created_at
		FROM vocabulary
		ORDER BY id ASC
		LIMIT $1 OFFSET $2
		`
	
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query all vocabulary: %w", err)
	}
	defer rows.Close()

	var vocabs []models.Vocabulary
	for rows.Next() {
		var vocab models.Vocabulary
		err := rows.Scan(&vocab.ID, &vocab.Word, &vocab.Reading, &vocab.Meaning, &vocab.PartOfSpeech, &vocab.Example, &vocab.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan vocabulary: %w", err)
		}
		vocabs = append(vocabs, vocab)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating vocabulary: %w", err)
	}

	return vocabs, nil
}