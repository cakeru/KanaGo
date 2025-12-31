# Frontend Setup Guide: Building with React + Vite

This guide walks you through setting up a modern React frontend for your Japanese learning platform.

---

## Prerequisites

- **Node.js 18+** - https://nodejs.org/
- **npm or yarn** - Package manager
- **Git** - Version control

---

## Project Initialization

### Step 1: Create React + TypeScript Project with Vite

```bash
npm create vite@latest go-kanadojo-frontend -- --template react-ts
cd go-kanadojo-frontend
npm install
```

### Step 2: Create Directory Structure

```bash
mkdir -p src/pages
mkdir -p src/features/{kana,kanji,vocabulary,achievements,progress,preferences,auth}
mkdir -p src/shared/{components,hooks,lib,store,types}
mkdir -p src/assets/{sounds,wallpapers}
mkdir -p src/styles
mkdir -p public/data
mkdir -p src/__tests__
```

---

## Core Dependencies

### Install Required Packages

```bash
# UI Framework & Styling
npm install @radix-ui/react-dialog @radix-ui/react-select tailwind-merge clsx class-variance-authority
npm install tailwindcss postcss autoprefixer
npm install lucide-react @fortawesome/fontawesome-svg-core @fortawesome/react-fontawesome @fortawesome/free-solid-svg-icons

# State Management
npm install zustand
npm install immer

# Routing
npm install react-router-dom

# HTTP Client
npm install axios

# Form Handling & Validation
npm install zod react-hook-form @hookform/resolvers

# Animations
npm install framer-motion

# Internationalization (optional, can use simple JSON)
npm install i18next react-i18next

# Utilities
npm install date-fns uuid

# Development Dependencies
npm install -D typescript @types/react @types/react-dom @types/node
npm install -D @vitejs/plugin-react
npm install -D eslint eslint-config-prettier prettier
npm install -D vitest @testing-library/react @testing-library/dom jsdom
npm install -D tailwindcss postcss autoprefixer
```

### Optional for Advanced Features

```bash
# WebSocket support (for real-time features)
npm install socket.io-client

# Data visualization (for stats)
npm install recharts

# Notifications
npm install react-hot-toast

# Local storage with persistence
npm install js-cookie
```

---

## Configuration Files

### `tailwind.config.js`

```javascript
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        // Add custom colors as needed
        primary: '#3B82F6',
        secondary: '#8B5CF6',
      },
      fontFamily: {
        // Add Japanese fonts
        'jp': ['"Noto Sans JP"', 'sans-serif'],
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
        slideUp: {
          '0%': { transform: 'translateY(10px)', opacity: '0' },
          '100%': { transform: 'translateY(0)', opacity: '1' },
        },
      },
      animation: {
        fadeIn: 'fadeIn 0.3s ease-in-out',
        slideUp: 'slideUp 0.3s ease-in-out',
      },
    },
  },
  plugins: [],
}
```

### `postcss.config.js`

```javascript
export default {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
}
```

### `vite.config.ts`

```typescript
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
```

### `tsconfig.json`

```json
{
  "compilerOptions": {
    "target": "ES2020",
    "useDefineForClassFields": true,
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "skipLibCheck": true,
    "esModuleInterop": true,
    "allowSyntheticDefaultImports": true,
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noFallthroughCasesInSwitch": true,
    "resolveJsonModule": true,
    "baseUrl": ".",
    "paths": {
      "@/*": ["./src/*"]
    }
  },
  "include": ["src"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
```

### `vitest.config.ts`

```typescript
import { defineConfig } from 'vitest/config'
import react from '@vitejs/plugin-react'
import path from 'path'

export default defineConfig({
  plugins: [react()],
  test: {
    globals: true,
    environment: 'jsdom',
    setupFiles: [],
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
})
```

---

## Core Setup Files

### `src/main.tsx` - Entry Point

```typescript
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './styles/globals.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
```

### `index.html`

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="/vite.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>KanaDojo - Master Japanese</title>
    <meta name="description" content="Learn Japanese with KanaDojo" />
  </head>
  <body>
    <div id="root"></div>
    <script type="module" src="/src/main.tsx"></script>
  </body>
</html>
```

---

## State Management

### `src/shared/store/authStore.ts`

```typescript
import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import { User } from '@/shared/types'

interface AuthState {
  user: User | null
  accessToken: string | null
  refreshToken: string | null
  isLoading: boolean
  error: string | null

