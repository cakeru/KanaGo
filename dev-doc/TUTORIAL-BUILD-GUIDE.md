# 🎬 KanaDojo Clone - Complete Build Tutorial
## Building a Japanese Learning App with Go + React

**By:** Your Development Lecturer  
**Level:** Intermediate (assumes basic programming knowledge)  
**Duration:** ~30 days of active development  
**Goal:** Build a production-ready Japanese learning platform  

---

## 📺 Welcome to the Course!

Hey everyone! 👋 Welcome to the complete build guide for KanaDojo Clone. I'm your instructor, and in this course, we're going to build a **full-stack Japanese learning application** from scratch.

Think of this guide like a lecture series. I'm going to walk you through every step, explain the "why" behind our decisions, and guide you through the building process without writing code for you. You'll learn, understand, and build this yourself.

### What You'll Learn

By the end of this course, you'll understand:
- ✅ How to build a REST API with Go
- ✅ How to build a React SPA with TypeScript
- ✅ How to design a scalable database
- ✅ How authentication works (JWT tokens)
- ✅ How to connect frontend and backend
- ✅ How to track user progress and data
- ✅ How to deploy to production

### Prerequisites

Before we start, make sure you have:
- **Basic Go knowledge** - You should understand goroutines, packages, interfaces
- **React/JavaScript knowledge** - Comfortable with hooks, components, state
- **SQL basics** - Can write simple SELECT, INSERT, UPDATE queries
- **Command line** - Comfortable with terminal/bash
- **Git** - Basic version control knowledge

If you're weak on any of these, I recommend brushing up first. This won't be easy, but it'll be rewarding!

---

## 🏗️ What We're Building

Let me show you the big picture before we dive into the code.

### The Application (30-Second Pitch)

**KanaDojo** is a Japanese learning platform where users:
1. Log in and create an account
2. Choose what to study (Hiragana, Katakana, Kanji, Vocabulary)
3. Play games to learn (4 different game modes)
4. Track their progress
5. Earn achievements and streaks
6. Customize their learning experience

### The Tech Stack

```
┌──────────────────────────────────────┐
│  FRONTEND: React + TypeScript        │
│  What user sees in the browser       │
└──────────────────┬───────────────────┘
                   │ HTTP Requests/Responses
┌──────────────────▼───────────────────┐
│  BACKEND: Go (Gin Framework)         │
│  API logic, authentication, data     │
└──────────────────┬───────────────────┘
                   │ SQL Queries
┌──────────────────▼───────────────────┐
│  DATABASE: PostgreSQL                │
│  Stores all user data persistently   │
└──────────────────────────────────────┘
```

### Why These Technologies?

**Go for Backend:**
- ⚡ Super fast (compiled language)
- 📦 Simple syntax, easy to learn
- 🔄 Great for handling lots of concurrent users
- 🎯 Perfect for REST APIs

**React for Frontend:**
- 🎨 Beautiful, interactive UIs
- ⚡ Fast and responsive
- 🔧 Huge ecosystem of libraries
- 📱 Can build mobile apps later with same knowledge

**PostgreSQL for Database:**
- 🔐 Rock-solid reliability
- 🧮 Handles complex queries
- 🚀 Can scale to millions of users
- 📊 Great for analytics

---

## 📅 The 30-Day Timeline (7 Phases)

Here's how we're going to organize the next 30 days:

### **WEEK 1: Foundation (Days 1-7)**

#### Phase 1: Setup (Days 1-2)
- **What:** Get everything installed and ready
- **Why:** You can't build a house without tools
- **Topics:**
  - Install Go, Node.js, PostgreSQL
  - Create project folders
  - Initialize Git repositories
  - Set up IDE and tools

#### Phase 2: Backend Foundation (Days 3-5)
- **What:** Create the Go server structure
- **Why:** The backend needs a solid foundation before we add features
- **Topics:**
  - Initialize Go project
  - Set up the web framework (Gin)
  - Create configuration system
  - Connect to PostgreSQL database
  - Run your first "Hello World" API

#### Phase 3: Authentication (Days 6-7)
- **What:** Build user login/signup
- **Why:** Users need to create accounts and log in securely
- **Topics:**
  - Create User table in database
  - Build JWT authentication system
  - Implement password hashing
  - Create register and login endpoints
  - Test with Postman

### **WEEK 2: Content & Frontend (Days 8-14)**

