# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Heiwadai is a hotel/hospitality management API server built with Go and Connect-Go (gRPC-compatible RPC framework). The system handles check-ins, reservations, coupons, mail magazines, and multi-tenant store management with Supabase integration.

## Architecture

The codebase follows Clean Architecture/DDD patterns:
- `server/core/entity/`: Domain entities
- `server/core/usecase/`: Business logic
- `server/core/infra/`: Infrastructure interfaces  
- `server/infrastructure/`: Infrastructure implementations (DB, cache, storage)
- `server/controller/`: API controllers (Connect-Go handlers)
- `server/api/v1/`: Generated protobuf code
- `server/v1/`: Protocol buffer definitions

## Common Development Commands

### Development
```bash
# Start local environment (PostgreSQL via Supabase, Redis)
docker compose up -d

# Run development server with hot reload
cd server && make dev

# Run server without hot reload
cd server && make run

# Run linter
cd server && make lint
```

### Testing
```bash
# Run all integration tests
cd server && make test-all

# Run specific test suites
cd server && make bookTest      # Booking system tests
cd server && make mailTest      # Mail system tests
cd server && make couponTest    # Coupon system tests
cd server && make checkinTest   # Check-in system tests
```

### Code Generation
```bash
# Generate API code from proto files
cd server && make buf

# Generate ORM models from database
cd server && make sqlboiler
```

### Database Operations
```bash
# Run migrations
cd server && make migrate-up

# Rollback migrations
cd server && make migrate-down

# Create new migration
cd server && make migrate-create TABLE_NAME=table_name

# Initialize database (migrate + generate models)
cd server && make init-db

# Run seeders
cd server && make seeder-store
cd server && make seeder-user
cd server && make seeder-admin
cd server && make seeder-coupon
```

### Installation Commands
```bash
# Install development tools
cd server && make install-devtools

# Install buf and protobuf tools
cd server && make install-buf

# Install migration tool
cd server && make install-migrate

# Install SQLBoiler
cd server && make install-sqlboiler

# Install Node tools (for OpenAPI)
cd server && make install-node
```

### AWS Lambda Deployment
```bash
# Deploy birthday coupon Lambda function (one-time setup)
export CRON_ACCESS_ENDPOINT="your-endpoint-url"
export CRON_ACCESS_SECRET="your-secret"
export CRON_ACCESS_KEY="your-key"
export AWS_PROFILE="your-aws-profile"  # Optional, defaults to 'default'
make deploy-birthday-coupon

# Update Lambda function only
make deploy-lambda

# Test Lambda function manually
make lambda-test

# Delete Lambda resources
make delete-birthday-coupon

# Show Lambda help
make help-lambda

# Use specific AWS profile
AWS_PROFILE=production make deploy-birthday-coupon
```

## Tech Stack

- **Backend**: Go 1.20
- **API Framework**: Connect-Go (Buf)
- **Database**: PostgreSQL (via Supabase)
- **Cache**: Redis 7.0.12
- **ORM**: SQLBoiler v4
- **Migrations**: golang-migrate
- **Auth**: Supabase Auth
- **Storage**: Supabase Storage / AWS S3
- **Deployment**: Fly.io

## Environment Variables

Required `.env` file with:
- `PSQL_*`: PostgreSQL connection details
- `REDIS_*`: Redis connection details  
- `SUPABASE_*`: Supabase project details
- `AWS_*`: S3 configuration (if using S3)
- `MAIL_*`: Email service configuration

## Testing Approach

- Unit tests: Located alongside code files (`*_test.go`)
- Integration tests: Located in `/server/test/` directory
- Test framework: Go standard `testing` package with `testify` assertions
- Database mocking: `go-sqlmock` for unit tests

## Development Workflow

1. Always run `make lint` before committing
2. Generate code after proto changes: `make buf`
3. Regenerate models after DB schema changes: `make sqlboiler`
4. Use `air` for hot reload during development (`make dev`)
5. Test changes with integration tests in `/server/test/`

## Key Implementation Patterns

- Dependency injection via interfaces
- Repository pattern for data access
- Use case pattern for business logic
- Connect-Go interceptors for auth/logging
- Null handling with `volatiletech/null/v8`
- Cache layer with Redis for performance