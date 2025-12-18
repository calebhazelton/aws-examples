# aws-examples
This codebase contains all AWS examples created through the 'AWS Certified Solution Architect Associate' certification study courses.  This repository contains...

- devcontainer setup to copy dependencies in a separate workspace
- bash scripts demonstrating AWS CLI use for a particular service
- go packages demonstrating AWS SDK use for a particular service
- IAC 

### github codespace setup
- Configure all required github codespace secrets listed in `.env.example` to match your environment.
- Create a github codespace, and build off the included `.devcontainer`, all dependencies should be installed upon building the container.

### local/other setup
- Install the aws cli on your machine https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html.
- Configure the environment variables listed in `.env.example` based on your environment.