#### Phase 4: Content API (Days 8-9)
- **What:** Create endpoints to serve learning content
- **Why:** The game needs data to display (kana, kanji, vocabulary)
- **Topics:**
  - Load static JSON data (kana.json, kanji.json)
  - Create GET endpoints for content
  - Implement filtering and pagination
  - Test all endpoints

#### Phase 5: Frontend Setup (Days 10-12)
- **What:** Create React app with routing and state
- **Why:** We need a place to display the UI
- **Topics:**
  - Initialize React with Vite
  - Set up TypeScript
  - Configure Zustand for state management
  - Set up React Router
  - Build navigation structure
  - Connect to backend authentication

#### Phase 6: Game Logic (Days 13-14)
- **What:** Build the actual game
- **Why:** The learning happens in the game
- **Topics:**
  - Design the 4 game modes (Pick, Reverse-Pick, Input, Reverse-Input)
  - Create game UI components
  - Handle user answers
  - Send answers to backend
  - Show feedback (correct/incorrect)
  - Track progress in database

### **WEEK 3: Polish & Features (Days 15-21)**

#### Phase 7: Features & Optimization (Days 15-30)
- **What:** Add nice-to-have features and optimize
- **Why:** Every good product needs polish
- **Topics:**
  - Statistics dashboard (accuracy, streaks)
  - Achievements system
  - User preferences (theme, font, language)
  - Sound effects and animations
  - Testing and bug fixes
  - Performance optimization
  - Deployment preparation

---

## 🎓 Core Concepts You Need to Understand

Before you start coding, let me explain some key concepts. These are the "laws of the land" in your app.

### 1. **Client-Server Architecture**

```
CLIENT (Your Browser)          SERVER (Go)              DATABASE
┌──────────────────┐          ┌──────────────┐        ┌──────────┐
│                  │          │              │        │          │
│ 1. User clicks   │──HTTP──→ │ 2. Server    │──SQL→  │ 3. Store │
│    "Login"       │          │    processes │        │   data   │
│                  │←──JSON─── │ 4. Responds  │←──Row── │          │
│ 5. Shows result  │          │              │        │          │
│                  │          │              │        │          │
└──────────────────┘          └──────────────┘        └──────────┘
```

**Key Point:** Frontend doesn't talk directly to database. It sends HTTP requests to the server, which decides what to do.

### 2. **Authentication with JWT (JSON Web Tokens)**

Imagine JWT like a ticket system:

```
USER:                          SERVER:
1. Enter username/password  →   Verify credentials
2. Server creates JWT       ←   Send back a token
   (like a ticket)              
3. Store token in browser   
4. Use token for every   →     Verify token
   future request              Is it valid? Is it expired?
```

The token is like a signed ticket that proves you're logged in without sending your password every time.

### 3. **The 4 Game Modes**

Our app has 4 different ways to learn:

**Mode 1: Pick (Multiple Choice)**
```
Display: あ
Options: [a] [i] [u] [e] [o]
User clicks: [a] ✓ Correct!
```

**Mode 2: Reverse-Pick (Reverse Multiple Choice)**
```
Display: "a"
Options: [あ] [い] [う] [え] [お]
User clicks: [あ] ✓ Correct!
```

**Mode 3: Input (Type the Answer)**
```
Display: あ
User types: a ✓ Correct!
```

**Mode 4: Reverse-Input (Type the Character)**
```
Display: "a"
User types: あ ✓ Correct!
```

All 4 modes help learning from different angles.

### 4. **Progress Tracking**

Every time user answers a question, we:
1. Check if answer is correct
2. Update their progress (how many correct/incorrect)
3. Calculate "mastery level" (0-100%)
4. Update statistics (streak, accuracy)
5. Check if they unlocked any achievements

```
Answer submitted
    ↓
Is it correct?
    ├─ Yes → Streak +1, Accuracy increases
    └─ No → Streak reset to 0, Try again
    ↓
Update progress level for this content
    ↓
Check achievements (Did they reach 100 correct? Award achievement!)
    ↓
Save everything to database
```

### 5. **State Management on Frontend**

On the frontend, we use **Zustand** to manage state (think of it as a place to store data that multiple components need):

```
Store:
├── authStore
│   ├── currentUser
│   ├── accessToken
│   └── isLoggedIn
├── gameStore
│   ├── currentQuestion
│   ├── userAnswer
│   └── score
└── preferencesStore
    ├── theme
    ├── fontSize
    └── language
```

