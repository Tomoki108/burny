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

## bucket
- Create GCS bucket named "burny-tfstate" mannually for tfstate backend. 

## each environment setting (dev, prod)
- Create `terraform.tfvars` from `terraform.tfvars.sample`.
- Init module by `terraform init` and `terraform apply`.
- Put secret values mannually at Secret Manager console. Values must be the same as defined by terraform.tfvars.