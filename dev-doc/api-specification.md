# API Specification: KanaDojo Go + React

Complete RESTful API specification for the Japanese learning platform.

---

## Base Information

- **Base URL**: `http://localhost:8080/api/v1` (development)
- **Content Type**: `application/json`
- **Authentication**: Bearer Token (JWT)
- **Response Format**: Standardized JSON with `success`, `data`, and `error` fields

---

## Response Format

### Success Response (2xx)

```json
{
  "success": true,
  "data": {
    // Response data here
  },
  "message": "Optional message"
}
```

### Error Response (4xx, 5xx)

```json
{
  "success": false,
  "error": "Error description",
  "message": "Optional message"
}
```

---

## Authentication Endpoints

### 1. Register User

**POST** `/auth/register`

Register a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "username": "username",
  "password": "securePassword123"
}
```

**Validation Rules:**
- `email`: Required, valid email format, unique
- `username`: Required, 3-100 characters, unique, alphanumeric + underscores
- `password`: Required, minimum 8 characters

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "user@example.com",
      "username": "username",
      "email_verified": false,
      "created_at": "2024-01-15T10:30:00Z"
    },
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
  },
  "message": "User registered successfully"
}
```

**Error Responses:**
- `400 Bad Request`: Validation failed
- `409 Conflict`: Email or username already exists

---

### 2. Login User

**POST** `/auth/login`

Authenticate user and return JWT tokens.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "user@example.com",
      "username": "username",
      "email_verified": true,
      "last_login": "2024-01-15T10:30:00Z"
    },
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

**Error Responses:**
- `401 Unauthorized`: Invalid credentials
- `404 Not Found`: User not found

---

### 3. Refresh Token

**POST** `/auth/refresh`

Get new access token using refresh token.

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

**Error Responses:**
- `401 Unauthorized`: Invalid or expired refresh token

---

### 4. Get Current User

**GET** `/auth/me`

Get the authenticated user's profile.

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "user@example.com",
    "username": "username",
    "email_verified": true,
    "last_login": "2024-01-15T10:30:00Z",
    "created_at": "2024-01-10T08:00:00Z"
  }
}
```

**Error Responses:**
- `401 Unauthorized`: Missing or invalid token

---

### 5. Logout

**POST** `/auth/logout`

Logout user (client should discard tokens).

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Logged out successfully"
}
```

---

## Content Endpoints

### 1. Get All Kana

**GET** `/kana`

Get all hiragana and katakana characters.

**Query Parameters:**
- `type` (optional): "hiragana" or "katakana"
- `subset` (optional): "base", "dakuon", "yoon", "foreign"

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": "kana-a",
      "character": "あ",
      "romanization": "a",
      "type": "hiragana",
      "subset": "base"
    },
    {
      "id": "kana-b",
      "character": "い",
      "romanization": "i",
      "type": "hiragana",
      "subset": "base"
    }
  ]
}
```

---

### 2. Get Kana by ID

**GET** `/kana/:id`

Get specific kana character by ID.

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "kana-a",
    "character": "あ",
    "romanization": "a",
    "type": "hiragana",
    "subset": "base"
  }
}
```

---

### 3. Get All Kanji

**GET** `/kanji`

Get all kanji characters with optional filtering.

**Query Parameters:**
- `jlpt` (optional): "n5", "n4", "n3", "n2", "n1"
- `limit` (optional): Maximum results (default: 50, max: 500)
- `offset` (optional): Pagination offset (default: 0)

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": "kanji-001",
      "character": "漢",
      "jlpt_level": "n2",
      "meanings": ["Sino-", "China"],
      "onyomi": ["カン"],
      "kunyomi": ["から"],
      "strokes": 13
    }
  ]
}
```

---

### 4. Get Kanji by JLPT Level

**GET** `/kanji/jlpt/:level`

Get all kanji for a specific JLPT level.

**Path Parameters:**
- `level`: "n5", "n4", "n3", "n2", or "n1"

**Query Parameters:**
- `limit` (optional): Maximum results
- `offset` (optional): Pagination offset

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "jlpt_level": "n5",
    "total_count": 100,
    "kanji": [
      {
        "id": "kanji-001",
        "character": "一",
        "meanings": ["one"],
        "onyomi": ["イチ"],
        "kunyomi": ["ひと"],
        "strokes": 1
      }
    ]
  }
}
```

