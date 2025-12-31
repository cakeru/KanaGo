# Building KanaDojo with Go + React: Complete Architecture Plan

## 📋 Executive Summary

KanaDojo is a sophisticated Japanese learning platform with:
- **Multiple learning dojos** (Kana, Kanji, Vocabulary, Calligraphy)
- **4 game modes per dojo** (Pick, Reverse-Pick, Input, Reverse-Input)
- **100+ themes, 28 fonts**, extensive customization
- **Progress tracking, achievements, streaks**
- **Multi-language support** (EN, ES, JA)
- **Real-time feedback with audio/visual effects**

By building with **Go + React**, you'll have a powerful, scalable foundation with Go's performance for the backend and React's rich UI capabilities for the frontend.

---

## 🏗️ Architecture Overview

### Layer Stack

```
┌───────────────────────────────────────┐
│    React Frontend (Single-Page App)   │
│  Components, State, Hooks, Routing    │
└──────────────┬────────────────────────┘
               │ HTTP/WebSocket
┌──────────────▼────────────────────────┐
│      Go Backend (REST/GraphQL)        │
│ Routing, Business Logic, Middleware   │
└──────────────┬────────────────────────┘
               │ SQL
┌──────────────▼────────────────────────┐
│        PostgreSQL Database            │
│ Users, Progress, Stats, Achievements  │
└───────────────────────────────────────┘
```

### Technology Stack

| Layer     | Technology                  | Purpose                                |
| --------- | --------------------------- | -------------------------------------- |
| Frontend  | React 18+ with TypeScript   | UI components, state, routing          |
| CSS       | Tailwind CSS + shadcn       | Styling, component library             |
| State     | Zustand / Context API       | Client-side state management           |
| Backend   | Go 1.21+                    | REST API, business logic               |
| Framework | Gin / Echo / Chi            | HTTP routing, middleware               |
| Database  | PostgreSQL 14+              | Persistent data storage                |
| Cache     | Redis (optional)            | Session caching, real-time features    |
| Auth      | JWT + Refresh Tokens        | Secure user authentication             |
| Files     | JSON files in public folder | Static content (kana, kanji, vocab)    |

---

## 📁 Directory Structure

### Backend (Go)

```
go-kanadojo-backend/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/            # HTTP request handlers
│   │   │   ├── kana.go
│   │   │   ├── kanji.go
│   │   │   ├── vocabulary.go
│   │   │   ├── auth.go
│   │   │   ├── user.go
│   │   │   ├── stats.go
│   │   │   └── achievements.go
│   │   ├── middleware/          # Middleware (auth, CORS, logging)
│   │   ├── routes.go            # Route registration
│   │   └── response.go          # Standard response formatter
│   ├── models/                  # Data models
│   │   ├── user.go
│   │   ├── progress.go
│   │   ├── achievement.go
│   │   ├── preference.go
│   │   └── content.go
│   ├── services/                # Business logic
│   │   ├── auth_service.go
│   │   ├── user_service.go
│   │   ├── progress_service.go
│   │   ├── stats_service.go
│   │   ├── achievement_service.go
│   │   └── content_service.go
│   ├── repository/              # Database access layer
│   │   ├── user_repo.go
│   │   ├── progress_repo.go
│   │   ├── achievement_repo.go
│   │   └── preference_repo.go
│   ├── db/
│   │   ├── db.go                # Database connection pool
│   │   └── migrations/          # SQL migration files
│   │       ├── 001_init_schema.sql
│   │       ├── 002_add_achievements.sql
│   │       └── ...
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── logger/
│   │   └── logger.go            # Logging utility
│   └── auth/
│       └── jwt.go               # JWT utilities
├── pkg/
│   ├── utils/
│   │   ├── validators.go
│   │   └── helpers.go
│   └── constants/
│       └── constants.go
├── migrations/                  # Database migration files
├── data/
│   ├── kana.json                # Static content
│   ├── kanji.json
│   └── vocabulary.json
├── tests/
│   ├── handlers_test.go
│   ├── services_test.go
│   └── integration_test.go
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml
```

### Frontend (React)

