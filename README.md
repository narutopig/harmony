# Gin with React Template

## Client (`client/`)

- The client uses React with Typescript to provide a simple, reactive UI
- File layout (recommended)
  - `components/` - components that do not make up a full page
  - `pages/` - a component that makes up the full page (should be added under Routes in `App.tsx`)

## Server (`server/`)

- The server uses the Gin framework with Golang to create a fast backend

## Troubleshooting

`server/main.go:6:2: no required module provides package github.com/gin-gonic/contrib/static: go.mod file not found in current directory or any parent directory; see 'go help modules'`

- Fix: Make sure that your current directory is `server/`; this error is caused by the lack of a go.mod file in the current directory
