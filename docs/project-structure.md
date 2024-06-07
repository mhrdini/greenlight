# Project Structure

```bash
greenlight
|
+-- bin         # compiled application binaries, ready to be deployed to production server
|
+-- cmd
|    |
|    +-- api    # application-specific code for Greenlight API: running the server, processing HTTP requests, managing authentication
|
+-- internal    # ancillary, non-application-specific, potentially reusable code: interacting with our database, doing data validation, sending emails, etc.
|
+-- migrations  # SQL migrations for our DB
|
+-- remote      # config files and setup scripts for production server
|
+-- go.mod      # project dependencies, versions, and module path
|
+-- Makefile    # recipes for automating common tasks: auditing Go code, building binaries, executing database migrations
```