Different components can read from and write to this store.

---

## 🛠️ Development Setup Walkthrough

### STEP 1: Install Required Tools

#### 1.1 Install Go
- Go to **golang.org**
- Download Go 1.21 or later
- Follow installation instructions for your OS
- Verify: Open terminal, type `go version`, you should see the version number

#### 1.2 Install Node.js
- Go to **nodejs.org**
- Download Node.js 18+ (LTS version)
- Verify: Type `node --version` and `npm --version` in terminal

#### 1.3 Install PostgreSQL
- Go to **postgresql.org**
- Download PostgreSQL 14 or later
- During installation, remember the password you set for "postgres" user
- Verify: Type `psql --version` in terminal

#### 1.4 Install Git
- Go to **git-scm.com**
- Download and install
- Verify: Type `git --version`

#### 1.5 Install IDE
- Download **VS Code** from code.visualstudio.com
- Install these extensions:
  - "Go" (by Go Team)
  - "ES7+ React/Redux/React-Native snippets" (by dsznajder.es7-react-js-snippets)
  - "Thunder Client" (for testing API)

### STEP 2: Create Project Structure

In your terminal:

```bash
# Create main project folder
mkdir kanago-clone
cd kanago-clone

# Create backend folder
mkdir backend
cd backend

# Create frontend folder (from kanago-clone)
cd ..
mkdir frontend
```

So your structure looks like:
```
kanago-clone/
├── backend/    (Go code goes here)
└── frontend/   (React code goes here)
```

### STEP 3: Initialize Git Repositories

```bash
# In backend folder
cd backend
git init
git config user.name "Your Name"
git config user.email "your@email.com"

# Create .gitignore file (I'll guide you through this in the Backend Setup section)

# In frontend folder
cd ../frontend
git init
```

---

## 🔧 Backend Development Workflow

### Phase 1-2: Setting Up the Go Server

#### What You'll Do:
1. **Initialize Go Project**
   - Create `go.mod` and `go.sum` files
   - These track your project dependencies

2. **Install Dependencies**
   - Gin (web framework)
   - sqlx (database driver)
   - JWT library (authentication)
   - dotenv (configuration)

3. **Create Basic Structure**
   ```
   backend/
   ├── cmd/server/main.go         ← Entry point
   ├── internal/
   │   ├── api/handlers/          ← Handle HTTP requests
   │   ├── services/              ← Business logic
   │   ├── repository/            ← Database queries
   │   ├── models/                ← Data structures
   │   ├── config/                ← Settings
   │   └── middleware/            ← Auth middleware
   ├── migrations/                ← Database setup scripts
   ├── go.mod
   └── go.sum
   ```

4. **Create Your First Endpoint**
   - `/health` - Returns "OK" to test server is running
   - This verifies Gin is working

5. **Set Up Configuration**
   - Create `.env` file with settings:
     - Database URL
     - JWT Secret
     - Server Port
   - Load these into your app

6. **Connect to PostgreSQL**
   - Create database connection
   - Test that connection works
   - You should be able to query the database

#### Key Questions to Answer:
- ❓ Why use Gin instead of writing HTTP handling from scratch?
  - **Answer:** Gin handles routing, middleware, and common patterns for us
  
- ❓ What's the difference between sqlx and other database libraries?
  - **Answer:** sqlx gives us simple, fast SQL queries without complex ORMs

- ❓ Why put code in `internal/` folder?
  - **Answer:** Go respects this convention - code in `internal/` can't be imported by other projects

---

### Phase 3: Authentication System

#### What You'll Do:

1. **Design Database Schema**
   - Create `users` table with:
     - id (unique identifier)
     - email (for login)
     - username (display name)
     - password_hash (encrypted password, never store plain password!)
     - created_at, updated_at (timestamps)

2. **Create User Model**
   - Go struct representing a user
   - This is the "shape" of user data

3. **Build Authentication Logic**
   - **Password Hashing:** Use bcrypt to hash passwords securely
   - **JWT Generation:** Create tokens that prove user is logged in
   - **Token Verification:** Verify tokens are valid and not expired

4. **Create Auth Endpoints**
   - `POST /auth/register` - New user signs up
   - `POST /auth/login` - Existing user logs in
   - `POST /auth/refresh` - Get new token when old one expires
   - `GET /auth/me` - Get current user info