---

### 5. Get All Vocabulary

**GET** `/vocabulary`

Get vocabulary words with optional filtering.

**Query Parameters:**
- `jlpt` (optional): "n5", "n4", "n3", "n2", "n1"
- `part_of_speech` (optional): "noun", "verb", "adjective", "adverb"
- `limit` (optional): Maximum results
- `offset` (optional): Pagination offset

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": "vocab-001",
      "word": "日本",
      "reading": "にほん",
      "romaji": "nihon",
      "meaning": "Japan",
      "jlpt_level": "n5",
      "part_of_speech": "noun"
    }
  ]
}
```

---

### 6. Get Vocabulary by JLPT Level

**GET** `/vocabulary/jlpt/:level`

Get all vocabulary for a specific JLPT level.

**Path Parameters:**
- `level`: "n5", "n4", "n3", "n2", or "n1"

**Query Parameters:**
- `limit` (optional): Maximum results
- `offset` (optional): Pagination offset

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "jlpt_level": "n5",
    "total_count": 800,
    "vocabulary": [
      {
        "id": "vocab-001",
        "word": "私",
        "reading": "わたし",
        "romaji": "watashi",
        "meaning": "I, me",
        "part_of_speech": "pronoun"
      }
    ]
  }
}
```

---

## Progress & Statistics Endpoints

### 1. Get User Statistics

**GET** `/stats`

Get aggregated user statistics.

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "total_answers": 5000,
    "correct_answers": 4750,
    "accuracy": 95.0,
    "current_streak": 12,
    "longest_streak": 47,
    "last_practice_date": "2024-01-15T10:30:00Z",
    "content_stats": {
      "kana": {
        "total_answers": 2000,
        "correct_answers": 1950,
        "accuracy": 97.5
      },
      "kanji": {
        "total_answers": 2000,
        "correct_answers": 1850,
        "accuracy": 92.5
      },
      "vocabulary": {
        "total_answers": 1000,
        "correct_answers": 950,
        "accuracy": 95.0
      }
    }
  }
}
```

---

### 2. Get Content-Specific Statistics

**GET** `/stats/:contentType`

Get statistics for specific content type (kana, kanji, vocabulary).

**Path Parameters:**
- `contentType`: "kana", "kanji", or "vocabulary"

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "content_type": "kana",
    "total_items": 46,
    "mastered_items": 40,
    "in_progress_items": 5,
    "accuracy": 97.5,
    "total_answers": 2000,
    "correct_answers": 1950,
    "practice_time_minutes": 180
  }
}
```

---

### 3. Update User Progress

**POST** `/progress/update`

Record an answer and update progress.

**Headers:**
```
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
  "content_type": "kana",
  "content_id": "kana-a",
  "is_correct": true,
  "time_spent_ms": 2500
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "progress": {
      "id": "progress-001",
      "user_id": "user-001",
      "content_type": "kana",
      "content_id": "kana-a",
      "correct_count": 15,
      "wrong_count": 2,
      "mastery_level": 88,
      "last_practiced": "2024-01-15T10:35:00Z"
    },
    "accuracy": 88.2,
    "mastery_level": 88,
    "streak_updated": true,
    "current_streak": 12,
    "achievement_unlocked": null
  }
}
```

---

### 4. Get Progress History

**GET** `/progress/history`

Get detailed practice history.

**Headers:**
```
Authorization: Bearer {access_token}
```

**Query Parameters:**
- `content_type` (optional): Filter by content type
- `limit` (optional): Maximum results (default: 50)
- `offset` (optional): Pagination offset
- `from_date` (optional): ISO 8601 date
- `to_date` (optional): ISO 8601 date

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "total_count": 5000,
    "history": [
      {
        "id": "session-001",
        "content_type": "kana",
        "game_mode": "pick",
        "started_at": "2024-01-15T10:00:00Z",
        "ended_at": "2024-01-15T10:30:00Z",
        "duration_seconds": 1800,
        "correct_answers": 28,
        "total_answers": 30,
        "accuracy": 93.3
      }
    ]
  }
}
```

---

### 5. Get Current Streak

**GET** `/streak`

Get current and longest streak information.

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "current_streak": 12,
    "longest_streak": 47,
    "streak_last_updated": "2024-01-15T10:35:00Z",
    "days_since_last_practice": 0
  }
}
```

