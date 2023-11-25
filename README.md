## 🎒🚀 LeapKit

LeapKit is a template for building web applications with Go, HTMX and Tailwind CSS. It integrates useful features such as hot code reload and css recompiling.

### Getting started

Install Gonew:

```
go install rsc.io/tmp/gonew@latest 
```

Create a new project from template:

```
gonew github.com/leapkit/template@latest github.com/your/app 
```
⚠️ Important: Gonew does not support modules without github.com or similar use github.com/your/app as the pattern for the module path of the newly created project.

### Setup

Install dependencies:

```
go mod download
go run ./cmd/setup
```

### Running the application

To run the application in development mode execute:

```
go run ./cmd/dev
```
And open http://localhost:3000 in your browser.