5. **Create Auth Middleware**
   - Middleware is code that runs before your endpoint handler
   - It checks: Does request have a valid token?
   - If yes, let request proceed
   - If no, return "Unauthorized" error

#### The Auth Flow (Detailed):

```
USER SIGNUP:
1. Frontend: User fills form (email, password) → Sends to /auth/register
2. Backend:
   - Validate email format (is it valid?)
   - Check if email already exists
   - Hash password using bcrypt (one-way encryption)
   - Save user to database
   - Create JWT token (proof of login)
3. Response: Send back user data + access token + refresh token
4. Frontend: Store tokens in browser (localStorage or cookies)

USER LOGIN:
1. Frontend: User enters email/password → Sends to /auth/login
2. Backend:
   - Find user by email
   - Use bcrypt to compare password hash (can't decrypt)
   - If match, create JWT token
   - If no match, return error
3. Response: Send back user data + tokens
4. Frontend: Store tokens

SUBSEQUENT REQUESTS:
1. Frontend: Every API call includes token in header:
   Authorization: Bearer <token>
2. Backend Middleware: Extracts token, verifies it's valid
3. Proceeds only if token is valid
4. Handler knows which user made the request (from token)
```

#### Key Decisions:
- ❓ Why hash passwords?
  - **Answer:** If database is compromised, attackers can't use password to log into user's other accounts
  
- ❓ What's a refresh token?
  - **Answer:** Access tokens expire in 15 minutes. Refresh token lasts 7 days and is used to get new access tokens without re-login.

- ❓ Why JWT instead of sessions?
  - **Answer:** JWT is stateless - server doesn't need to look up session data. Scales better.

---

### Phase 4: Content API

#### What You'll Do:

1. **Load Static Content**
   - Put JSON files in your backend: `kana.json`, `kanji.json`, `vocabulary.json`
   - Each file contains arrays of learning items
   - Example kana.json:
     ```json
     [
       {"id": 1, "character": "あ", "romanji": "a", "category": "hiragana"},
       {"id": 2, "character": "い", "romanji": "i", "category": "hiragana"}
     ]
     ```

2. **Create Handlers for Content**
   - `GET /kana` - Get all kana
   - `GET /kana/:id` - Get specific kana
   - `GET /kanji` - Get all kanji
   - Add filtering: `GET /kanji?level=N1` (only N1 level kanji)
   - Add pagination: `GET /kanji?page=1&limit=20`

3. **Add Filtering & Pagination**
   - **Why filtering?** Users want to study specific things
   - **Why pagination?** Don't send 10,000 items at once (slow!)
   - Send 20 at a time, user navigates with "next page"

4. **Response Formatting**
   - Every response should have consistent format:
     ```json
     {
       "success": true,
       "data": [...],
       "error": null,
       "total": 2370
     }
     ```

#### Key Concepts:
- ❓ Why separate static content from user data?
  - **Answer:** Learning content doesn't change. User progress data does. Different handling.

- ❓ What's pagination?
  - **Answer:** Instead of getting all 2,370 kanji, get 20 at a time. Like pages in a book.

---

## 🎨 Frontend Development Workflow

### Phase 5: Setting Up React

#### What You'll Do:

1. **Initialize React Project with Vite**
   - Vite is the new, faster way to build React apps
   - Sets up project structure automatically
   - Much faster than Create React App

2. **Install Dependencies**
   - React Router (page navigation)
   - Zustand (state management)
   - Axios or Fetch (HTTP requests)
   - Tailwind CSS (styling)
   - shadcn/ui (pre-built components)

3. **Project Structure**
   ```
   frontend/
   ├── src/
   │   ├── app/
   │   │   ├── App.tsx        ← Root component
   │   │   └── main.tsx       ← Entry point
   │   ├── pages/
   │   │   ├── Home.tsx
   │   │   ├── Login.tsx
   │   │   ├── GamePage.tsx
   │   │   └── ProgressPage.tsx
   │   ├── features/
   │   │   ├── kana/
   │   │   ├── kanji/
   │   │   └── achievements/
   │   ├── shared/
   │   │   ├── components/    ← Reusable UI components
   │   │   ├── hooks/         ← Custom logic
   │   │   ├── store/         ← State management
   │   │   ├── types/         ← TypeScript types
   │   │   └── lib/api.ts     ← API communication
   ├── public/
   │   └── data/              ← JSON content files
   └── package.json
   ```

