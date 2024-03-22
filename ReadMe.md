# CBAM API

![Successful XML File Upload](https://postimg.cc/6TZ0DyRm)

The CBAM API enables seamless integration of EU CBAM XML reports into your systems, transforming complex XML data into clear, actionable endpoints. Designed for developers and enterprises alike, this API simplifies the handling of carbon pricing data, making it accessible and actionable.

## Features

- **Simple Integration**: Easy setup allows for direct integration into existing systems.
- **Actionable Endpoints**: Convert complex XML data into clear and concise information.
- **EU CBAM Compliance**: Assists with compliance with the EU Carbon Border Adjustment Mechanism (CBAM).

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

The API server should now be running locally.

## Future Roadmap
- docs
- Implement user auth
- Add reportId validation
- Add reportId generation
- Add report modification support
- Enhance error handling and logging
- Improve API documentation and code comments