```
go-kanadojo-frontend/
├── src/
│   ├── app/
│   │   ├── App.tsx              # Root app component
│   │   ├── App.css
│   │   └── main.tsx             # Entry point
│   ├── pages/
│   │   ├── Home.tsx             # Home page
│   │   ├── KanaTraining.tsx
│   │   ├── KanjiTraining.tsx
│   │   ├── VocabularyTraining.tsx
│   │   ├── Progress.tsx
│   │   ├── Achievements.tsx
│   │   ├── Preferences.tsx
│   │   ├── Login.tsx
│   │   ├── Register.tsx
│   │   └── NotFound.tsx
│   ├── features/
│   │   ├── kana/
│   │   │   ├── components/
│   │   │   │   ├── KanaCards.tsx
│   │   │   │   ├── KanaGameModes.tsx
│   │   │   │   └── KanaMenu.tsx
│   │   │   ├── hooks/
│   │   │   │   └── useKanaGame.ts
│   │   │   ├── store/
│   │   │   │   └── kanaStore.ts
│   │   │   └── types.ts
│   │   ├── kanji/
│   │   ├── vocabulary/
│   │   ├── achievements/
│   │   ├── progress/
│   │   └── preferences/
│   ├── shared/
│   │   ├── components/
│   │   │   ├── Button.tsx
│   │   │   ├── Card.tsx
│   │   │   ├── Modal.tsx
│   │   │   ├── ThemeToggle.tsx
│   │   │   ├── LanguageSelector.tsx
│   │   │   └── ui/              # shadcn components
│   │   ├── hooks/
│   │   │   ├── useAudio.ts
│   │   │   ├── useFetch.ts
│   │   │   └── useLocalStorage.ts
│   │   ├── lib/
│   │   │   ├── api.ts           # API client
│   │   │   ├── utils.ts
│   │   │   └── constants.ts
│   │   ├── store/
│   │   │   ├── authStore.ts
│   │   │   ├── themeStore.ts
│   │   │   └── preferencesStore.ts
│   │   └── types/
│   │       └── index.ts         # Shared types
│   ├── assets/
│   │   ├── sounds/
│   │   │   ├── click.mp3
│   │   │   ├── correct.mp3
│   │   │   └── error.mp3
│   │   └── wallpapers/
│   ├── styles/
│   │   ├── globals.css
│   │   ├── tailwind.css
│   │   └── animations.css
│   └── __tests__/               # Test files
├── public/
│   ├── data/
│   │   ├── kana.json
│   │   ├── kanji.json
│   │   └── vocabulary.json
│   ├── sounds/
│   └── wallpapers/
├── index.html
├── package.json
├── tsconfig.json
├── vite.config.ts
├── tailwind.config.js
├── vitest.config.ts
└── .env.example
```

---

## 🔌 API Design

### Authentication Endpoints

```
POST   /api/v1/auth/register         Register new user
POST   /api/v1/auth/login            Login user (returns JWT)
POST   /api/v1/auth/refresh          Refresh access token
POST   /api/v1/auth/logout           Logout user
GET    /api/v1/auth/me               Get current user info
POST   /api/v1/auth/verify-email     Verify email address
POST   /api/v1/auth/forgot-password  Initiate password reset
```

### Content Endpoints

```
GET    /api/v1/kana                  Get all kana characters
GET    /api/v1/kana/:id              Get specific kana
GET    /api/v1/kana/subset/:subset   Get kana by subset
GET    /api/v1/kanji                 Get all kanji
GET    /api/v1/kanji/:id             Get specific kanji
GET    /api/v1/kanji/jlpt/:level     Get kanji by JLPT level
GET    /api/v1/vocabulary            Get all vocabulary
GET    /api/v1/vocabulary/:id        Get specific word
GET    /api/v1/vocabulary/jlpt/:level Get words by JLPT level
```

### Progress & Stats Endpoints

```
GET    /api/v1/stats                 Get user statistics
GET    /api/v1/stats/kana            Get kana-specific stats
GET    /api/v1/stats/kanji           Get kanji-specific stats
GET    /api/v1/stats/vocabulary      Get vocabulary-specific stats
POST   /api/v1/progress/update       Update user progress
POST   /api/v1/progress/answer       Submit answer (calculate score)
GET    /api/v1/progress/history      Get practice history
GET    /api/v1/streak                Get current streak
```

### User Preferences