4. **Set Up Routing**
   - `/` - Home page
   - `/login` - Login page
   - `/register` - Signup page
   - `/kana` - Kana game
   - `/progress` - Progress dashboard
   - `/settings` - User preferences

5. **Create API Client**
   - Function that makes HTTP requests to backend
   - Automatically adds JWT token to requests
   - Handles common errors

#### Key Decisions:
- ❓ What's the difference between pages and components?
  - **Answer:** Pages are full screens (Login, Game, Progress). Components are reusable pieces (Button, Card, Modal).

- ❓ Why Zustand instead of Redux?
  - **Answer:** Redux is powerful but heavy. Zustand is simpler, lighter, perfect for this size app.

- ❓ What's TypeScript?
  - **Answer:** Adds type checking to JavaScript. Catches bugs before they happen.

---

### Phase 6: Building the Game

#### What You'll Do:

1. **Design Game Component Architecture**
   ```
   GamePage
   ├── GameModeSelector (Pick, Reverse-Pick, Input, Reverse-Input)
   ├── ContentSelector (Choose Kana, Kanji, or Vocabulary)
   └── GameArea
       ├── QuestionDisplay
       ├── AnswerInput/Options
       ├── SubmitButton
       ├── Feedback (Correct/Incorrect)
       └── ScoreDisplay
   ```

2. **Create Custom Game Hook (useGame)**
   - This handles all game logic:
     - Get random question from content
     - Check if answer is correct
     - Update score
     - Get next question
     - Track time spent

3. **Implement Game Loop**
   ```
   1. Display question (random from content)
   2. User provides answer (click/type)
   3. Send to backend to verify + save
   4. Backend returns: correct/incorrect
   5. Show feedback animation
   6. If correct, celebrate + next question
   7. If incorrect, show right answer + next question
   8. Repeat until user finishes session
   ```

4. **Connect to Backend**
   - API call: `POST /progress/update` with answer data
   - Backend checks answer
   - Saves progress to database
   - Returns updated user statistics

5. **Add Visual Feedback**
   - Correct answer: Green highlight, checkmark, celebration animation
   - Incorrect answer: Red highlight, X, show correct answer
   - Sound effects: Click, correct, incorrect

#### Key Flow:

```
USER INTERACTION → GAME LOGIC → API CALL → BACKEND → DATABASE → RESPONSE → UPDATE UI

1. User sees: "あ" with options [a] [i] [u] [e] [o]
2. User clicks: [a]
3. Game logic checks: Is [a] correct for あ?
4. Sends to backend: {question_id: 1, user_answer: "a"}
5. Backend checks database, verifies answer
6. Returns: {is_correct: true, mastery_level: 45}
7. Frontend updates: Shows ✓ Correct!, increases score
8. Loads next question
```

#### Key Concepts:
- ❓ Why custom hooks?
  - **Answer:** Keep game logic separate from UI. Easy to test and reuse.

- ❓ What's the difference between Zustand and component state?
  - **Answer:** Component state (useState) for single component. Zustand for data that multiple components need.

---

### Phase 7: Dashboard & Features

#### What You'll Do:

1. **Progress Dashboard**
   - Show mastery level for each content type (Kana, Kanji, Vocab)
   - Display as progress bars (0-100%)
   - Show accuracy percentage
   - Show current streak (consecutive correct answers)

2. **Statistics Page**
   - Total questions answered
   - Overall accuracy
   - Average time per question
   - Streak history
   - Charts/graphs of progress over time

3. **Achievements System**
   - Unlock achievements for milestones:
     - "First Steps" - Answer 10 questions
     - "Kana Master" - 100% mastery in hiragana
     - "Century Club" - 100 correct answers in one session
     - "Streak King" - 50 consecutive correct answers
   - Show locked/unlocked achievements with progress

4. **Preferences Panel**
   - **Theme:** Dark mode, light mode
   - **Font:** Multiple Japanese font options
   - **Language:** English, Spanish, Japanese
   - **Difficulty:** Easy, Medium, Hard
   - **Sound:** On/Off

5. **Polish & Optimizations**
   - Loading spinners while data loads
   - Error messages if something goes wrong
   - Mobile responsive design
   - Smooth animations and transitions

#### Database Updates:

