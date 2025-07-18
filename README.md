# Goilerplate

A modern Go web application boilerplate with empl templating and Tailwind CSS.

## Quick Start

```bash
# Install dependencies
go mod download

# Run development server
make dev
```

The server will start at `http://localhost:8090`

## Tech Stack

- **Go** - Backend language
- **templ** - Type-safe templating
- **Tailwind CSS** - Utility-first CSS
- **Alpine.js** - Lightweight JavaScript framework
- **HTMX** - Hypermedia Framework

## Structure

```
   assets/          # Static assets (CSS, JS, images)
   handlers/        # HTTP handlers
   ui/              # UI components and pages
      components/  # Reusable components
      layouts/     # Page layouts
      pages/       # Page templates
   main.go          # Application entry point
```

## License

This is proprietary software. See LICENSE file for details.

