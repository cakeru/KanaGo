# Visual Architecture & Quick Reference

## System Architecture Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                     USER BROWSER                           │
│  React SPA (TypeScript, Zustand, React Router)             │
│                                                             │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │   Pages      │  │  Components  │  │   Stores     │     │
│  │              │  │              │  │              │     │
│  │  • Login     │  │  • Button    │  │  • Auth      │     │
│  │  • Register  │  │  • Card      │  │  • Progress  │     │
│  │  • Game      │  │  • Modal     │  │  • Prefs     │     │
│  │  • Progress  │  │  • Form      │  │              │     │
│  │  • Settings  │  │              │  │              │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
│                                                             │
│                      HTTPS/WebSocket                       │
└─────────────────┬───────────────────────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────────────────────┐
│               GO REST API SERVER (Port 8080)                │
│                                                             │
│  ┌──────────────────────────────────────────────────────┐  │
│  │              HTTP Handlers                           │  │
│  │  • Auth Handlers                                     │  │
│  │  • Content Handlers (Kana, Kanji, Vocab)           │  │
│  │  • Progress Handlers                                │  │
│  │  • Stats Handlers                                   │  │
│  │  • Achievement Handlers                             │  │
│  └──────────────────────────────────────────────────────┘  │
│                          ▼                                  │
│  ┌──────────────────────────────────────────────────────┐  │
│  │            Business Logic (Services)                 │  │
│  │  • User Service (auth, validation)                  │  │
│  │  • Progress Service (tracking, mastery calc)        │  │
│  │  • Statistics Service (accuracy, streaks)           │  │
│  │  • Achievement Service (unlock logic)               │  │
│  └──────────────────────────────────────────────────────┘  │
│                          ▼                                  │
│  ┌──────────────────────────────────────────────────────┐  │
│  │          Data Access (Repositories)                  │  │
│  │  • User Repository                                  │  │
│  │  • Progress Repository                              │  │
│  │  • Statistics Repository                            │  │
│  │  • Achievement Repository                           │  │
│  └──────────────────────────────────────────────────────┘  │
└─────────────────┬───────────────────────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────────────────────┐
│             PostgreSQL Database (Port 5432)                 │
│                                                             │
│  Tables:                                                   │
│  ├── users (id, email, username, password_hash)          │
│  ├── statistics (user_id, total_answers, accuracy)       │
│  ├── progress (user_id, content_type, mastery_level)     │
│  ├── preferences (user_id, theme, font, language)        │
│  ├── achievements (id, name, description, points)        │
│  ├── user_achievements (user_id, achievement_id)         │
│  ├── practice_sessions (user_id, game_mode, duration)    │
│  └── answers (session_id, user_answer, is_correct)       │
└─────────────────────────────────────────────────────────────┘
```

---

## Data Flow: Complete Answer Submission

```
USER INTERACTION
      │
      ▼
[User sees Hiragana character: あ]
[User selects: "a" from multiple choice]
      │
      ▼
FRONTEND - GAME COMPONENT
┌─────────────────────────────────┐
│ validate answer                 │
│ is_correct = true              │
│ time_spent_ms = 2500           │
└────────────┬────────────────────┘
             │
             ▼ (HTTP POST /progress/update)
BACKEND - PROGRESS HANDLER
┌─────────────────────────────────┐
│ extract user_id from JWT        │
│ validate request                │
│ call progress service           │
└────────────┬────────────────────┘
             │
             ▼
SERVICE LAYER - PROGRESS SERVICE
┌─────────────────────────────────┐
│ find or create progress entry   │
│ update counts                   │
│ calculate mastery_level         │
│ update user statistics          │
│ check for achievements unlock   │
└────────────┬────────────────────┘
             │
             ▼
REPOSITORY LAYER
┌─────────────────────────────────┐
│ UPDATE progress SET             │
│   correct_count = correct_count + 1,
│   mastery_level = 100           │
│                                 │
│ UPDATE statistics SET           │
│   total_answers = total_answers + 1,
│   correct_answers = correct_answers + 1,
│   current_streak = current_streak + 1
│                                 │
│ INSERT answers (...)            │
└────────────┬────────────────────┘
             │
             ▼
DATABASE - PostgreSQL
┌─────────────────────────────────┐
│ PERSISTED TO DISK               │
└─────────────────────────────────┘
             │
             ▼ (RESPONSE: 200 OK)
FRONTEND - UPDATE LOCAL STATE
┌─────────────────────────────────┐
│ store.addProgress(...)          │
│ play correct sound              │
│ show feedback: "✓ Correct!"     │
│ update streak: 5                │
│ update score: +10               │
│ show next question button       │
└─────────────────────────────────┘
             │
             ▼
