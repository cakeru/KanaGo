package services

import (
	"fmt"

	"github.com/cakeru/kanago-backend/internal/repository"
)

// KanaService handles buisness logic for Kana
type KanaService struct {
	kanaRepo *repository.KanaRepository
}

// NewKanaService creates a new KanaService
func NewKanaService() *KanaService {
	return &KanaService{
		kanaRepo: repository.NewKanaRepository(),
	}
}

// GetKana retrieves a kana character by its ID
func (s *KanaService) GetKana(id int) (map[string]interface{}, error) {
	// validate input
	if id <= 0 {
		return nil, fmt.Errorf("invalid kana ID")
	}

	// Get from repository
	kana, err := s.kanaRepo.GetKanaByID(id)
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"id":         kana.ID,
		"character":  kana.Character,
		"romaji":     kana.Romaji,
		"type":       kana.Type,
		"created_at": kana.CreatedAt,
	}

	return response, nil
}

// GetKanaByType retrieves all kana characters of a specific type (hiragana or katakana)
func (s *KanaService) GetKanaByType(kanaType string) ([]map[string]interface{}, error) {
	// validate input
	if kanaType == "" {
		return nil, fmt.Errorf("kana type is required")
	}

	if kanaType != "hiragana" && kanaType != "katakana" {
		return nil, fmt.Errorf("invalid kana type, must be 'hiragana' or 'katakana'")
	}

	// Get from repository
	kanas, err := s.kanaRepo.GetKanaByType(kanaType)
	if err != nil {
		return nil, err
	}

	//Build Response
	var response []map[string]interface{}
	for _, kana := range kanas {
		item := map[string]interface{}{
			"id":         kana.ID,
			"character":  kana.Character,
			"romaji":     kana.Romaji,
			"type":       kana.Type,
			"created_at": kana.CreatedAt,
		}
		response = append(response, item)
	}
	
	return response, nil
}

// GetAllKana retrieves all kana characters
func (s *KanaService) GetAllKana() (map[string]interface{}, error) {
	// Get from repository
	kanas, err := s.kanaRepo.GetAllKana()
	if err != nil {
		return nil, err
	}

	// Seperate by type
	var hiraganas []map[string]interface{}
	var katakanas []map[string]interface{}

	for _, kana := range kanas {
		item := map[string]interface{}{
			"id":         kana.ID,
			"character":  kana.Character,
			"romaji":     kana.Romaji,
			"type":       kana.Type,
			"created_at": kana.CreatedAt,
		}

		if kana.Type == "hiragana" {
			hiraganas = append(hiraganas, item)
		} else if kana.Type == "katakana" {
			katakanas = append(katakanas, item)
		}
	}

	// Build Response
	response := map[string]interface{}{
		"hiragana": hiraganas,
		"katakana": katakanas,
		"total":    len(kanas),
	}
	return response, nil
}