// Project Structure
// https://github.com/golang-standards/project-layout - the standard to follow, but there isn't an correct way. Best to be consitent
// /cmd - the main sources files. The main.go of a foo application should live in /cmd/foo/main.go
// /internal - private code that we don't others importing
// /pkg - public code
// /test - additonal external test and test data
// /configs - configuration files
// /docs - Design and user docs
// /examples
// /api - API contract files (Swagger, Protocol Buffers, etc.)
// /web - web application-specific assets (static files, etc.)
// /build - packaging and continuous integration (CI) files
// /scripts - script for analysis, installation, and so on
// /vendor application dependencies (for example, go modules dependecies)