USER SEES FEEDBACK & CONTINUES
```

---

## API Endpoints Quick Reference

### Authentication (5 endpoints)
```
POST   /auth/register          → {user, access_token, refresh_token}
POST   /auth/login             → {user, access_token, refresh_token}
POST   /auth/refresh           → {access_token, refresh_token}
GET    /auth/me                → {user}
POST   /auth/logout            → {message}
```

### Content (6 endpoints)
```
GET    /kana                   → [kana_chars]
GET    /kanji                  → [kanji_chars]
GET    /kanji/jlpt/n5          → [kanji_by_level]
GET    /vocabulary             → [vocab_words]
GET    /vocabulary/jlpt/n5     → [vocab_by_level]
GET    /kana/:id               → {kana_char}
```

### Progress (5 endpoints)
```
POST   /progress/update        → {progress, accuracy, mastery_level}
GET    /stats                  → {total_answers, accuracy, streak}
GET    /stats/:contentType     → {content_stats}
GET    /progress/history       → [sessions]
GET    /streak                 → {current_streak, longest_streak}
```

### Preferences (2 endpoints)
```
GET    /preferences            → {theme, font, sound_enabled, ...}
PUT    /preferences            → {updated_preferences}
```

### Achievements (3 endpoints)
```
GET    /achievements           → [all_achievements]
GET    /achievements/unlocked  → [user_achievements]
POST   /achievements/check     → {unlocked_achievements, progress}
```

---

## Database Schema (Simplified)

```
USERS
├── id (PK, UUID)
├── email (UNIQUE)
├── username (UNIQUE)
├── password_hash
├── created_at
└── last_login

PROGRESS
├── id (PK)
├── user_id (FK → users.id)
├── content_type ("kana", "kanji", "vocab")
├── content_id
├── correct_count
├── wrong_count
├── mastery_level (0-100)
└── last_practiced

STATISTICS
├── id (PK)
├── user_id (FK, UNIQUE → users.id)
├── total_answers
├── correct_answers
├── accuracy
├── current_streak
├── longest_streak
└── last_practice_date

PREFERENCES
├── id (PK)
├── user_id (FK, UNIQUE → users.id)
├── theme_name
├── font_name
├── sound_enabled
├── dark_mode
└── language

ACHIEVEMENTS
├── id (PK)
├── name (UNIQUE)
├── description
└── icon_url

USER_ACHIEVEMENTS
├── id (PK)
├── user_id (FK → users.id)
├── achievement_id (FK → achievements.id)
└── unlocked_at

PRACTICE_SESSIONS
├── id (PK)
├── user_id (FK → users.id)
├── game_mode ("pick", "reverse_pick", ...)
├── started_at
├── ended_at
├── correct_answers
├── total_answers
└── duration_seconds

ANSWERS
├── id (PK)
├── session_id (FK → practice_sessions.id)
├── user_id (FK → users.id)
├── user_answer
├── correct_answer
├── is_correct
└── time_spent_ms
```

---

## State Management Structure

### Frontend Stores (Zustand)

```
useAuthStore
├── user: User | null
├── accessToken: string | null
├── refreshToken: string | null
├── login(email, password)
├── register(email, username, password)
└── logout()

usePreferencesStore
├── theme: string
├── font: string
├── soundEnabled: boolean
├── darkMode: boolean
├── language: string
├── setTheme(theme)
└── setFont(font)

useProgressStore
├── progress: {[key]: ContentProgress}
├── totalAnswers: number
├── correctAnswers: number
├── accuracy: number
├── currentStreak: number
├── longestStreak: number
└── addProgress(contentId, isCorrect)
```

---

## Game Mode Flowcharts

### Pick Mode (Show Character → User Selects Romanization)
```
Backend: Select random kana
    ↓
Backend: Generate 3 wrong romanizations
    ↓
Backend: Shuffle all 4 and send to frontend
    ↓
Frontend: Display character + 4 choices
    ↓
User: Click correct romanization
    ↓
Frontend: Send answer to backend
    ↓
Backend: Validate (is answer == correct romanization?)
    ↓
Backend: Update progress and stats
    ↓
Frontend: Display feedback and next button
```

### Reverse-Pick Mode (Show Romanization → User Selects Character)
```
(Same as above but with romanization shown instead of character)
```

### Input Mode (Show Character → User Types Romanization)
```
Backend: Select random kana
    ↓
Frontend: Display character
    ↓
User: Type romanization
    ↓
Frontend: On Enter key press
    ↓
Frontend: Send typed answer to backend
    ↓
Backend: Validate (case-insensitive comparison)
    ↓
Backend: Update progress
    ↓