  login: (email: string, password: string) => Promise<void>
  register: (email: string, username: string, password: string) => Promise<void>
  logout: () => void
  setUser: (user: User) => void
  setTokens: (accessToken: string, refreshToken: string) => void
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      accessToken: null,
      refreshToken: null,
      isLoading: false,
      error: null,

      login: async (email, password) => {
        set({ isLoading: true, error: null })
        try {
          const response = await fetch('/api/v1/auth/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password }),
          })
          const data = await response.json()
          
          if (!response.ok) {
            throw new Error(data.error || 'Login failed')
          }

          set({
            user: data.data.user,
            accessToken: data.data.access_token,
            refreshToken: data.data.refresh_token,
            isLoading: false,
          })
        } catch (error) {
          set({
            error: error instanceof Error ? error.message : 'An error occurred',
            isLoading: false,
          })
          throw error
        }
      },

      register: async (email, username, password) => {
        set({ isLoading: true, error: null })
        try {
          const response = await fetch('/api/v1/auth/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, username, password }),
          })
          const data = await response.json()

          if (!response.ok) {
            throw new Error(data.error || 'Registration failed')
          }

          set({
            user: data.data.user,
            accessToken: data.data.access_token,
            refreshToken: data.data.refresh_token,
            isLoading: false,
          })
        } catch (error) {
          set({
            error: error instanceof Error ? error.message : 'An error occurred',
            isLoading: false,
          })
          throw error
        }
      },

      logout: () => {
        set({ user: null, accessToken: null, refreshToken: null })
      },

      setUser: (user) => set({ user }),
      setTokens: (accessToken, refreshToken) =>
        set({ accessToken, refreshToken }),
    }),
    {
      name: 'auth-storage',
      partialize: (state) => ({
        user: state.user,
        accessToken: state.accessToken,
        refreshToken: state.refreshToken,
      }),
    }
  )
)
```

### `src/shared/store/preferencesStore.ts`

```typescript
import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface PreferencesState {
  theme: string
  font: string
  soundEnabled: boolean
  darkMode: boolean
  language: string

  setTheme: (theme: string) => void
  setFont: (font: string) => void
  toggleSound: () => void
  toggleDarkMode: () => void
  setLanguage: (language: string) => void
}

export const usePreferencesStore = create<PreferencesState>()(
  persist(
    (set) => ({
      theme: 'default',
      font: 'default',
      soundEnabled: true,
      darkMode: false,
      language: 'en',

      setTheme: (theme) => set({ theme }),
      setFont: (font) => set({ font }),
      toggleSound: () => set((state) => ({ soundEnabled: !state.soundEnabled })),
      toggleDarkMode: () => set((state) => ({ darkMode: !state.darkMode })),
      setLanguage: (language) => set({ language }),
    }),
    {
      name: 'preferences-storage',
    }
  )
)
```

---

## API Client

### `src/shared/lib/api.ts`

```typescript
import { useAuthStore } from '@/shared/store/authStore'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'

interface RequestOptions extends RequestInit {
  headers?: Record<string, string>
}

async function apiCall<T>(
  endpoint: string,
  options: RequestOptions = {}
): Promise<T> {
  const { accessToken } = useAuthStore.getState()

  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    ...options.headers,
  }

  if (accessToken) {
    headers['Authorization'] = `Bearer ${accessToken}`
  }

  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    headers,
  })

  if (!response.ok) {
    // Handle token refresh on 401
    if (response.status === 401 && accessToken) {
      const { refreshToken } = useAuthStore.getState()
      if (refreshToken) {
        try {
          const refreshResponse = await fetch(`${API_BASE_URL}/auth/refresh`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ refresh_token: refreshToken }),
          })

          if (refreshResponse.ok) {
            const data = await refreshResponse.json()
            const { setTokens } = useAuthStore.getState()
            setTokens(data.data.access_token, data.data.refresh_token)

            // Retry original request
            return apiCall(endpoint, options)
          }
        } catch (error) {
          useAuthStore.getState().logout()
        }
      }
    }

    const errorData = await response.json().catch(() => ({}))
    throw new Error(errorData.error || `API error: ${response.status}`)
  }

  const data = await response.json()
  return data.data as T
}