```
GET    /api/v1/preferences           Get user preferences
PUT    /api/v1/preferences           Update preferences
GET    /api/v1/preferences/theme     Get theme settings
PUT    /api/v1/preferences/theme     Update theme settings
```

### Achievements

```
GET    /api/v1/achievements          Get all achievements
GET    /api/v1/achievements/unlocked Get user's unlocked achievements
POST   /api/v1/achievements/check    Check and unlock achievements
```

---

## 🗄️ Database Schema

### Core Tables

```sql
-- Users Table
CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  username VARCHAR(100) UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  email_verified BOOLEAN DEFAULT FALSE,
  last_login TIMESTAMP,
  INDEX idx_email (email),
  INDEX idx_username (username)
);

-- User Progress (answers per content)
CREATE TABLE progress (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  content_type VARCHAR(50) NOT NULL, -- 'kana', 'kanji', 'vocabulary'
  content_id VARCHAR(100) NOT NULL,
  correct_count INT DEFAULT 0,
  wrong_count INT DEFAULT 0,
  last_practiced TIMESTAMP,
  mastery_level INT DEFAULT 0, -- 0-100
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(user_id, content_type, content_id),
  INDEX idx_user_id (user_id),
  INDEX idx_content (content_type, content_id)
);

-- Statistics (aggregated)
CREATE TABLE statistics (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
  total_answers INT DEFAULT 0,
  correct_answers INT DEFAULT 0,
  accuracy FLOAT DEFAULT 0,
  current_streak INT DEFAULT 0,
  longest_streak INT DEFAULT 0,
  last_practice_date DATE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_user_id (user_id)
);

-- User Preferences
CREATE TABLE preferences (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
  theme_name VARCHAR(100) DEFAULT 'default',
  font_name VARCHAR(100) DEFAULT 'default',
  sound_enabled BOOLEAN DEFAULT TRUE,
  hotkeys_enabled BOOLEAN DEFAULT TRUE,
  dark_mode BOOLEAN DEFAULT FALSE,
  language VARCHAR(10) DEFAULT 'en',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_user_id (user_id)
);

-- Achievements
CREATE TABLE achievements (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL UNIQUE,
  description TEXT,
  icon_url VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User Achievements (unlocked)
CREATE TABLE user_achievements (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  achievement_id UUID NOT NULL REFERENCES achievements(id) ON DELETE CASCADE,
  unlocked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(user_id, achievement_id),
  INDEX idx_user_id (user_id),
  INDEX idx_achievement_id (achievement_id)
);

-- Practice Sessions
CREATE TABLE practice_sessions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  content_type VARCHAR(50) NOT NULL,
  game_mode VARCHAR(50) NOT NULL, -- 'pick', 'reverse_pick', 'input', 'reverse_input'
  started_at TIMESTAMP NOT NULL,
  ended_at TIMESTAMP,
  correct_answers INT DEFAULT 0,
  total_answers INT DEFAULT 0,
  duration_seconds INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_user_id (user_id),
  INDEX idx_started_at (started_at)
);

-- Answer History (detailed tracking)
CREATE TABLE answers (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  session_id UUID NOT NULL REFERENCES practice_sessions(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  content_id VARCHAR(100) NOT NULL,
  user_answer VARCHAR(255),
  correct_answer VARCHAR(255) NOT NULL,
  is_correct BOOLEAN NOT NULL,
  time_spent_ms INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_session_id (session_id),
  INDEX idx_user_id (user_id)
);
```

---

## 🎯 Key Features Implementation

### 1. Game Mode Logic

**4 Game Modes per Content Type:**

1. **Pick**: Show character, user selects from 4 choices
   - Backend: Generates random distractors
   - Frontend: Multiple choice UI

2. **Reverse-Pick**: Show romaji/translation, user selects character
   - Backend: Same as Pick but shows answer side
   - Frontend: Same UI, different data mapping

3. **Input**: Show character, user types response
   - Backend: Validates user input against correct answer
   - Frontend: Text input with validation feedback

4. **Reverse-Input**: Show romaji/translation, user types character
   - Backend: Same validation as Input
   - Frontend: Same input UI

### 2. Progress Tracking

