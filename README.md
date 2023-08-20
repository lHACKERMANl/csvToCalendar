# csvToCalendar

**csvToCalendar** is a tool developed in Go that facilitates the automatic loading of data from a CSV file into Google Calendar. It simplifies the process of adding events to Google Calendar based on data from a CSV file.

## Usage

To use **csvToCalendar**, follow these steps:

1. Prepare the `credentials.json` file:

   To interact with Google Calendar, you need to create an OAuth 2.0 client in the Google Cloud Console and upload the `credentials.json` file, which contains your credentials.

2. Prepare the `data.csv` CSV file:

   Your CSV file should have the following structure:

   ```plaintext
   Date,Subject,MSK Format,Start (MSK),IZH Format,Start IZH,Instructor,Record
   DD/MM/YYYY,Title,Type,HH:MM(begin),Location,HH:MM(end),speaker,link```
Note: Some columns like "MSK Format," "IZH Format," "Instructor," and "Record" can be left empty.

    go run main.go

## data.csv Signature

plaintext

Date,Subject,MSK Format,Start (MSK),IZH Format,Start IZH,Instructor,Record
DD/MM/YYYY,Title,Type,HH:MM(begin),Location,HH:MM(end),speaker,link

The signature describes the structure of the data.csv CSV file, where each row represents event information including date, subject, start format (MSK), start time (MSK), and other parameters.

Note: The signature provided here does not list all possible columns and values, only those provided in your example.

## `credentials.json` Signature

The `credentials.json` file is required for interacting with Google Calendar using the **csvToCalendar** tool. It contains the OAuth 2.0 client credentials needed to authorize access to your Google account. Here is an overview of its structure:

```json
{
  "installed": {
    "client_id": "YOUR_CLIENT_ID",
    "project_id": "YOUR_PROJECT_ID",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://accounts.google.com/o/oauth2/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_secret": "YOUR_CLIENT_SECRET",
    "redirect_uris": ["urn:ietf:wg:oauth:2.0:oob", "http://localhost"]
  }
}
```

Replace the following placeholders with your actual information:

    YOUR_CLIENT_ID: The client ID assigned when you create an OAuth 2.0 client in the Google Cloud Console.
    YOUR_PROJECT_ID: Your Google Cloud project ID.
    YOUR_CLIENT_SECRET: The client secret associated with your OAuth 2.0 client.

For detailed instructions on obtaining the credentials.json file and setting up OAuth 2.0 credentials, refer to the Google Calendar API documentation or related guides.
