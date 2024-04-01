# CBAM-API 

- **Centralize Emissions**: Gather supplier embedded emissions data for analysis.
- **Optimize Forecasts**: Align carbon obligation forecasts with real-time market activities.
- **Automate Reporting**: Streamline the collection, validation, and reporting of emissions data.
  
![cbam-apicbam-2](https://github.com/JNDUNLAP/cbam-api/assets/125301054/3ccfc631-f328-48fd-9d5f-b17b909e2e46)

## Introduction

CBAM-API is a open-source platform designed to streamline emissions management and reporting processes. It empowers users to centralize supplier embedded emissions for analysis, optimize carbon obligation forecasts, and automate data collection, validation, and reporting.

## Features

- **Integrated Database**: Utilize CBAM-API's built-in database for storing Quarterly CBAM Reports.
- **Modern Endpoints**: Convert EU's XML reports into dynamic REST endpoints for seamless data updates.
- **Custom Validation**: Validate entire reports against EU standards at the character level.


## Methodology

- The EU CBAM has intensive requirements.
- CBAM requirements are defined to the character level.
- CBAM-API centralizes report  requirements, apply these constraints, and returns validated JSON.

[Centralized Requirements](https://github.com/JNDUNLAP/cbam-api/tree/main/files/requirements)

1. Define Report Data Model with Constraints
2. Read XML into Memory
3. Read Constraints and Map to Custom Types
4. Apply Constraints
5. Convert Data to JSON


![Methodology](https://github.com/JNDUNLAP/cbam-api/assets/125301054/c6177062-d6f3-4b6d-b10d-7e9dba8f021d)

## Dependencies

- `github.com/cosmtrek/air`
- `github.com/joho/godotenv`
- `go.mongodb.org/mongo-driver`



## Getting Started

### Installation

- **Install Docker**: Refer to [Docker Documentation](https://docs.docker.com/engine/install/) for installation instructions.

- **Clone the Repository**:

    ```sh
    git clone https://github.com/jncbam_api/cbam-api.git
    ```

### Configuration

- **Environment Variables**:

    - Create a `.env` file.
    - Define the following variables:

    ```env
    APP_USER=some_username
    BUILD_TARGET=development
    # BUILD_TARGET=production
    MONGO_ROOT_USERNAME=your_admin_username
    MONGO_ROOT_PASSWORD=super_secret_password
    MONGO_HOST=mongodb
    MONGO_PORT=27017
    DB_NAME=database_name
    REPORTCOLLECTION=your_first_cbam_collection
    ```

### Usage

- **Build and Run Locally**:

    ```sh
    docker-compose up --build
    ```

