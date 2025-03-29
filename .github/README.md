# Burny CI/CD

This document describes the CI/CD workflows for the Burny project.

## Flows

- When code is pushed to `main` or `dev` branches:
  - `api_scenario.yml` and `api_scenario.yml` will be triggered.
  - After corresponding scenario test passed, `api_deployment.yml` and `web_deployment.yml` will be triggered.
- When PR is created and including `/api` diff or `/web` diff:
  - `api_scenario.yml` and `api_scenario.yml` will be triggered.
- Each workflows can be triggered mannually via GitHub UI.
