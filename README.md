# Jira CLI Tracker

## Overview

The Jira CLI Tracker is a command-line application designed to help you manage and track your Jira tasks directly from the terminal. This tool allows you to view tasks assigned to you with specific labels, filter tasks by status, and more.

## Features

- **Task Listing**: View all tasks assigned to you with a specific label.
- **Status Filtering**: Filter tasks by status (Open, In Progress, Done).
- **Clear Console**: Clears the console screen for a clean output.
- **Environment Variables**: Uses environment variables for configuration (JIRA_URL, USERNAME, API_TOKEN, LABEL).

## Prerequisites

- Go 1.16 or later
- Jira account with appropriate permissions
- Environment variables set up (`.env` file or system environment variables)

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/mrashed98/jiraCliTracker.git
   cd jiraCliTracker
   ```

2. Install dependencies:
   ```sh
   go mod download
   ```

3. Set up your environment variables by creating a `.env` file in the root directory with the following content:
   ```env
   JIRA_URL=https://your-jira-instance.atlassian.net
   USERNAME=your-jira-username
   API_TOKEN=your-jira-api-token
   LABEL=your-label
   ```

## Usage

### Commands

- **List All Tasks**:
  ```sh
  go run main.go tasks
  ```

- **Filter Tasks by Status**:
  - Open Tasks:
    ```sh
    go run main.go open
    ```
  - In Progress Tasks:
    ```sh
    go run main.go inprogress
    ```
  - Done Tasks:
    ```sh
    go run main.go done
    ```

### Example

To list all tasks assigned to you with the label specified in your `.env` file:
```sh
go run main.go tasks
```

To list all tasks marked as "Done":
```sh
go run main.go done
```

## Configuration

The application uses environment variables for configuration. Ensure that the `.env` file is correctly set up with your Jira instance details, username, API token, and label.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or support, please contact [Mahmoud Rashed](mailto:mahmoudrashed2806@gmail.com).

---

Thank you for using the Jira CLI Tracker! We hope this tool helps you manage your Jira tasks more efficiently.