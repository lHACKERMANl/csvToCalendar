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
   DD/MM/YYYY,Title,Type,HH:MM(begin),Location,HH:MM(end),speaker,link

Note: Some columns like "MSK Format," "IZH Format," "Instructor," and "Record" can be left empty.

    Run the application:

    In the command line, navigate to the directory containing the application and execute the following command:

    sh

    go run main.go

    The application will prompt you to authorize access to Google Calendar and then load events from the data.csv file into your calendar.

data.csv Signature

plaintext

Date,Subject,MSK Format,Start (MSK),IZH Format,Start IZH,Instructor,Record
DD/MM/YYYY,Title,Type,HH:MM(begin),Location,HH:MM(end),speaker,link

The signature describes the structure of the data.csv CSV file, where each row represents event information including date, subject, start format (MSK), start time (MSK), and other parameters.

Note: The signature provided here does not list all possible columns and values, only those provided in your example.

