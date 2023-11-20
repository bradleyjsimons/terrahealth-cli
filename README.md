# TerraHealth CLI

## Overview

TerraHealth CLI is an open-source command-line tool designed for monitoring and reporting the health of AWS resources. It aims to assist DevOps teams and individuals in efficiently managing their AWS infrastructure. The tool is in active development, starting with basic AWS functionalities and planning to expand into more comprehensive features.

## Features

### Current Features

- **AWS Connection**: Establishes a connection to AWS and interacts with various services.
- **EC2 Instance Monitoring**: Lists the current EC2 instances, providing a snapshot of their statuses.

### Planned Features

- **Health Monitoring**: Extend functionality to check the health status of various AWS services, including EC2, RDS, and Lambda.
- **Change Reporting**: Implement change detection features to track and report modifications in the infrastructure.
- **Compliance Checks**: Add compliance checking to ensure that the infrastructure adheres to specific standards and best practices.
- **CLI-to-Web UI Bridging**: Develop a basic web interface for more detailed reports, accessible directly from the CLI tool.
- **Scalability and Extensibility**: Design the tool to easily accommodate new AWS services and features in the future.

## Getting Started

### Prerequisites

- Go (version 1.x or later)
- AWS account and AWS CLI configured with appropriate credentials

### Installation

(Currently, the installation process involves cloning the repository and building the tool locally. Instructions will be updated as the project progresses.)

1. **Clone the Repository**:

   git clone https://github.com/yourusername/TerraHealth-CLI.git

2. **Build the Tool**:

   cd TerraHealth-CLI
   go build cmd/main.go

### Usage

Run the TerraHealth CLI with the desired command. For example:

./main check-aws

This command checks and lists the AWS EC2 instances in your configured AWS account.

## Development Status

As of now, TerraHealth CLI:

- Sets up the basic structure and connection to AWS.
- Implements functionality to list EC2 instances.
- Includes tests for the implemented features.

Future development will focus on expanding the capabilities as outlined in the Planned Features section.

## Contributing

We welcome contributions from the community! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature.
3. Commit your changes.
4. Push to the branch and open a pull request.

## License

[MIT License](LICENSE) - see the `LICENSE` file for details.
