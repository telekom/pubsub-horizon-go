<!--
Copyright 2025 Deutsche Telekom AG

SPDX-License-Identifier: Apache-2.0
-->
# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Horizon Go is a shared library for Go components within the Horizon ecosystem, a pub/sub messaging platform. It provides common abstractions, types, and utilities used across multiple Horizon services, including:

- **Cache abstraction**: Unified interface for Hazelcast cache operations with listener support
- **Message types**: Data structures for published messages, events, and circuit breaker messages
- **Resource types**: Subscription resource definitions with metadata
- **Validation**: Custom validators for event types, ISO timestamps, and other domain-specific formats
- **Tracing**: OpenTelemetry-based distributed tracing wrapper with detailed span management
- **Enums**: Type-safe enumerations for message status, delivery types, event retention times, circuit breaker status, and response filter modes
- **Testing utilities**: Docker-based integration test helpers for Hazelcast

This library is consumed by other Horizon Go services (such as Quasar) and provides the foundation for consistent data handling across the ecosystem.

## Commands

### Build
```bash
go build
```

### Run Tests
```bash
# Run all tests with coverage
go test -v ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./... -tags=testing

# Run tests for a specific package
go test -v ./cache -tags=testing

# Run a single test
go test -v ./cache -run TestNewCache -tags=testing
```

### Linting
Run linting to check code quality and format/fix:
```bash
golangci-lint run --fix
```

## Architecture

### Core Packages

**cache/** - Hazelcast Cache Abstraction
- `Cache[T]` interface: Generic cache interface for Put/Get/Delete/AddListener operations
- `HazelcastCache[T]`: Hazelcast-specific implementation with query support using predicates
- `Listener[T]` interface: Event listener for cache entry changes (Add/Update/Delete)
- Thread-safe operations using Hazelcast Go client
- Supports generic types for type-safe cache operations

**message/** - Message Data Structures
- `PublishedMessage`: Core message structure with UUID, environment, status, event, HTTP headers, and additional fields
- `Event`: Event data with type, ID, timestamp, and data payload
- `CircuitBreakerMessage`: Circuit breaker state tracking for subscriptions
- `Status`: Message processing status tracking

**resource/** - Kubernetes Resource Definitions
- `SubscriptionResource`: Kubernetes CR structure for subscriptions
- `Subscription`: Subscription configuration with triggers, callbacks, delivery types, circuit breaker settings
- `SubscriptionTrigger`: Filter configurations for events (response filters, selection filters)
- Used by services that watch or manage subscription resources

**enum/** - Type-Safe Enumerations
- `DeliveryType`: SSE (server-sent events) vs callback delivery modes
- `MessageStatus`: Message processing lifecycle states
- `CircuitBreakerStatus`: Circuit breaker states (CLOSED, OPEN, HALF_OPEN)
- `EventRetentionTime`: Event retention duration configurations
- `ResponseFilterMode`: Filter matching modes (EXCLUDE_FILTER, INCLUDE_FILTER)
- All enums support JSON marshaling/unmarshaling with custom parsing

**tracing/** - OpenTelemetry Tracing Wrapper
- `TraceContext`: Manages hierarchical spans with OpenTelemetry
- Supports detailed tracing mode for fine-grained observability
- Provides span lifecycle management (start, end, attributes)
- Integrates with B3 propagation for distributed tracing
- Helper functions for extracting trace context from Kafka headers

**validation/** - Custom Validators
- Extends `go-playground/validator` with domain-specific validation rules
- `ValidateEventType`: Validates event type format (domain/type/version)
- `ValidateIsoTime`: Validates ISO 8601 timestamp format
- Use `NewValidator()` to get a pre-configured validator instance with custom rules

**util/** - Utility Functions
- `HazelcastZerologLogger`: Bridges Hazelcast logger to zerolog for unified logging
- Translates Hazelcast log weights to zerolog levels

**test/** - Testing Utilities
- `StartDocker()/StopDocker()`: Manages Hazelcast container lifecycle using dockertest
- `EnvOrDefault()`: Helper for environment variable fallbacks in tests
- Integration tests use `//go:build testing` tag
- Hazelcast container configuration: cluster name "horizon", exposed on port 5701

**types/** - Custom Types
- `Timestamp`: Custom timestamp type with JSON marshaling

## Code Conventions

This project follows the Uber Go Style Guide with additional project-specific conventions. See @pubsub-horizon-shareddata/CLAUDE.md

### Linting
Run linting to check code quality and format/fix:
```bash
golangci-lint run --fix
```

## Important Implementation Notes

### Generic Types
- The cache package uses Go generics (`Cache[T]`, `HazelcastCache[T]`) for type-safe operations
- When using the cache, specify the concrete type: `NewHazelcastCache[MyStruct](config)`

### Hazelcast Integration
- Default cluster name: "horizon"
- Uses JSON serialization for Hazelcast entries
- Supports predicate-based queries via `GetQuery()`
- Listener pattern for cache entry notifications (add/update/delete events)

### Testing with Docker
- Integration tests require Docker for Hazelcast container
- Use `test.StartDocker()` in `TestMain()` to set up containers
- Use `test.StopDocker()` to clean up after tests
- Environment variables for configuration: `HAZELCAST_IMAGE`, `HAZELCAST_TAG`, `HAZELCAST_HOST`, `HAZELCAST_PORT`
- Tests must have `//go:build testing` build tag

### Tracing
- `TraceContext` wraps OpenTelemetry tracing with simplified API
- Detailed mode enables fine-grained tracing (use `StartDetailedSpan()` for optional spans)
- Always end spans to prevent resource leaks
- Use `CurrentSpan()` to access the active span for setting attributes

### Validation
- Use `validation.NewValidator()` to get validator with custom rules registered
- Custom validators: `eventType`, `isoTime`
- Struct tags: `validate:"eventType,required"`, `validate:"isoTime"`

### Enum Parsing
- All enums support string parsing via `Parse<EnumType>(string)` functions
- Custom JSON unmarshaling handles various formats (quoted, unquoted, aliases)
- Example: `ParseDeliveryType("sse")` returns `DeliveryTypeSse`

## Code Quality

### Pre-commit Hooks
The project uses pre-commit hooks:
- REUSE compliance checks (copyright/licensing)
- Conventional commits enforcement

### Licensing
- Every file must have SPDX license header
- Project follows REUSE standard for software licensing
- Use Apache-2.0 license for new files

## Additional Instructions
- Use zerolog instead of fmt for logging
- Whenever you finish work, run the linter and the entire test suite. If issues arise, fix them.
