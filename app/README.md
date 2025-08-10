# Google Slides API Playground

This is a sample Go application that demonstrates how to use the Google Slides API and Google Drive API to create and manage presentations.

## Features

- Create a new Google Slides presentation.
- Search for an existing presentation by title.
- Add a new slide to a presentation.

## Setup

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.21 or later)
- A Google Cloud Platform project with the Google Slides API and Google Drive API enabled.

### Configuration

1.  **Enable APIs**
    - Go to the [Google Cloud Console](https://console.cloud.google.com/).
    - Create a new project or select an existing one.
    - Enable the "Google Slides API" and "Google Drive API".

2.  **Create Credentials**
    - Go to the "APIs & Services" > "Credentials" page.
    - Click "Create Credentials" and choose "OAuth client ID".
    - Select "Desktop app" for the application type.
    - Download the JSON file.
    - Rename the downloaded file to `credentials.json` and place it in the root of this project.

3.  **Install Dependencies**
    - Open your terminal in the project root and run:
      ```sh
      go mod tidy
      ```

## Running the Application

1.  Run the application from your terminal:
    ```sh
    go run quickstart.go
    ```

2.  The first time you run the application, you will be prompted to authorize access:
    - A URL will be printed to your console.
    - Copy and paste this URL into your web browser.
    - Log in to your Google account and grant the requested permissions.
    - A code will be displayed in your browser. Copy this code.
    - Paste the code back into your terminal where the application is running and press Enter.

3.  A `token.json` file will be created in the project root. This file stores your authorization tokens. If you change the API scopes in the code, you will need to delete this file and re-authorize.

## How it Works

When you run the application, it will:
1.  Search for a Google Slides presentation named "New Presentation Title" in your Google Drive.
2.  If found, it will use that presentation.
3.  If not found, it will create a new presentation with that title.
4.  It will then add a new slide with a "Title and Two Columns" layout to the presentation.