- **Per-Content Tracking**: Track correct/wrong for each character
- **Session Tracking**: Log practice sessions with duration
- **Answer History**: Keep detailed answer logs for analytics
- **Mastery Calculation**: Percentage based on recent performance
- **Streak System**: Current and longest streaks

### 3. Achievement System

Example achievements:
- First steps (complete 10 answers)
- Streak master (100+ streak)
- Accuracy champion (95%+ accuracy)
- Speed demon (complete session in < 2 mins)
- Master of kana (100% accuracy on all kana)
- JLPT N5 warrior (complete all N5 kanji)
- Vocabulary builder (1000+ vocabulary items)

Trigger when:
- Specific milestones reached
- Accuracy thresholds exceeded
- Time-based challenges completed
- Content mastery achieved

### 4. Preferences & Customization

**Themes**: Store theme data in frontend (JSON), user preference in database
**Fonts**: Similar structure with dropdown selection
**Audio**: Toggle in preferences table
**Language**: Multi-language UI with i18n library
**Display Options**: Romaji/Kana toggle, English translation toggle

### 5. Statistics Dashboard

**Real-time Stats:**
- Accuracy percentage
- Total answers
- Streak information
- Most practiced content
- Recent activity
- Weekly/monthly progress charts

---

## 🔐 Authentication & Security

### JWT Authentication Flow

```
1. User registers → Create user in DB → Return JWT tokens
2. User logs in → Validate credentials → Return JWT tokens
3. Client stores tokens in secure storage
4. Client sends Authorization: Bearer {token} header
5. Server validates token → Process request
6. Token expires → Use refresh token to get new access token
```

### Implementation

**Backend (Go):**
```go
// Create JWT
func CreateToken(user *User) (accessToken, refreshToken string, err error)

// Verify JWT
func VerifyToken(tokenString string) (*Claims, error)

// Middleware
func AuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    token := c.GetHeader("Authorization")
    claims, err := VerifyToken(token)
    if err != nil {
      c.JSON(401, gin.H{"error": "unauthorized"})
      c.Abort()
      return
    }
    c.Set("user_id", claims.UserID)
    c.Next()
  }
}
```

**Frontend (React):**
```typescript
// Store tokens in secure storage (httpOnly cookies or localStorage)
// Add to every API request
function apiCall(url: string, options = {}) {
  const token = localStorage.getItem('accessToken');
  return fetch(url, {
    ...options,
    headers: {
      ...options.headers,
      'Authorization': `Bearer ${token}`
    }
  });
}
```

---

## 🚀 Development Workflow

### Phase 1: Project Setup (Week 1)

1. Initialize Go backend with Gin framework
2. Set up PostgreSQL database
3. Create database migrations
4. Initialize React project with Vite
5. Set up Tailwind CSS and shadcn/ui
6. Configure TypeScript, ESLint, Prettier

### Phase 2: Backend Foundation (Weeks 2-3)

1. Database models and migrations
2. Authentication system (register, login, JWT)
3. User service layer
4. Content endpoints (kana, kanji, vocabulary)
5. Middleware (CORS, logging, auth)
6. Input validation and error handling

### Phase 3: Frontend Foundation (Weeks 2-3)

1. Component library setup with shadcn/ui
2. Routing with React Router
3. State management with Zustand
4. API client wrapper
5. Authentication flow (login, register, token refresh)
6. Theme and preference system

### Phase 4: Game Logic (Weeks 4-5)

1. Content selection UI
2. Game mode implementations (all 4 modes)
3. Question generation logic
4. Answer validation
5. Real-time feedback (visual + audio)
6. Progress update API integration

### Phase 5: Features (Weeks 6-7)

1. Progress tracking UI
2. Statistics dashboard
3. Achievement system
4. Preferences/settings panel
5. Customization (themes, fonts, sounds)
6. Multi-language support

### Phase 6: Polish & Testing (Weeks 8+)

1. Unit tests (both backend and frontend)
2. Integration tests
3. Performance optimization
4. Accessibility audit
5. Responsive design verification
6. Security audit

---

## 📊 Technology Decisions & Rationale