---

## Preferences Endpoints

### 1. Get User Preferences

**GET** `/preferences`

Get user's customization preferences.

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "pref-001",
    "user_id": "user-001",
    "theme_name": "nord",
    "font_name": "noto_sans_jp",
    "sound_enabled": true,
    "hotkeys_enabled": true,
    "dark_mode": true,
    "language": "en"
  }
}
```

---

### 2. Update User Preferences

**PUT** `/preferences`

Update user preferences.

**Headers:**
```
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
  "theme_name": "dracula",
  "font_name": "noto_serif_jp",
  "sound_enabled": false,
  "hotkeys_enabled": true,
  "dark_mode": false,
  "language": "ja"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "pref-001",
    "user_id": "user-001",
    "theme_name": "dracula",
    "font_name": "noto_serif_jp",
    "sound_enabled": false,
    "hotkeys_enabled": true,
    "dark_mode": false,
    "language": "ja"
  },
  "message": "Preferences updated successfully"
}
```

---

## Achievements Endpoints

### 1. Get All Achievements

**GET** `/achievements`

Get list of all available achievements.

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": "achievement-001",
      "name": "First Steps",
      "description": "Complete your first 10 answers",
      "icon_url": "/icons/first-steps.png",
      "points": 10
    },
    {
      "id": "achievement-002",
      "name": "Streak Master",
      "description": "Reach a 100+ day streak",
      "icon_url": "/icons/streak-master.png",
      "points": 50
    }
  ]
}
```

---

### 2. Get User's Unlocked Achievements

**GET** `/achievements/unlocked`

Get achievements the user has unlocked.

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "total_achievements": 50,
    "unlocked_count": 15,
    "achievements": [
      {
        "id": "achievement-001",
        "name": "First Steps",
        "description": "Complete your first 10 answers",
        "icon_url": "/icons/first-steps.png",
        "unlocked_at": "2024-01-10T08:00:00Z",
        "points": 10
      }
    ]
  }
}
```

---

### 3. Check Achievements

**POST** `/achievements/check`

Check and unlock achievements if conditions are met.

**Headers:**
```
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
  "content_type": "kana",
  "action": "completed_session"
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "unlocked_achievements": [
      {
        "id": "achievement-025",
        "name": "Kana Master",
        "description": "Achieve 100% accuracy on all kana characters",
        "icon_url": "/icons/kana-master.png"
      }
    ],
    "progress": {
      "achievement_id": "achievement-026",
      "name": "Accuracy Striker",
      "current_progress": 92,
      "required_progress": 95
    }
  }
}
```

---

## Error Codes

| Code | Status | Description                          |
| ---- | ------ | ------------------------------------ |
| 400  | Bad Request | Invalid request parameters           |
| 401  | Unauthorized | Missing or invalid authentication    |
| 403  | Forbidden | User lacks permission                |
| 404  | Not Found | Resource not found                   |
| 409  | Conflict | Resource already exists (email/user) |
| 422  | Unprocessable Entity | Validation failed                    |
| 500  | Internal Server Error | Server error                         |

---

## Rate Limiting

- **Requests per minute**: 60 (public endpoints), 100 (authenticated)
- **Header**: `X-RateLimit-Remaining`
- **Exceeding limit response**: `429 Too Many Requests`

---

## Pagination

For endpoints returning lists:

**Query Parameters:**
- `limit`: Items per page (default: 50, max: 500)
- `offset`: Starting position (default: 0)

**Response:**
```json
{
  "success": true,
  "data": {
    "items": [...],
    "pagination": {
      "limit": 50,
      "offset": 0,
      "total_count": 1000
    }
  }
}
```

---

## Notes

- All timestamps are in ISO 8601 format with UTC timezone
- Token expiry: Access token (15 minutes), Refresh token (7 days)
- All endpoints require `Content-Type: application/json` for requests
- Empty/null response data is omitted from responses
