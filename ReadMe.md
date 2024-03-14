# CBAM API Project

This project provides a RESTful API for managing quarterly reports using the CBAM (Carbon Border Adjustment Mechanism) framework.

## Installation

### 1. Install GoLang

GoLang is required for running this project. If you haven't already installed GoLang, follow these steps:

- Download the installer for your operating system from the [official GoLang website](https://golang.org/dl/).
- Follow the installation instructions provided.

### 2. Set Up Workspace

Create a workspace directory for your Go projects. By convention, this directory is often named `go` and is placed in your home directory. Within this directory, you can create subdirectories for your Go projects.

```bash
mkdir ~/go
```

### 3. Set Up GOPATH

Set the `GOPATH` environment variable to point to your workspace directory. You can add this to your shell profile (`~/.bashrc`, `~/.bash_profile`, `~/.zshrc`, etc.) to make it persistent.

```bash
export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin
```

### 4. Clone the Repository

Clone this repository into your Go workspace:

```bash
git clone https://github.com/yourusername/cbam-api.git ~/go/src/github.com/yourusername/cbam-api
```

### 5. Install Dependencies

Navigate to the project directory and install dependencies:

```bash
cd ~/go/src/github.com/jndunlap/cbam-api
go mod download
```

### 6. Run

run the project:

```bash
air
```

The API server should now be running locally.

### 7. Verify Installation

You can test the API by sending requests to `http://localhost:8080` using tools like cURL, Postman, or your preferred HTTP client.

## Future Roadmap

- Implement user authentication and authorization.
- Add support for exporting reports to various formats (XML, JSON, CSV).
- Enhance error handling and logging.
- Improve API documentation and code comments.
