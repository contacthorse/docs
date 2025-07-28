# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This repository contains a Go web server that redirects `docs.contacthorse.com` to a Google Doc containing the actual documentation. The app deploys to fly.io and serves as a simple redirect service.

## Architecture

- **Language**: Go 1.23
- **Deployment Target**: fly.io
- **Functionality**: HTTP 307 temporary redirect to Google Doc
- **Domain**: docs.contacthorse.com
- **Google Doc URL**: https://docs.google.com/document/d/e/2PACX-1vSsmM8yaTYx5tQst5tOLT85r7JjvEb0G-abrOZiblYvG2Eymkp4qZYjoSFr9_qemrRy44JNklGgmv96/pub

## Development Commands

```bash
# Run locally
go run main.go

# Build binary
go build -o main .

# Deploy to fly.io
fly deploy

# Check deployment status
fly status

# View logs
fly logs

# Open the deployed app
fly open
```

## File Structure

- `main.go` - Go web server with redirect handler
- `go.mod` - Go module configuration
- `Dockerfile` - Multi-stage build for production deployment
- `fly.toml` - fly.io configuration
- `README.md` - Contains the Google Doc URL for reference

## Key Notes

- Uses Go's standard library HTTP package for minimal dependencies
- Multi-stage Docker build for optimized container size
- Configured for auto-scaling with min 0 machines
- 256MB memory allocation for lightweight operation