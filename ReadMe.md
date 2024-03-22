# CBAM API
```bash
CBAM API Routes
---------------------------------------------

NationalCompetentAuth
GET /reports/{id}/national-competent-auth

Signatures
GET /reports/{id}/signatures

GoodsEmissions
GET /reports/{id}/imported-goods/emissions

Declarant
GET /reports/{id}/declarent

Importer
GET /reports/{id}/importer

ImportedGoods
GET /reports/{id}/imported-goods

Documents
GET /reports/{id}/imported-goods/{goodId}/supporting-documents

Files
POST /api/upload

QuarterlyReports
GET /reports
GET /reports/{id}
POST /reports/create
DELETE /reports/delete/{id}

2024/03/22 17:07:32 CBAM API is Running on PORT 8080
2024/03/22 17:08:03 [POST]: [200], Duration: [1.113584ms]  -  /api/upload
```

## Features

- **REST Endpoints**: Convert complex CBAM XML data into REST endpoints.
- **(BYOD) Bring Your Own Database**: Easy setup allows for direct integration into existing systems.

## Prerequisites

Before you begin, ensure you have the following installed:

[Install Go](https://go.dev/doc/install)

## Clone the Repository

Clone this repository into your Go workspace:

```bash
git clone https://github.com/jndunlap/cbam-api.git 
```

## Install Dependencies

Navigate to the project directory and install dependencies:

```bash
cd ~/go/src/github.com/jndunlap/cbam-api
go mod download
```

## Run

Run the project:

```bash
air
```
