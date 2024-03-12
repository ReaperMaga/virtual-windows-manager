<h1><img align="center" height="48" src="public/images/logo.svg"> Virtual Windows Manager</h3>

Virtual Windows Manager is a powerful tool designed to streamline the process of creating and managing Windows environments using Docker containers. Whether you're a developer, tester, or system administrator, Virtual Windows Manager simplifies the setup and maintenance of Windows environments, allowing you to focus on your tasks without worrying about compatibility issues or complex configurations.

## Features

- **Easy Setup**: Virtual Windows Manager makes it simple to set up and configure Windows environments using Docker containers.

- **Versatility**: Create multiple isolated Windows environments for various purposes, such as development, testing, or demonstration.

- **Efficiency**: Save time and resources by using lightweight Docker containers instead of traditional virtual machines.

- **Portability**: Docker containers ensure consistency across different platforms, making it easy to share environments with team members or deploy them to different machines.

- **Scalability**: Scale your Windows environments effortlessly by leveraging Docker's scalability features.

- **Customization**: Tailor your environments to specific requirements by installing custom software, configuring network settings, and more.

## Installation

To get started with Virtual Windows Manager, follow these simple steps:

1. **Install Docker**: Ensure that Docker is installed on your system. You can download and install Docker Desktop from [https://www.docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop).

2. **Clone the Repository**: Clone the Virtual Windows Manager repository to your local machine using Git:

    ```bash
    git clone https://github.com/reapermaga/virtual-windows-manager.git
    ```

3. **Run**: Navigate to the cloned directory and use docker compose to start your app.

    ```bash
    cd virtual-windows-manager
    docker compose up -d
    ```

## Development

This project uses go `1.21` <br />
First of all download all necessary dependencies with `go mod download`

To start working you need to fill up the `.env` file before starting.

Now to run the dev mode, just simple use `air` in the root directory

## Contributing

We welcome contributions from the community to enhance Virtual Windows Manager. If you encounter any bugs, have suggestions for new features, or would like to contribute code, please feel free to open an issue or submit a pull request on our [GitHub repository](https://github.com/reapermaga/virtual-windows-manager).

## License

Virtual Windows Manager is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Support

For any questions, feedback, or support requests, please contact me at [reapermaga@gmail.com](mailto:reapermaga@gmail.com).

---

Virtual Windows Manager - © 2024 ReaperMaga. All Rights Reserved.