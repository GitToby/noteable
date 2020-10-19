# CLI Note taker
noteable is a cli note tool which you can set up to take notes and share via email.

A help description is available via `notable --help`

```
Take, manage and share notes powerfully via cli.

Usage:
  noteable [command]

Available Commands:
  config      Set and retrieve config for notable
  help        Help about any command
  history     Get the history of notes you've taken
  search      Search all your notes for a string
  share       A brief description of your command
  take        Take a note and save it

Flags:
      --config string   config file path to use (default "C:\\Users\\Toby\\.noteable\\config.yml")
  -h, --help            help for noteable
  -v, --version         version for noteable

Use "noteable [command] --help" for more information about a command.
```

# Configuration
Config is driven by a .yml file (`~/.noteable/config.yml` by default, generated on first run) a description of each field is below.

- `database_path`: The default path to look for the database (default: `~/.noteable/noteable.db`)
- `date_format`: The format to use for printing dates, used in search, history. Uses [go date formatting](https://yourbasic.org/golang/format-parse-string-time-date-example/) (default: "Mon 02 Jan 2006 15:04:05")
- `share`:
    - `methods`: The list of methods to run when calling \"noteable share\". This can be empty if you don't want to share anything. (default: \[email\*\])
    - `schedule`: Defines the day of the week to share notes with when taking a note on that day (default: 5) (sunday=0, so 5=friday)
    - `date_format`: The format to use for printing dates, used in share. Uses [go date formatting](https://yourbasic.org/golang/format-parse-string-time-date-example/) (default: "Mon 02 Jan 03:04 pm")
    - `email`
        - `to`: The email addresses to send out notifications to when calling \"notable share\" with 'email' in the methods list (default: "")
        - `cc`: The email addresses cc notifications to when calling \"notable share\" with 'email' in the methods list (default: "")
        - `subject`: The subject line to add to an email when forming the draft.(default: "Notes from noteable!")

