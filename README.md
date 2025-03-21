# IncidentFlow

- cmd: CLI-related commands (Cobra)
- api: API-related files 
- internal: Core business logic shared by CLI and API

- internal/diagram: Mermaid.js related utilities 

- internal/models: Data models
- assets: Stores output files (generated diagrams)


**main.go** is primary entrypoint for Go application, it initializes and runs either the CLI or the API depending on how the program is executed.


1. So when the CLI is run, it delegates execution to cmd/root.go

2. When running the API, it starts the server from api/server.go.