// Specific API methods
export const api = {
  // Auth
  login: (email: string, password: string) =>
    apiCall('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    }),

  register: (email: string, username: string, password: string) =>
    apiCall('/auth/register', {
      method: 'POST',
      body: JSON.stringify({ email, username, password }),
    }),

  // Content
  getKana: () => apiCall('/kana'),
  getKanji: () => apiCall('/kanji'),
  getVocabulary: () => apiCall('/vocabulary'),

  // Progress
  getStats: () => apiCall('/stats'),
  updateProgress: (data: any) =>
    apiCall('/progress/update', {
      method: 'POST',
      body: JSON.stringify(data),
    }),

  // Preferences
  getPreferences: () => apiCall('/preferences'),
  updatePreferences: (data: any) =>
    apiCall('/preferences', {
      method: 'PUT',
      body: JSON.stringify(data),
    }),

  // Achievements
  getAchievements: () => apiCall('/achievements'),
  getUnlockedAchievements: () => apiCall('/achievements/unlocked'),
}
```

---

## Custom Hooks

### `src/shared/hooks/useAudio.ts`

```typescript
import { usePreferencesStore } from '@/shared/store/preferencesStore'

export function useAudio(soundFile: string) {
  const { soundEnabled } = usePreferencesStore()

  const play = () => {
    if (!soundEnabled) return

    try {
      const audio = new Audio(`/sounds/${soundFile}`)
      audio.play().catch((err) => console.warn('Audio play failed:', err))
    } catch (error) {
      console.warn('Audio error:', error)
    }
  }

  return { play }
}

export function useCorrectSound() {
  return useAudio('correct.mp3')
}

export function useErrorSound() {
  return useAudio('error.mp3')
}

export function useClickSound() {
  return useAudio('click.mp3')
}
```

### `src/shared/hooks/useFetch.ts`

```typescript
import { useState, useEffect } from 'react'
import { api } from '@/shared/lib/api'

export function useFetch<T>(apiCall: () => Promise<T>) {
  const [data, setData] = useState<T | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true)
        const result = await apiCall()
        setData(result)
      } catch (err) {
        setError(err instanceof Error ? err.message : 'An error occurred')
      } finally {
        setLoading(false)
      }
    }

    fetchData()
  }, [])

  return { data, loading, error }
}
```

---

## Basic Components

### `src/shared/components/Button.tsx`

```typescript
import { ReactNode } from 'react'
import { cn } from '@/shared/lib/utils'

interface ButtonProps {
  children: ReactNode
  onClick?: () => void
  className?: string
  variant?: 'primary' | 'secondary' | 'outline'
  disabled?: boolean
  type?: 'button' | 'submit' | 'reset'
}

export function Button({
  children,
  className,
  variant = 'primary',
  disabled = false,
  ...props
}: ButtonProps) {
  const baseStyles = 'px-4 py-2 rounded-lg font-medium transition-colors'
  const variants = {
    primary: 'bg-blue-500 text-white hover:bg-blue-600 disabled:bg-gray-400',
    secondary: 'bg-purple-500 text-white hover:bg-purple-600 disabled:bg-gray-400',
    outline: 'border-2 border-blue-500 text-blue-500 hover:bg-blue-50',
  }

  return (
    <button
      className={cn(baseStyles, variants[variant], disabled && 'cursor-not-allowed', className)}
      disabled={disabled}
      {...props}
    >
      {children}
    </button>
  )
}
```

### `src/shared/lib/utils.ts`

```typescript
export function cn(...classes: (string | undefined | false)[]): string {
  return classes.filter(Boolean).join(' ')
}

export function generateRandomChoices(correct: string, all: string[], count: number = 3): string[] {
  const others = all.filter((item) => item !== correct)
  const shuffled = others.sort(() => Math.random() - 0.5)
  return [correct, ...shuffled.slice(0, count)].sort(() => Math.random() - 0.5)
}

export function calculateAccuracy(correct: number, total: number): number {
  if (total === 0) return 0
  return Math.round((correct / total) * 100)
}
```

---

## Routing Setup

### `src/App.tsx`

```typescript
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { useAuthStore } from '@/shared/store/authStore'
import { usePreferencesStore } from '@/shared/store/preferencesStore'

// Pages
import Home from '@/pages/Home'
import Login from '@/pages/Login'
import Register from '@/pages/Register'
import KanaTraining from '@/pages/KanaTraining'
import KanjiTraining from '@/pages/KanjiTraining'
import VocabularyTraining from '@/pages/VocabularyTraining'
import Progress from '@/pages/Progress'
import Achievements from '@/pages/Achievements'
import Preferences from '@/pages/Preferences'
import NotFound from '@/pages/NotFound'

