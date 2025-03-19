# Backend Infrastructure Module

This module manages the backend infrastructure components for the Burny application in Google Cloud Platform.

## Components

- **Cloud Run Service**: Deploys and manages the backend API service
- **Cloud SQL**: PostgreSQL database instance
- **Artifact Registry**: Docker image repository for Cloud Run
- **IAM**: Service account and permissions for Cloud Run
- **Secret Manager**: Manages database credentials and other secrets

## Features

- Automatic HTTPS configuration
- Database instance with secure configuration
- Integration with GitHub Actions for deployment
- Secret management for sensitive data

## Module Configuration

For detailed configuration options, please refer to:

- `variables.tf`
- `main.tf`
