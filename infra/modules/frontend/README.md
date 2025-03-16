# Frontend Infrastructure Module

This module manages the frontend infrastructure components for the Burny application in Google Cloud Platform.

## Components

- **Cloud Storage**: Static website hosting bucket
- **Load Balancer**: HTTPS load balancer for the static website
- **Cloud CDN**: Content delivery network for static assets
- **SSL Certificate**: Managed SSL certificate for secure connections

## Features

- Automated SSL certificate management
- CDN-enabled static website hosting
- HTTP to HTTPS redirection
- SPA (Single Page Application) routing support

## Module Configuration

For detailed configuration options, please refer to:

- `variables.tf` - For input variables and their descriptions
- `main.tf` - For the main module implementation
- `output.tf` - For module outputs