For all these features, we need database tables:
- `progress` - Track mastery level per content
- `statistics` - Overall stats for user
- `achievements` - Achievement definitions
- `user_achievements` - Which user unlocked which achievements
- `preferences` - User settings
- `practice_sessions` - Each time user plays, save session
- `answers` - Every answer submitted

---

## 💾 Database Design Deep Dive

Let me explain the database structure:

### Core Tables:

1. **users**
   ```
   id (auto-increment)
   email (unique)
   username
   password_hash
   created_at
   updated_at
   ```

2. **progress**
   ```
   id
   user_id → links to users
   content_id (kanji, kana, vocab id)
   content_type (kana/kanji/vocabulary)
   correct_count
   incorrect_count
   mastery_level (0-100)
   last_reviewed_at
   ```
   
   **Why this structure?** When user plays kana game, we need to track their progress on each kana character. This table stores that.

3. **statistics**
   ```
   id
   user_id → links to users
   total_answers
   correct_answers
   current_streak
   longest_streak
   accuracy_percentage
   updated_at
   ```
   
   **Why separate?** For quick dashboard queries. Could calculate from answers table, but doing it now is faster.

4. **practice_sessions**
   ```
   id
   user_id → links to users
   game_mode (pick/reverse-pick/input/reverse-input)
   content_type
   duration_seconds
   correct_count
   incorrect_count
   created_at
   ```
   
   **Why?** Track each study session. Help users see patterns. "I studied kana for 45 minutes yesterday."

5. **answers**
   ```
   id
   session_id → links to practice_sessions
   question_id
   user_answer
   is_correct
   time_spent_ms
   created_at
   ```
   
   **Why?** Keep record of every answer. Useful for analytics: "Which kanji do users struggle with?"

6. **achievements**
   ```
   id
   name ("Kana Master")
   description
   icon_url
   points (for gamification)
   ```
   
   **Why?** Achievement definitions. These don't change.

7. **user_achievements**
   ```
   id
   user_id → links to users
   achievement_id → links to achievements
   unlocked_at
   ```
   
   **Why?** Track which user unlocked which achievement. Many-to-many relationship.

8. **preferences**
   ```
   id
   user_id → links to users
   theme (dark/light)
   font (gothic/mincho/modern)
   language (en/es/ja)
   sound_enabled (true/false)
   updated_at
   ```

### Relationships:
- **One user** → **Many progress entries** (user studies many content pieces)
- **One user** → **Many practice sessions** (user plays many times)
- **One user** → **Many achievements** (user unlocks many achievements)
- **One session** → **Many answers** (each session has many answers)

Think of it like:
- **users** table is the "root"
- Everything else branches off from there
- `user_id` is the connection

---

## 🔌 Frontend-Backend Integration

This is where the magic happens!

### How They Communicate:

1. **HTTP Requests (Frontend → Backend)**
   ```
   POST /auth/login
   Content-Type: application/json
   
   {
     "email": "user@example.com",
     "password": "securepassword123"
   }
   ```

2. **HTTP Responses (Backend → Frontend)**
   ```
   200 OK
   Content-Type: application/json
   
   {
     "success": true,
     "data": {
       "user": {
         "id": 1,
         "email": "user@example.com",
         "username": "john_doe"
       },
       "access_token": "eyJhbGciOiJIUzI1NiIs...",
       "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
     }
   }
   ```

3. **Protected Requests** (with JWT token)
   ```
   GET /progress
   Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
   ```

### API Contract (The Agreement):

The **api-specification.md** document is the agreement between frontend and backend developers:

- Frontend team: "We'll send requests in this format"
- Backend team: "We'll respond in this format"
- Both agree on what endpoint URLs exist
- Both agree on error codes

This way, frontend can build UI, backend can build API, they work independently, then combine!

### Error Handling:

```
Frontend                    Backend                   Result
─────────────────          ─────────────────         ─────────
Send request    →          Process request
                           Check for errors
                           ├─ Bad input?     ← 400 Bad Request
                           ├─ Not logged in? ← 401 Unauthorized
                           ├─ No permission? ← 403 Forbidden
                           ├─ Not found?     ← 404 Not Found
                           ├─ Server error?  ← 500 Server Error
                           └─ Success?       ← 200 OK
                           Return response   →     Show to user
```

Frontend shows user-friendly messages:
- "Please enter a valid email" (from 400 error)
- "You need to log in" (from 401 error)
- "Something went wrong, try again" (from 500 error)

