# CBAM API

Directly plug EU CBAM XML reports into your systems. Turn complex data into clear, actionable endpoints.

## Installation

#### Installing Go

[Install Go](https://go.dev/doc/install)

### 4. Clone the Repository

Clone this repository into your Go workspace:

```bash
git clone https://github.com/jndunlap/cbam-api.git 
```

### 5. Install Dependencies

Navigate to the project directory and install dependencies:

```bash
cd ~/go/src/github.com/jndunlap/cbam-api
go mod download
```

### 6. Run

Run the project:

```bash
air
```

The API server should now be running locally.

### 7. Verify Installation

You can test the API by sending requests to `http://localhost:8080` using tools like cURL, Postman, or your preferred HTTP client.

## Future Roadmap
- Something you suggest 
- Implement user auth
- Add reportId validation
- Add reportId generation
- Add report modification support
- Enhance error handling and logging
- Improve API documentation and code comments
