# Go Tutorials

This folder contains my walkthroughs of several official [Go tutorials](https://go.dev/doc/tutorial), focused on modern backend development, modules, and advanced Go features. Each subfolder corresponds to a complete hands-on example or concept.

---

## 📚 Tutorials Completed

- [Create a module](https://go.dev/doc/tutorial/create-module) → [Module Folder](https://github.com/dhu2022-dev/GO_TUTORIAL/tree/main/Tutorials/Module): Initialize and build a CLI app with `go mod`
- [Develop a web service](https://go.dev/doc/tutorial/web-service-gin) → [Web_Service_Gin Folder](https://github.com/dhu2022-dev/GO_TUTORIAL/tree/main/Tutorials/Web_Service_Gin): Build a RESTful API using the Gin framework
- [Access a database](https://go.dev/doc/tutorial/database-access) → [Data_Access Folder](https://github.com/dhu2022-dev/GO_TUTORIAL/tree/main/Tutorials/Data_Access): Connect and query a **MySQL** database using Go and environment configs
- [Work with multiple modules](https://go.dev/doc/tutorial/workspaces) → [Module_Multiple Folder](https://github.com/dhu2022-dev/GO_TUTORIAL/tree/main/Tutorials/Module_Multiple): Learn to use `go.work` for multi-module workspaces
- [Fuzzing in Go](https://go.dev/doc/tutorial/fuzz) → [Fuzz Folder](https://github.com/dhu2022-dev/GO_TUTORIAL/tree/main/Tutorials/Fuzz): Use Go’s built-in fuzzing tools for input testing
- [Generics](https://go.dev/doc/tutorial/generics) → [Generics Folder](https://github.com/dhu2022-dev/GO_TUTORIAL/tree/main/Tutorials/Generics): Understand and apply generics using type parameters
- [Vulnerability scanning](https://go.dev/doc/tutorial/security) → [Vulnerabilities Folder](https://github.com/dhu2022-dev/GO_TUTORIAL/tree/main/Tutorials/Vulnerabilities): Detect and patch security issues with `govulncheck`

---

## 🗂️ Folder Structure

```
Tutorials/
├── Module/                    # Single-module CLI app
├── Web_Service_Gin/           # REST API using Gin
├── Data_Access/               # MySQL database + .env + SQL setup
├── Module_Multiple/           # Multi-module workspace (hello/, example/, go.work)
├── Fuzz/                      # Fuzz testing
├── Generics/                  # Go generics practice
├── Vulnerabilities/           # Go security tools and scanning
```

---

## 🚀 Running the Code

Make sure [Go is installed](https://go.dev/doc/install).

1. Navigate into a tutorial folder:
   ```bash
   cd Tutorials/Data_Access
   go run .
   ```

2. For workspace-based projects like `Module_Multiple`:
   ```bash
   cd Tutorials/Module_Multiple
   go run .
   ```

---

## 📝 Notes

- All code follows the official Go documentation.
- `.env` files and `.sql` scripts are included where relevant.
- The purpose of this repo is to deepen practical Go understanding through hands-on coding.