---

## 🚀 Deployment Strategy

### Development Environment (Your Computer)
```
Running locally on ports:
- Frontend: http://localhost:5173 (Vite)
- Backend: http://localhost:8080 (Go)
- Database: postgres://localhost:5432 (PostgreSQL)

You test everything here before deploying.
```

### Production Environment (Live Server)
```
Running on actual server with domains:
- Frontend: https://kanadojo.com (Static hosting or server)
- Backend: https://api.kanadojo.com (Cloud server)
- Database: PostgreSQL on managed service

Users access the real app here.
```

### Deployment Process:

1. **Backend Deployment**
   - Code → Server
   - Database migrations run
   - Server starts accepting requests
   - Endpoints are live

2. **Frontend Deployment**
   - React code → Built into HTML/CSS/JS
   - Uploaded to static hosting
   - Points to production backend API
   - Users download app

3. **Continuous Integration/Deployment**
   - Push code to GitHub
   - Tests run automatically
   - If tests pass, deploy automatically
   - If tests fail, stop deployment

### Key Platforms:

**Backend hosting options:**
- Railway.app (great for Go)
- Heroku (legacy but simple)
- AWS/Google Cloud (powerful but complex)

**Frontend hosting options:**
- Vercel (built for React)
- Netlify (simple deployment)
- AWS S3 + CloudFront (scalable)

**Database hosting:**
- Managed PostgreSQL services (Supabase, AWS RDS, Railway)
- Don't self-host in production (hassle)

---

## 📋 Testing Strategy

### What to Test:

1. **Backend Testing**
   - Does login work? (Try with correct/wrong password)
   - Does game endpoint return right question? (Try multiple times, should be random)
   - Does answer submission save correctly? (Check database)
   - Do protected endpoints reject requests without token?

2. **Frontend Testing**
   - Does form validation work? (Try empty fields)
   - Does game display questions correctly?
   - Do sounds play when they should?
   - Does dark mode toggle work?

3. **Integration Testing**
   - Register new user → Login → Play game → Check progress updated
   - Full user journey

### How to Test:

**Manual Testing:**
- Use Postman to test API endpoints
- Use browser dev tools to check console for errors
- Click buttons and use app normally

**Automated Testing:**
- Backend: Go testing framework
- Frontend: Jest + React Testing Library

### Testing Tools:

- **Postman** - Test API endpoints (visual)
- **Thunder Client** - VS Code extension for API testing
- **Browser DevTools** - Check for JavaScript errors
- **Go test framework** - Write automated tests for backend
- **Jest** - Write automated tests for frontend

---

## 🎯 Daily Development Workflow

Here's what a typical development day looks like:

### Morning (Planning)
1. Pick today's task from checklist
2. Read relevant documentation
3. Plan what you'll build
4. Create a branch in Git: `git checkout -b feature/user-auth`

### Midday (Coding)
1. Write code for the feature
2. Test as you go
3. Keep making commits: `git commit -am "Add password hashing"`
4. Takes 2-4 hours for a feature

### Afternoon (Testing)
1. Manual testing with Postman/browser
2. Fix bugs that come up
3. Check database to verify data saved
4. Write simple tests

### Evening (Wrapping Up)
1. Clean up code (remove unused imports, fix formatting)
2. Write clear commit message
3. Push to GitHub: `git push origin feature/user-auth`
4. Create Pull Request (code review)

### Repeat Next Day
- Build next feature
- Integration with previous features
- Continue progressing through phases

---

## 🎓 Learning Path

If you feel weak on any concepts:

### Go Fundamentals:
- Watch: "Go for Beginners" (YouTube)
- Read: gobyexample.com
- Build: Simple REST API before this project

### React Fundamentals:
- Watch: "React Course" (YouTube)
- Read: Official React documentation
- Build: Simple todo app before this project

### Database/SQL:
- Read: "SQL Tutorial" (mode.com)
- Practice: Write queries for sample database
- Build: Design database schema for simple app

### HTTP/REST:
- Read: "REST API Tutorial"
- Test: Use Postman to interact with public APIs
- Understand: Request/response cycle

---

## 🚨 Common Pitfalls to Avoid

