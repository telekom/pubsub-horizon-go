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

This project follows the Uber Go Style Guide with additional project-specific conventions. Key principles include:

### Naming Conventions
- **Meaningful names** for all variables, functions, structs
- **Error variables**: Named after the function that returns them, e.g., `errValidateCustomerFeedback := validateCustomerFeedback(...)`
- **Error types**: Prefix with `Err` or `err` for exported/unexported error variables; suffix with `Error` for custom error types
- **Type aliases**: Wrap external types in package-local type aliases for easier import management
- **Package names**: Short, concise, evocative, lowercase, single-word names (avoid "common", "util", "shared", "lib" as catch-all packages)
- **Interface names**: One-method interfaces use method name + "-er" suffix (Reader, Writer, Formatter)
- **MixedCaps**: Use MixedCaps or mixedCaps rather than underscores for multiword names
- **Functions/Variables**: Exported start with uppercase, unexported with lowercase (camel case)
- **Constants**: Use all capital letters with underscores, e.g., `MAX_RETRY_COUNT`
- **Boolean variables**: Prefix with Has, Is, Can, or Allow, e.g., `isConnected`, `hasPermission`
- **Getters**: Avoid "Get" prefix; use `user.Name()` instead of `user.GetName()`
- **File names**: Single lowercase words; compound names use underscores; test files use `_test.go` suffix
- **Unexported globals**: Prefix with `_` (except error values with `err` prefix)

### Pointers and Interfaces
- **Never use pointers to interfaces** - interfaces are already reference types
- **Verify interface compliance at compile time**: Use `var _ Interface = (*Type)(nil)` for exported types
- **Receivers**: Methods with value receivers can be called on pointers and values; pointer receivers only on pointers/addressable values
- **Accept interfaces, return structs**: Interfaces declared on consumer side, not producer side

### Error Handling
- **Always check and handle errors explicitly**
- **Handle errors once**: Don't log and return; choose one approach
- **Error wrapping**: Use `fmt.Errorf("context: %w", err)` to wrap errors for traceability
- **Error matching**: Use `%w` if callers should match the error; use `%v` to obfuscate
- **Avoid "failed to" prefix**: Keep context succinct (use "new store" not "failed to create new store")
- **Return errors from functions**: Only call `os.Exit` or `log.Fatal` in `main()`

### Nil Handling and Zero Values
- **Internal functions**: Do NOT check input parameters for nil (caller's responsibility)
- **External functions**: DO check input parameters for nil
- **Exception**: Always validate input structs
- **nil is a valid slice**: Return `nil` instead of `[]T{}` for empty slices; check `len(s) == 0` not `s == nil`
- **Zero-value mutexes are valid**: Don't use `new(sync.Mutex)`; use `var mu sync.Mutex`

### Structs
- **Constructors required**: Always use constructors to instantiate structs (except parameter structs)
- **Field names in initialization**: Always specify field names when initializing structs (enforced by `go vet`)
- **Omit zero values**: Don't specify zero-value fields unless they provide meaningful context
- **Use `var` for zero-value structs**: `var user User` instead of `user := User{}`
- **Struct references**: Use `&T{}` instead of `new(T)` for consistency
- **Parameter structs**: Instantiate inline, must be validated in the function
- **Avoid embedding in public structs**: Embedding leaks implementation details and inhibits evolution
- **Embedded fields**: Place at top of struct with blank line separator

### Immutability and State
- **Favor immutability**: Pass structs by value and return new structs rather than mutating pointers (unless performance-critical)
- **Data flow transparency**: Write code with straightforward, transparent data flow
- **Avoid mutable globals**: Use dependency injection instead of global variables
- **Avoid `init()`**: If unavoidable, be deterministic, avoid I/O, no global state; use constructors or main() instead

### Dependencies and Concurrency
- Use **dependency injection** (constructor functions)
- Avoid global state
- Follow **inversion of control** principle
- **Don't panic**: Return errors instead of panicking (except for truly irrecoverable situations)
- **Use goroutines safely**: Always have a way to stop goroutines and wait for them to exit
- **No goroutines in `init()`**: Expose objects that manage goroutine lifetimes
- **Propagate context**: Always propagate `context.Context` for cancellation
- **Defer to clean up**: Use defer for resource cleanup (files, locks, etc.)

### Maps, Slices, and Collections
- **Copy slices and maps at boundaries**: Prevent unintended mutations when receiving/returning
- **Specify container capacity**: Use `make(map[T]T, size)` and `make([]T, 0, capacity)` when size is known
- **Channel size**: Channels should be unbuffered or size 1 (any other size requires scrutiny)
- **Map initialization**: Use `make()` for empty maps; use literals for fixed sets of elements

### Time Handling
- **Use `time.Time`** for instants of time (not int/string)
- **Use `time.Duration`** for periods of time (not int/float)
- **Use RFC 3339** format for string timestamps when needed
- **AddDate vs Add**: Use `AddDate` for calendar arithmetic; use `Add` for duration arithmetic

### Variables and Scope
- **Short variable declarations**: Use `:=` when setting explicit values; use `var` for zero values
- **Reduce scope**: Declare variables in smallest scope possible; use inline declarations with if statements
- **Local constants**: Don't make constants global unless used across functions/files
- **Top-level declarations**: Use `var` keyword without type (unless type differs from expression)

### Code Organization and Style
- **Use guard clauses (reverse ifs)**: Check for error/invalid conditions first and return early to avoid deeply nested code
- **Reduce nesting**: Handle errors/special cases first and return early
- **Unnecessary else**: Eliminate else blocks when variable can be set with single if
- **Function grouping**: Sort functions by receiver; place utility functions at end
- **Import groups**: Standard library, then everything else (blank line between)
- **Group similar declarations**: Use `const ()`, `var ()`, `type ()` blocks for related declarations
- **Be consistent**: Consistency is more important than individual preferences

### Testing
- **Use table-driven tests** with subtests for repetitive test logic
- **Avoid unnecessary complexity**: Split complex table tests into multiple tests or tables
- **Test tables convention**: Slice named `tests`, variable `tt`, fields prefixed with `give`/`want`
- **Mock external interfaces**: Use Mockery for generating mocks
- **Define mock expectations**: In mutation functions within test tables
- **Separate test types**: Unit tests (fast) vs integration tests (slower, use `dockertest`)
- **Coverage**: Ensure test coverage for all exported functions
- **Use goroutine leak detection**: Use `go.uber.org/goleak` to test for goroutine leaks, if needed
- **Test build tag**: Integration tests use `//go:build testing` tag

### Performance
- **Prefer `strconv` over `fmt`** for primitive conversions
- **Avoid repeated string-to-byte conversions**: Convert once and reuse
- **Atomic operations**: If needed, use `sync/atomic` or `go.uber.org/atomic` package for thread-safe operations

### Patterns
- **Functional Options**: Use for optional constructor arguments (variadic `...Option`)
- **Exit Once**: Prefer single `os.Exit` call in `main()`; use `run() error` pattern
- **Field tags**: Always use field tags in marshaled structs (json, yaml, etc.)

### Type Safety
- **Handle type assertion failures**: Always use "comma ok" idiom: `t, ok := i.(string)`
- **Start enums at one**: Use `iota + 1` so zero value is invalid/unknown
- **Avoid built-in names**: Don't shadow built-in identifiers (error, string, etc.)
- **Use raw string literals**: Use backticks for strings with quotes/backslashes

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