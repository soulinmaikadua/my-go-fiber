# MY-GO-FIBER


Initial Setup

```bash
go mod init github.com/your/repo
```

1. Close this repository:

```bash
git clone https://github.com/soulinmaikadua/my-go-fiber.git
```

2. Navigate to the project directory:

```bash
cd my-go-fiber
```

### Environment Variables

This project uses environment variables for configuration. Before running the application, make sure to create a `.env` file in the root directory and define the following variables:

```bash
# Example .env file

# Database configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=username
DB_PASS=password
DB_NAME=dbname

# Other configuration variables
# Add any other environment variables your application needs here
```

Make sure to replace the placeholder values with your actual configuration.

### Build

To build the Go application, run the following command:

```bash
go build -o <output-file-name>
```

Replace <output-file-name> with the desired name of the executable file.

### Run

After building the application, you can run it using the following command:

```bash
./<output-file-name>
```

Replace `<output-file-name>` with the name of the executable file generated during the build process.