### Backend Mistakes:
1. **Storing plain passwords** - Always hash!
2. **Forgetting CORS headers** - Frontend can't call backend without these
3. **Not validating input** - User sends email: "'; DROP TABLE users; --"
4. **Hardcoding secrets** - Never commit API keys to Git!
5. **No error handling** - Users see confusing error messages

### Frontend Mistakes:
1. **Storing tokens in localStorage** - Can be stolen. Use httpOnly cookies if possible
2. **Not handling loading states** - App freezes while waiting for data
3. **Not handling errors** - API fails, user sees nothing
4. **Not using environment variables** - Hardcode backend URL, can't deploy
5. **Missing TypeScript types** - Catch errors before they happen

### Database Mistakes:
1. **No indexes** - Queries become slow with large datasets
2. **Forgetting ON DELETE CASCADE** - Delete user, orphaned data stays
3. **Wrong data types** - Store password as TEXT instead of VARCHAR
4. **No backups** - Database corrupts, lose all user data

### Git Mistakes:
1. **Committing node_modules** - Makes repo huge
2. **Committing .env files** - Expose secrets
3. **Not writing commit messages** - Future you won't know what changed
4. **Force pushing** - Overwrites team member's work: `git push --force` ☠️

---

## ✅ Success Criteria

How do you know when you're done?

### MVP (Minimum Viable Product) - Week 2
- ✅ Users can register and login
- ✅ Users can play one game mode on kana
- ✅ Progress is tracked and saved
- ✅ User can see their progress on dashboard

### v1.0 - Week 3-4
- ✅ All 4 game modes work
- ✅ All 3 content types (kana, kanji, vocab)
- ✅ Achievement system
- ✅ Statistics/analytics
- ✅ Preferences working (theme, font, language)
- ✅ 90% of endpoints tested
- ✅ App deployed and live

### Post-Launch
- ✅ Users find bugs, you fix them
- ✅ Gather user feedback
- ✅ Add features based on feedback
- ✅ Optimize performance
- ✅ Monitor with analytics

---

## 🎬 Ready to Build?

You now understand:

✅ The overall architecture  
✅ Why we chose these technologies  
✅ How each piece fits together  
✅ What you'll build in each phase  
✅ Common pitfalls to avoid  
✅ How to test and deploy  

**Next steps:**

1. **Set up your development environment** (install all tools)
2. **Read the setup guides** (go-backend-setup-guide.md, react-frontend-setup-guide.md)
3. **Create project structure** (folders, git repos)
4. **Start Phase 1** (Day 1-2 setup)
5. **Work through each phase** in order

**The journey of 1,000 lines of code starts with a single `go run main.go`.**

You've got this! 💪

---

## 📚 Quick Reference

### Useful Documentation Links (within dev-doc):
- `PROJECT-PLAN-SUMMARY.md` - 30-day timeline
- `go-react-kanadojo-plan.md` - Detailed architecture
- `VISUAL-ARCHITECTURE.md` - Diagrams
- `go-backend-setup-guide.md` - Backend setup steps
- `react-frontend-setup-guide.md` - Frontend setup steps
- `api-specification.md` - API endpoints
- `IMPLEMENTATION-CHECKLIST.md` - Track progress
- `code-samples.md` - Code examples

### Commands You'll Use Daily:
```bash
# Go backend
go run ./cmd/server/main.go          # Run server
go build ./cmd/server                 # Build executable
go test ./...                         # Run tests

# React frontend
npm run dev                           # Start dev server
npm run build                         # Build for production
npm test                              # Run tests

# Git
git status                            # See changes
git add .                             # Stage changes
git commit -m "message"               # Commit
git push origin branch-name           # Push to GitHub

# Database
psql -U postgres                      # Connect to PostgreSQL
\l                                    # List databases
\d                                    # List tables
```

### Key Tools:
- **VS Code:** Editor
- **Postman:** Test API
- **DBeaver:** Explore database
- **GitHub:** Version control
- **Terminal:** Run commands

---

## 🎥 You're Ready for the Lecture Series!

This tutorial is the foundation. The actual build guides will walk you through:
- Exact terminal commands to run
- Exact code to write
- Tests to verify everything works
- Common errors and how to fix them

Now, whenever you get stuck on something:
1. Check the specific guide for that phase
2. Review the architecture docs
3. Look at code samples
4. Test with Postman/browser dev tools
5. Ask questions in development communities

**Let's build something awesome!** 🚀

---

**End of Tutorial - You're Ready to Code!**
