# Burny CI/CD

## Flows

- When code is pushed to `main` or `dev` branches:
  - `api_scenario.yml` and `api_scenario.yml` will be triggered.
  - After corresponding scenario test passed, `api_deployment.yml` and `web_deployment.yml` will be triggered.
- When PR is created and including `/api` diff or `/web` diff:
  - `api_scenario.yml` and `api_scenario.yml` will be triggered.
- Each workflows can be triggered mannually via GitHub UI.

## Google Cloud Authentication

- For authentication with Google Cloud for deployemnt, [Workload Identity](https://cloud.google.com/blog/ja/products/devops-sre/infrastructure-as-code-with-terraform-and-identity-federation) is used. Relevant resources are defined in [/infra/modules/github](/infra/modules/github/).
