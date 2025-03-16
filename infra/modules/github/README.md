# GitHub Infrastructure Module

This module manages GitHub Actions integration with Google Cloud Platform for the Burny application.

## Components

- **Service Account**: Dedicated service account for GitHub Actions
- **Workload Identity Pool**: Federation with GitHub Actions OIDC
- **IAM Bindings**: Required permissions for CI/CD workflows

## Features

- Secure authentication between GitHub Actions and GCP
- Workload identity federation setup
- Least privilege access configuration
- Support for artifact registry and Cloud Run deployments

## Module Configuration

For detailed configuration options, please refer to:

- `variables.tf` - For input variables and their descriptions
- `main.tf` - For the main module implementation
- `outputs.tf` - For module outputs
