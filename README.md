# Harmony

A Discord clone made by two programmers who, frankly, don't really like the way that Discord operates

## Code structure

### Client - `client/`

- The client uses React with Typescript to provide a simple, reactive UI
- File layout
  - `components/` - components that do not make up a full page
  - `pages/` - a component that makes up the full page (should be added under Routes in `App.tsx`)

### Server - `server/`

- The server uses the Gin framework with Golang to create a fast backend

## Self-hosting

## Links

- This project was created from my own [Gin + React template](https://github.com/narutopig/harmony)