| Decision                | Rationale                                          |
| ----------------------- | -------------------------------------------------- |
| **Go** for backend      | Fast, concurrent, excellent for I/O-heavy APIs    |
| **React** for frontend  | Rich ecosystem, component reusability, large team  |
| **PostgreSQL**          | Mature, reliable RDBMS with strong JSON support   |
| **JWT** for auth        | Stateless, scalable, industry standard             |
| **REST API**            | Simpler than GraphQL for this use case             |
| **Tailwind CSS**        | Utility-first, rapid development, minimal bundle  |
| **Zustand**             | Lightweight state management, no boilerplate       |
| **Vite**                | Fast build tool, excellent DX                     |

---

## 🔄 Data Flow Example: Answering a Question

### User Flow
```
1. User starts Kana Pick mode
2. Frontend requests random kana from backend
3. Backend selects random kana, generates 3 wrong answers
4. Backend returns: { correct: 'あ', choices: ['あ', 'い', 'う', 'え'] }
5. Frontend displays multiple choice UI
6. User selects answer
7. Frontend sends: { answer: 'あ', content_id: 'kana-a' } to backend
8. Backend validates and calculates score
9. Backend returns: { correct: true, streak: 5, points: 10 }
10. Frontend plays sound, updates UI, stores progress locally
11. Frontend updates stats store with new data
12. User continues or ends session
```

---

## 📈 Scaling Considerations

### Caching Strategy
- **Frontend**: Cache content lists (kana, kanji, vocab) with periodic sync
- **Backend**: Cache frequently accessed content, user preferences
- **Database**: Indexes on commonly queried columns

### Performance Optimizations
- Lazy load large content sets
- Paginate progress history
- Compress audio files (MP3/OGG)
- CDN for static assets
- Database connection pooling

### Future Scaling
- Add Redis for session caching
- Implement GraphQL for complex queries
- WebSocket for real-time leaderboards
- Consider microservices architecture

---

## 🧪 Testing Strategy

### Backend Testing
```go
// Unit tests for services
func TestCreateUser(t *testing.T) {
  // Test user creation, validation, etc.
}

// Integration tests for handlers
func TestPostRegister(t *testing.T) {
  // Test full HTTP request/response
}
```

### Frontend Testing
```typescript
// Component tests with React Testing Library
test('renders kana card with character', () => {
  render(<KanaCard character="あ" />);
  expect(screen.getByText('あ')).toBeInTheDocument();
});

// Integration tests for game flow
test('completes practice session', async () => {
  // Test full game flow from start to finish
});
```

---

## 📋 Checklist for Implementation

### Backend
- [ ] Database schema created and migrated
- [ ] User authentication working (register, login, logout)
- [ ] Content endpoints returning correct data
- [ ] Progress tracking implemented
- [ ] Statistics calculation working
- [ ] Achievement system functional
- [ ] Error handling consistent
- [ ] Logging in place
- [ ] Tests written and passing
- [ ] API documentation complete

### Frontend
- [ ] React Router configured
- [ ] Authentication flow complete
- [ ] All 4 game modes implemented
- [ ] Theme system working
- [ ] Sound effects functional
- [ ] Progress tracking UI displaying data
- [ ] Responsive design verified
- [ ] Accessibility audit passed
- [ ] Tests written and passing
- [ ] Production build optimized

---

## 🎓 Resources to Learn

### Go Backend
- **Gin Framework**: https://gin-gonic.com/
- **GORM**: https://gorm.io/ (ORM alternative)
- **Database Migrations**: https://github.com/golang-migrate/migrate

### React Frontend
- **React Router**: https://reactrouter.com/
- **Zustand**: https://github.com/pmndrs/zustand
- **Shadcn/ui**: https://ui.shadcn.com/
- **Vite**: https://vitejs.dev/

### Full Stack
- **API Design Best Practices**: https://restfulapi.net/
- **JWT**: https://jwt.io/
- **PostgreSQL**: https://www.postgresql.org/docs/

---

## 📞 Next Steps

1. **Choose your Go framework** (Gin, Echo, or Chi)
2. **Set up PostgreSQL** locally or cloud
3. **Initialize project structures** for both backend and frontend
4. **Start with authentication** - this is foundational
5. **Build content endpoints** and hook up frontend
6. **Implement game logic** for one mode, then replicate to others
7. **Add progress tracking**
8. **Build the UI** feature by feature

Good luck building your learning platform! 🚀
