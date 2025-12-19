# AWS Example Usage
This codebase contains a variety of examples, demonstrating different ways to configure cloud infrastructure using Amazon Web Services.  This repository covers the three main ways to programatically interact with AWS services, via...
1. The AWS Command Line Interface (CLI)
2. The AWS Software Development Kit (SDK)
3. Infrastructure as Code (IAC) tools
    * These examples utilize Terraform and HCL

The contents of this repository include...
- The devcontainer setup to copy dependencies in a separate workspace
- The bash scripts demonstrating AWS CLI use for a particular service
- The Go packages demonstrating AWS SDK use for a particular service
- IAC

## Github codespace setup
- Configure all required github codespace secrets listed in `.env.example` to match your environment.
- Create a github codespace, and build off the included `.devcontainer`, all dependencies should be installed upon building the container.

## Local/other setup
### For all demonstrations
- Configure the environment variables listed in `.env.example` based on your environment.
    - This can be done simply by writing to the terminal, eg. `export AWS_REGION="us-east-1"`

### For AWS CLI demonstrations
- Install the aws cli on your machine https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html.

### For AWS SDK demonstrations
- Install Go (latest stable from https://go.dev/dl/) and ensure it’s on your PATH.

## How to run
### For all demonstrations
- Ensure AWS credentials are available (e.g., `AWS_PROFILE`/`AWS_REGION`) and environment variables from `.env.example` are set.
- Verify AWS access:
  - `aws sts get-caller-identity`

### For AWS CLI demonstrations
- Get a list of available bash scripts for a service.
  - `ls <service>/bash-scripts`.  Example usage: `ls s3/bash-scripts`
- Run selected bash script for input parameters.  Example `./s3/bash-scripts/create-bucket`

### For AWS SDK demonstrations
- Run the Go module to see available commands (from the repo root):
  - `go run .`
- Command example, Create Bucket
  - `go run . create --bucket <bucket_name> --region <region_name>`