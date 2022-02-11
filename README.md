# qfy

Self-hosted implementation of Synthetics - Monitoring checks to validate your service availability.

## Running
```bash
# Start the API server with
go run cmd/cmd.go server

# Start an agent with
go run cmd/cmd.go agent

# Populate DB with fixtures
go run cmd/cmd.go fixtures
```