# Timetravel CLI 🚗⚡

Welcome to the `timetravel` CLI! Ever wanted to dial back the clock and revisit moments in history? Or perhaps shoot forward into the future? Inspired by the classic "Back to the Future" films, `timetravel` allows you to (figuratively) hop into your DeLorean, punch in a date and time, and embark on a fantastic voyage through the temporal realm! (1.21 gigawatts not required.)

## Table of Contents
- [Timetravel CLI 🚗⚡](#timetravel-cli-)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Building the CLI](#building-the-cli)
  - [Usage](#usage)
    - [Traveling through time](#traveling-through-time)
    - [Feedback Survey](#feedback-survey)
  - [Contact](#contact)
  
## Getting Started

### Prerequisites

To build and run the `timetravel` CLI, you will need:

- [Go](https://golang.org/dl/) (version 1.17 or later)

### Building the CLI

1. Clone the repository:
```bash
git clone https://github.com/marianina8/timetravel.git
```

2. Build the CLI:
   * For MacOS and Linux:
```bash
go build -o timetravel
```
 * For Windows:
```bash
go build -o timetravel.exe
```
Now, you have a timetravel binary (or timetravel.exe on Windows) ready for some temporal adventures!

## Usage
### Traveling through time

To embark on a journey:
```bash
./timetravel to 070417761200
```

For customized outputs, use the -o flag:
```bash
./timetravel to 070417761200 -o=dashboard
```

Valid formats for `-o, --output` include: text (default), json, yaml, dashboard.

### Feedback Survey
Share your time-traveling experience:
```bash
./timetravel feedback
```
For a non-interactive mode, use the --no-input flag.

## Contact

For questions, suggestions, or temporal anomalies you've spotted, reach out to me at [mmontagnino@gmail.com](mailto:mmontagnino@gmail.com)!
