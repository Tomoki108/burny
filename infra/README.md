# Burny Infrastructure


# Prerequisites

## tools
- Install tfenv and sepcified terraform version.
    ```shell
    brew install tfenv
    tfenv install 1.10.5
    tfenv use 1.10.5
    ```

- Install [gcloud CLI](https://cloud.google.com/sdk/auth_success?hl=ja) and login.
    ```
    gcloud auth application-default login
    ```

## each environment setting (currently only dev)
- Create `terraform.tfvars` from `terraform.tfvars.sample`. Secret values can be referenced at Secret Manager console of burny-{env} project.
- Init module by `terraform init` and `terraform apply`.
