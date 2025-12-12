# aws-examples
This codebase contains all AWS examples created through the 'AWS Certified Solution Architect Associate' certification study courses.  This repository contains the devcontainer setup and shell scripts to reproduce these examples in your own AWS environment, each aws sub directory contains CLI commands to configure that particular cloud service. 

### github codespace setup
- Configure all required github codespace secrets listed in `.env.example` to match your environment.
- Create a github codespace, and build off the included `.devcontainer`, all dependencies should be installed upon building the container.

### local/other setup
- Install the aws cli on your machine https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html.
- Configure the environment variables listed in `.env.example` based on your environment.