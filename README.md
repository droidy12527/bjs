# BJS - Blazingly Fast JavaScript (Compiled JS)

## Overview

BJS is a powerful compiled language designed to bring blazing-fast performance to JavaScript development. Currently in its early stages, BJS provides a JavaScript-like syntax to ensure a smooth transition for JavaScript developers, with ambitious plans for future optimization and features.

## Current Features

* JavaScript-like syntax for familiar development experience
* Basic REPL interpreter for development and testing
* Simple code execution capabilities

## Installation

```bash
# Set debug mode
export BJS_DEBUG=true

# Start the BJS interpreter
go run .

```

## Additionally you can start the compiler mode by using the following command: (Experimental)
```bash
# Set Compile mode
export BJS_COMPILE_MODE=true

# Start the BJS compiler
go run .

```

## You can compile the codebase to native binary by using the following code
```bash
# Export as single file and add to compuer ENV for accessing from anywhere
go build -o bjs

```

## You can set the binary in your binary folder to access from anywhere by using this command
```bash
# Compile the binary by using this command
sudo ./compile.sh
# Move the BJS binary to bin folder
sudo mv bjs /usr/local/bin/ 
# Make it executable
sudo chmod +x /usr/local/bin/bjs
```

## Future Releases

### Compilation Benefits (Coming Soon)
* Native-like performance through Golang backend compilation
* Automatic optimization of code execution paths
* Static type checking for enhanced reliability
* Dead code elimination and tree shaking

### Multithreading Support (Planned)
* Seamless multithreading capabilities
* Automatic thread management
* Built-in concurrency patterns
* Efficient worker pool implementation

### JavaScript Compatibility (In Development)
* Full ES6+ feature support
* Seamless integration with existing JavaScript projects
* NPM package compatibility
* Advanced runtime optimizations

### Additional Planned Features
* Built-in package manager
* Development server with hot reloading
* Production build optimization
* Advanced configuration options
* IDE support and tooling
* Debugging tools and profiler

## Current State

BJS is currently in early development, focusing on providing a familiar syntax for JavaScript developers. The interpreter allows basic code execution while we work on implementing the advanced features that will make BJS truly blazingly fast.

## Roadmap

1. Q2 2025: Stable interpreter with basic JavaScript syntax support
2. Q3 2025: Initial compilation features and basic optimizations
3. Q4 2025: Multithreading support and advanced compilation features
4. Q1 2026: Full JavaScript compatibility and ecosystem integration

*Note: This project is in early development. Most features are planned for future releases. Current version provides basic JavaScript-like syntax interpretation.*