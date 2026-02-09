# Markito API

REST API for creating and sharing Markdown documents.

## Overview

Markito is a Markdown document sharing service with a simple REST API for programmatic document creation and updates.

Base URL: `http://localhost:3000` (or your deployed instance)

## Endpoints

### Create Document

Creates a new Markdown document.

```
POST /api/documents
```

**Request**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `content` | string | Yes | Markdown content to store |

**Example**

```bash
curl -X POST http://localhost:3000/api/documents \
  -d "content=# Hello World\n\nThis is my document."
```

**Response (200 OK)**

```json
{
  "url": "http://localhost:3000/abc123",
  "markdown_url": "http://localhost:3000/raw/abc123/md",
  "html_url": "http://localhost:3000/raw/abc123/html"
}
```

| Field | Description |
|-------|-------------|
| `url` | View/edit the document in the web UI |
| `markdown_url` | Raw Markdown content |
| `html_url` | Rendered HTML content |

### Update Document

Updates an existing document by ID.

```
PUT /api/documents/{id}
```

**Parameters**

| Name | In | Required | Description |
|------|-----|----------|-------------|
| `id` | path | Yes | Document ID from create response |

**Request Body**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `content` | string | Yes | New Markdown content |

**Example**

```bash
curl -X PUT http://localhost:3000/api/documents/abc123 \
  -d "content=# Updated\n\nNew content here."
```

**Response (200 OK)**

Same format as Create Document.

## Error Responses

| Status | Response | Meaning |
|--------|----------|---------|
| 400 | `{"error":"content is required"}` | Missing content field |
| 400 | `{"error":"id is required"}` | Missing document ID (update) |
| 500 | `{"error":"failed to save document"}` | Database/server error |

## Raw Content Endpoints

These endpoints serve document content directly:

| Endpoint | Description | Content-Type |
|----------|-------------|--------------|
| `GET /raw/{id}/md` | Raw Markdown | `text/markdown` |
| `GET /raw/{id}/html` | Rendered HTML | `text/html` |

## Usage Patterns

### Creating a Shareable Document

1. POST to `/api/documents` with your Markdown
2. Store the returned `url` for sharing
3. Use `markdown_url` or `html_url` for embedding

### Updating Content

1. Keep the document ID from the initial create URL
2. PUT to `/api/documents/{id}` with new content
3. URLs remain the same, content updates immediately

## Environment Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `BASE_URL` | `http://localhost:3000` | URL prefix in API responses |

## Code Examples

### Go

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type DocumentResponse struct {
    URL      string `json:"url"`
    Markdown string `json:"markdown_url"`
    HTML     string `json:"html_url"`
}

func createDocument(content string) (*DocumentResponse, error) {
    data := bytes.NewBufferString("content=" + url.QueryEscape(content))
    resp, err := http.Post(
        "http://localhost:3000/api/documents",
        "application/x-www-form-urlencoded",
        data,
    )
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var doc DocumentResponse
    if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
        return nil, err
    }
    return &doc, nil
}
```

### Python

```python
import requests

def create_document(content: str) -> dict:
    response = requests.post(
        "http://localhost:3000/api/documents",
        data={"content": content}
    )
    response.raise_for_status()
    return response.json()

def update_document(doc_id: str, content: str) -> dict:
    response = requests.put(
        f"http://localhost:3000/api/documents/{doc_id}",
        data={"content": content}
    )
    response.raise_for_status()
    return response.json()
```

### JavaScript/TypeScript

```typescript
async function createDocument(content: string) {
    const response = await fetch('http://localhost:3000/api/documents', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: new URLSearchParams({ content })
    });
    return response.json();
}

async function updateDocument(id: string, content: string) {
    const response = await fetch(`http://localhost:3000/api/documents/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: new URLSearchParams({ content })
    });
    return response.json();
}
```

## Testing with cURL

```bash
# Create document
curl -X POST http://localhost:3000/api/documents \
  -d "content=# Test Document" | jq

# Update document (replace ID)
curl -X PUT http://localhost:3000/api/documents/YOUR_ID \
  -d "content=# Updated Document" | jq

# Get raw Markdown
curl http://localhost:3000/raw/YOUR_ID/md

# Get rendered HTML
curl http://localhost:3000/raw/YOUR_ID/html
```
