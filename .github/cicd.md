# Burny CI/CD

## Flows

- When code is pushed to `main` or `dev` branches:
  - `api_scenario.yml` and `web_scenario.yml` will be triggered.
  - After corresponding scenario test passed, `api_deployment.yml` or `web_deployment.yml` will be triggered.
- When PR is created and including `/api` or `/web` diff:
  - `api_scenario.yml` or `web_scenario.yml` will be triggered.
- Each workflows can be triggered mannually via GitHub UI.

## Google Cloud Authentication

- For Google Cloud authentication for deployment, [Workload Identity](https://cloud.google.com/blog/ja/products/devops-sre/infrastructure-as-code-with-terraform-and-identity-federation) is used. Relevant resources are defined in [/infra/modules/github](/infra/modules/github/).
