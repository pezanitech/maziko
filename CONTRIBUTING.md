# Contributing to Maziko

Thank you for your interest in contributing to Maziko! This document provides guidelines and instructions for contributing to this project.

## Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct. Please be respectful and considerate of others.

## How to Contribute

There are many ways to contribute to Maziko:

- Reporting bugs
- Suggesting enhancements
- Submitting pull requests
- Improving documentation
- Sharing feedback

## Submitting Issues

### Bug Reports

When submitting a bug report, please include:

1. **Title**: A clear, descriptive title
2. **Environment**: 
   - Maziko version
   - Go version
   - Node.js version
   - Operating system and version
   - Browser (if applicable)

3. **Steps to Reproduce**: Detailed steps to reproduce the issue
   - Example code if applicable
   - Error messages (if any)

4. **Expected Behavior**: What you expected to happen

5. **Actual Behavior**: What actually happened
   - Include screenshots or error logs if relevant

6. **Possible Solution**: If you have suggestions for fixing the bug

Example template:
```
## Bug Report

### Environment
- Maziko version: v0.1.0
- Go version: 1.24.2
- Node.js version: 18.12.0
- OS: Ubuntu 22.04
- Browser: Chrome 112.0.5615.49

### Steps to Reproduce
1. Create a new route at 'app/routes/test/'
2. Add only a page.tsx file without a handler.go
3. Run 'pnpm dev'

### Expected Behavior
The app should display an error message about missing handler.go

### Actual Behavior
The server crashes with the following error:
[error log here]

### Possible Solution
Add validation for missing handler files in the route generation process.
```

### Feature Requests

When submitting a feature request, please include:

1. **Title**: A clear, descriptive title
2. **Problem Statement**: The problem you're trying to solve
3. **Proposed Solution**: Your idea for implementing the feature
4. **Alternative Solutions**: Any alternatives you've considered
5. **Additional Context**: Any other information that might be helpful

Example template:
```
## Feature Request

### Problem Statement
Currently, there's no way to easily add middleware to specific routes.

### Proposed Solution
Add support for a middleware.go file within route directories that would contain middleware functions applied to that route.

### Alternative Solutions
- Global middleware configuration in maziko.json
- Explicit middleware attachment in handler.go files

### Additional Context
This would make it easier to implement route-specific authentication, logging, etc.
```

## Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Pull Request Guidelines

- Follow the existing code style and formatting
- Include tests for new features
- Update documentation as needed
- Keep pull requests focused on a single topic
- Reference any related issues

## Development Setup

Please refer to the README.md for instructions on setting up the development environment.

## License

By contributing to Maziko, you agree that your contributions will be licensed under the project's MIT license.