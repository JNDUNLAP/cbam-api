# CBAM API

- Leverage MongoDB for CBAM Reports
- Turn archaic CBAM XML into modern REST endpoints.
- Custom CBAM types validate report structure at the character level.

## Dependencies
- `github.com/joho/godotenv`
- `go.mongodb.org/mongo-driver`

## Installation

[Install Go](https://go.dev/doc/install)

## Clone the Repository

Clone this repository

```bash
git clone https://github.com/jndunlap/cbam-api.git 
```

## Install Dependencies

- Navigate to the project directory
- Install dependencies:

```bash
cd ~/go/src/github.com/jndunlap/cbam-api
go mod download
```

## Environment Variables

- Create .env file
- Set values for the following varaibles
```bash
PORT=8888
MONGODB_URI=your_mongodb_uri
DB_NAME=your_database_name
```
## Run

```bash
air
```

## This program is incomplete - use at your own risk. 

### Testing Data Model Against EU Standards

- **Data Model**: PASSED

### CBAM API Routes

- **Importer**: `GET /reports/{id}/importer`
- **ImportedGoods**: `GET /reports/{id}/imported-goods`
- **Files**: `POST /api/upload`
- **NationalCompetentAuth**: `GET /reports/{id}/national-competent-auth`
- **Declarant**: `GET /reports/{id}/declarent`
- **Signatures**: `GET /reports/{id}/signatures`
- **GoodsEmissions**: `GET /reports/{id}/imported-goods/emissions`
- **Documents**: `GET /reports/{id}/imported-goods/{goodId}/supporting-documents`
- **QuarterlyReports**:
  - `GET /reports`
  - `GET /reports/{id}`
  - `POST /reports/create`
  - `DELETE /reports/delete/{id}`

`2024/03/23 19:02:56 CBAM API is Running on PORT 8080`

