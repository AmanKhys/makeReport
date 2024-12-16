# Attendance Report Generator

This program helps you generate attendance reports for a session based on the list of members and other details like coordinators and batch names. It reads the list of members from a file, asks for the attendance, and outputs a well-formatted session report.

## Features

- Read members from a text file (`members.txt`).
- Allows input for the batch name.
- Allows users to select coordinators from the list of members.
- Allows input of attendance (present/absent) for each member.
- Sorts and formats the output for clarity.
- Optionally, input a date for the report (defaults to today's date).
- Shows a list of present and absent members.
- Prints session content from an external file (`content.txt`).

## Prerequisites

Before using the program, make sure you have the following files available in the same directory:

- `members.txt` – A text file containing the list of members, one per line.
- `content.txt` – A text file that contains content to be printed in the report (e.g., meeting agenda, session details).
  
## How to Use

### Linux
1. Make sure you have Go installed on your system.
2. Download the Linux executable from the project directory or build it from the source.
3. Run the program using the following command:
   ```bash
   ./makeReport
   ```

### Windows
1. Download the Windows executable (`makeReport.exe`) or build it from the source.
2. Run the executable on a Windows machine by double-clicking `makeReport.exe` or executing it from the command line:
   ```bash
   makeReport.exe
   ```

### Command-Line Flow

When you run the program, it will ask for the following inputs in order:

1. **Batch Name**: Enter the name of the batch (e.g., "Batch A").
2. **Date**: Enter today's date or leave it blank to use the current date.
3. **Select Coordinators**: You will see a list of members with numbers. Enter the number(s) of the coordinator(s) separated by spaces or press Enter when done.
4. **Attendance**: For each member, you'll be asked if they are present or absent. You can press Enter for "present" or type "y" for present and "n" for absent.

### Sample Interaction

```
Enter the batch name: Batch A
Enter today's date (or press enter to use the current date): 
Select coordinators from the list below (press Enter when done):
1. John
2. Mary
3. Alice
4. Bob

Enter the number of the coordinator (or press Enter to stop): 1
Enter the number of the coordinator (or press Enter to stop): 2

Is John present today?: y
Is Mary present today?: n
Is Alice present today?: y
Is Bob present today?: n

Session Report for 2024-12-16
------------------------------------------------
Batch Name: Batch A
Coordinators: John, Mary
------------------------------------------------
Meeting content: Today, we will discuss the upcoming project timeline and assign tasks.
We will also plan the next team's meeting.
------------------------------------------------
Present Members:
1. Alice
2. John

Absentees:
1. Bob
2. Mary
```

### Output

After entering all the necessary information, the program will output the following:
- **Batch Name**: The name of the batch (e.g., "Batch A").
- **Coordinators**: A list of coordinators you selected.
- **Content**: The content loaded from `content.txt`.
- **Present Members**: A numbered list of the members who are present.
- **Absentees**: A numbered list of the members who are absent. If there are no absentees, it will state "No absentees today."

## Building the Executable

If you'd like to build the executable from source, follow these steps:

1. Clone the repository or navigate to the project directory.
2. Install Go if it's not installed yet.
3. Use the following commands to build the Linux and Windows executables:

   ### For Linux:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o makeReport main.go
   ```

   ### For Windows:
   ```bash
   GOOS=windows GOARCH=amd64 go build -o makeReport.exe main.go
   ```

This will generate the respective executable files (`makeReport` for Linux and `makeReport.exe` for Windows).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

--- 

This `README.md` file provides a clear description of how to use the program, including installation instructions, expected output, and sample interaction, which should help any user interact with the program smoothly.