// Components
import PrivateRoute from '@/shared/components/PrivateRoute'

function App() {
  const { darkMode } = usePreferencesStore()

  return (
    <div className={darkMode ? 'dark' : ''}>
      <BrowserRouter>
        <Routes>
          {/* Public routes */}
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />

          {/* Protected routes */}
          <Route
            path="/kana"
            element={
              <PrivateRoute>
                <KanaTraining />
              </PrivateRoute>
            }
          />
          <Route
            path="/kanji"
            element={
              <PrivateRoute>
                <KanjiTraining />
              </PrivateRoute>
            }
          />
          <Route
            path="/vocabulary"
            element={
              <PrivateRoute>
                <VocabularyTraining />
              </PrivateRoute>
            }
          />
          <Route
            path="/progress"
            element={
              <PrivateRoute>
                <Progress />
              </PrivateRoute>
            }
          />
          <Route
            path="/achievements"
            element={
              <PrivateRoute>
                <Achievements />
              </PrivateRoute>
            }
          />
          <Route
            path="/preferences"
            element={
              <PrivateRoute>
                <Preferences />
              </PrivateRoute>
            }
          />

          {/* 404 */}
          <Route path="*" element={<NotFound />} />
        </Routes>
      </BrowserRouter>
    </div>
  )
}

export default App
```

### `src/shared/components/PrivateRoute.tsx`

```typescript
import { ReactNode } from 'react'
import { Navigate } from 'react-router-dom'
import { useAuthStore } from '@/shared/store/authStore'

interface PrivateRouteProps {
  children: ReactNode
}

export default function PrivateRoute({ children }: PrivateRouteProps) {
  const { user } = useAuthStore()

  if (!user) {
    return <Navigate to="/login" replace />
  }

  return <>{children}</>
}
```

---

## Package.json Scripts

### `package.json`

```json
{
  "name": "go-kanadojo-frontend",
  "private": true,
  "version": "0.1.0",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "tsc && vite build",
    "preview": "vite preview",
    "test": "vitest run",
    "test:watch": "vitest",
    "lint": "eslint src --ext ts,tsx --report-unused-disable-directives",
    "lint:fix": "eslint src --ext ts,tsx --fix",
    "format": "prettier --write src",
    "format:check": "prettier --check src",
    "type-check": "tsc --noEmit"
  },
  "dependencies": {
    "@radix-ui/react-dialog": "^1.1.1",
    "@radix-ui/react-select": "^2.0.0",
    "axios": "^1.6.0",
    "framer-motion": "^10.16.0",
    "lucide-react": "^0.292.0",
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "react-hook-form": "^7.48.0",
    "react-router-dom": "^6.20.0",
    "zustand": "^4.4.0"
  },
  "devDependencies": {
    "@testing-library/dom": "^9.3.3",
    "@testing-library/react": "^14.1.2",
    "@types/node": "^20.10.0",
    "@types/react": "^18.2.37",
    "@types/react-dom": "^18.2.15",
    "@vitejs/plugin-react": "^4.2.0",
    "autoprefixer": "^10.4.16",
    "eslint": "^8.54.0",
    "postcss": "^8.4.32",
    "prettier": "^3.1.0",
    "tailwindcss": "^3.3.6",
    "typescript": "^5.3.3",
    "vite": "^5.0.8",
    "vitest": "^1.0.4"
  }
}
```

---

## Environment Setup

### `.env.example`

```env
VITE_API_URL=http://localhost:8080/api/v1
VITE_ENV=development
```

---

## Running the Frontend

```bash
# Install dependencies
npm install

# Development server (hot reload at http://localhost:5173)
npm run dev

# Build for production
npm run build

# Run tests
npm run test

# Type checking
npm run type-check

# Linting
npm run lint

# Format code
npm run format
```

---

## Next Steps

1. **Create page components** in `src/pages/`
2. **Implement feature modules** in `src/features/`
3. **Add game logic** components (card display, multiple choice, input)
4. **Integrate with backend** via API calls
5. **Build statistics dashboard** with charts
6. **Add theme system** with 100+ themes
7. **Implement internationalization** with i18n
8. **Add unit and integration tests**

This foundation provides a modern, scalable React frontend!
