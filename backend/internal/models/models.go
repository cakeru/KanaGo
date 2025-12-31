package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID		     int    `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Progress represents user's progress on learning items
type Progress struct {
	ID		int       `json:"id"`
	UserID		int       `json:"user_id"`
	ContentID	int       `json:"content_id"`
	ContentType string    `json:"content_type"` // kana, kanji, vocabulary
	CorrectCount int 	 `json:"correct_count"`
	IncorrectCount int    `json:"incorrect_count"`
	MasteryLevel int 	 `json:"mastery_level"` //0-100
	LastReviewedAt *time.Time `json:"last_reviewed_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

//Statistics represents user statistics for reviews and accuracy
type Statistics struct {
	ID			  int       `json:"id"`
	UserID		  int       `json:"user_id"`
	TotalReviews  int       `json:"total_reviews"`
	CorrectReviews int      `json:"correct_reviews"`
	CurrentStreak int 	 `json:"current_streak"`
	LongestStreak  int       `json:"longest_streak"`
	AccuracyPercentage float64   `json:"accuracy_percentage"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Preference represents user preferences for the application
type Preference struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Theme     string `json:"theme"` // light, dark
	Font      string `json:"font"`
	Language  string `json:"language"` //en, jp, etc.
	SoundEnabled bool   `json:"sound_enabled"`
	UpdatedAt    time.Time `json:"updated_at"`
}

//PracticeSession represents a game session
type PracticeSession struct {
	ID 	   int       `json:"id"`
	UserID    int       `json:"user_id"`
	GameMode  string    `json:"game_mode"` // flashcards, multiple_choice, typing
	ContentType string    `json:"content_type"` // kana, kanji, vocabulary
	DurationSeconds int       `json:"duration_seconds"`
	CorrectCount int       `json:"correct_count"`
	IncorrectCount int     `json:"incorrect_count"`
	CreatedAt time.Time `json:"created_at"`
}

//Answer Represents a single answer in a session
type Answer struct {
	ID			  int       `json:"id"`
	SessionID 	  int       `json:"session_id"`
	QuestionID  int       `json:"question_id"`
	UserAnswer  string    `json:"user_answer"`
	IsCorrect   bool      `json:"is_correct"`
	TimeSpentMs int       `json:"time_spent_ms"`
	CreatedAt	time.Time `json:"created_at"`
}

type Kana struct {
	ID 	  int    `json:"id"`
	Character string `json:"character"`
	Romaji    string `json:"romaji"`
	Type      string `json:"type"` // hiragana, katakana
	CreatedAt time.Time `json:"created_at"`
}

type Kanji struct {
	ID        int    `json:"id"`
	Character string `json:"character"`
	Meaning  string `json:"meaning"`
	OnYomi    string `json:"onyomi"`
	KunYomi   string `json:"kunyomi"`
	Strokes   int    `json:"strokes"`
	CreatedAt time.Time `json:"created_at"`
}

type Vocabulary struct {
	ID        int    `json:"id"`
	Word      string `json:"word"`
	Meaning   string `json:"meaning"`
	Reading   string `json:"reading"`
	PartOfSpeech string `json:"part_of_speech"`
	Example string `json:"example"`
	CreatedAt time.Time `json:"created_at"`
}