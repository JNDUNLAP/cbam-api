# CBAM-API Documentation

## Introduction

CBAM-API is a versatile tool designed to streamline emissions management and reporting processes. It empowers users to centralize supplier embedded emissions for analysis, optimize carbon obligation forecasts, and automate data collection, validation, and reporting.

## Use Cases

- **Centralize Emissions**: Gather supplier embedded emissions data for comprehensive analysis.
- **Optimize Forecasts**: Align carbon obligation forecasts with real-time market activities.
- **Automate Reporting**: Streamline the collection, validation, and reporting of emissions data.

## Features

- **Integrated Database**: Utilize CBAM-API's built-in database for storing Quarterly CBAM Reports.
- **Modern Endpoints**: Convert EU's XML reports into dynamic REST endpoints for seamless data updates.
- **Custom Validation**: Validate entire reports against EU standards at the character level.

## Dependencies

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
    APP_USER=appuser
    BUILD_TARGET=development
    # BUILD_TARGET=production
    MONGO_ROOT_USERNAME=admin
    MONGO_ROOT_PASSWORD=secret
    MONGO_HOST=mongodb
    MONGO_PORT=27017
    DB_NAME=mydatabase
    REPORTCOLLECTION=test_report
    ```

### Usage

- **Build and Run**:

    Execute the following command to build and run the application using Docker Compose:

    ```sh
    docker-compose up --build
    ```