Frontend: Show feedback
```

### Reverse-Input Mode (Show Romanization → User Types Character)
```
(Same as input mode but user types Japanese character instead)
```

---

## Component Hierarchy

```
App
├── Router
│   ├── Home (public)
│   ├── Login (public)
│   ├── Register (public)
│   ├── Kana (protected)
│   │   ├── KanaMenu
│   │   │   ├── SubsetSelector
│   │   │   └── GameModeSelector
│   │   └── KanaGame
│   │       ├── QuestionDisplay
│   │       ├── ChoiceButtons (for pick modes)
│   │       ├── InputField (for input modes)
│   │       ├── FeedbackMessage
│   │       └── NextButton
│   ├── Kanji (protected)
│   │   └── (similar to Kana)
│   ├── Vocabulary (protected)
│   │   └── (similar to Kana)
│   ├── Progress (protected)
│   │   ├── StatsOverview
│   │   ├── ContentTypeStats
│   │   ├── AccuracyChart
│   │   └── HistoryList
│   ├── Achievements (protected)
│   │   ├── AchievementGrid
│   │   ├── AchievementCard
│   │   └── ProgressBar
│   ├── Preferences (protected)
│   │   ├── ThemeSelector
│   │   ├── FontSelector
│   │   ├── SettingsToggle
│   │   └── LanguageSelector
│   └── Layout
│       ├── Header
│       │   ├── Logo
│       │   ├── Navigation
│       │   └── UserMenu
│       ├── Sidebar
│       │   └── NavItems
│       └── Footer
```

---

## Typical User Journey

```
1. DISCOVERY
   ├── Visit /
   └── See home page with game options

2. AUTHENTICATION
   ├── Click register or login
   ├── Enter credentials
   ├── Receive JWT tokens
   └── Tokens stored in localStorage

3. CONTENT SELECTION
   ├── Navigate to /kana (or /kanji, /vocabulary)
   ├── Select subset (hiragana, dakuon, etc.)
   ├── Select game mode (pick, reverse-pick, input, reverse-input)
   └── Start game

4. GAMEPLAY
   ├── See question displayed
   ├── Answer 1st question (feedback plays)
   ├── Move to next question
   ├── Continue 10-20 questions
   └── Session ends

5. PROGRESS UPDATE
   ├── All answers sent to backend
   ├── Progress table updated
   ├── Statistics recalculated
   ├── Achievements checked
   └── Frontend updated with new stats

6. REVIEW
   ├── Navigate to /progress
   ├── View overall statistics
   ├── See streaks and accuracy
   ├── Review practice history
   └── View achievements

7. CUSTOMIZATION
   ├── Navigate to /preferences
   ├── Change theme
   ├── Change font
   ├── Toggle sound
   ├── Select language
   └── Settings persist to database
```

---

## Technology Stack at a Glance

| Layer | Technology | Why? |
|-------|-----------|------|
| **Frontend** | React 18 | Rich ecosystem, component reusability |
| **Frontend** | TypeScript | Type safety, excellent DX |
| **Frontend** | Vite | Fast builds, great dev experience |
| **Frontend** | Zustand | Lightweight state management |
| **Frontend** | React Router | Standard routing solution |
| **Frontend** | Tailwind CSS | Utility-first, rapid development |
| **Frontend** | Vitest | Fast unit testing |
| **Backend** | Go 1.21+ | Fast, concurrent, great for APIs |
| **Backend** | Gin | Lightweight routing framework |
| **Backend** | PostgreSQL | Mature, reliable, JSON support |
| **Backend** | sqlx | Type-safe database access |
| **Backend** | JWT | Stateless authentication |
| **Auth** | bcrypt | Secure password hashing |
| **Deployment** | Docker | Containerization |
| **CI/CD** | GitHub Actions | Automated testing/deployment |
| **Hosting** | Vercel/Netlify | Frontend (auto-deploy) |
| **Hosting** | Railway/Heroku | Backend (simple deployment) |

---

## Performance Targets

```
Frontend
├── Page load: < 3 seconds
├── API call: < 100ms (p95)
├── Bundle size: < 200KB gzipped
├── Lighthouse score: > 90
└── Responsive on all devices

Backend
├── Response time: < 100ms (p95)
├── Database query: < 50ms average
├── Throughput: > 1000 req/sec
├── Uptime: > 99.9%
└── Connection pool: 25 connections

Database
├── Query optimization: Index all FK
├── Connection pooling: 25 max connections
├── Backup frequency: Daily
└── Replication: Optional for scale
```

---

## Deployment Architecture

```
Production Environment
├── Frontend
│   ├── Hosted on Vercel/Netlify
│   ├── CDN for static assets
│   └── HTTPS enforced
├── Backend
│   ├── Docker container
│   ├── Hosted on Railway/Heroku
│   ├── Environment variables injected
│   └── Auto-scaling enabled
└── Database
    ├── PostgreSQL managed service
    ├── Automated backups
    ├── Read replicas (optional)
    └── SSL encryption
```

---

This visual guide should help you understand the complete system at a glance!
