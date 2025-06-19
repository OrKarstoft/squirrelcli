# squirrelcli 🐿️

**squirrelcli** is a command-line tool written in Go that automates the creation of JIRA Story issues on a board named `BOSS`. It assigns the new Story to the authenticated user, adds it to the active sprint, and immediately transitions it to "Done". 🚀

## Features ✨

- 📝 Creates a new JIRA Story with a summary and description.
- 👤 Assigns the Story to the authenticated user.
- 📋 Finds the `BOSS` board and the active sprint automatically.
- 🏃 Adds the Story to the current active sprint.
- ✅ Transitions the Story directly to "Done".

## Prerequisites ⚙️

- 🌐 Access to a JIRA instance via API.
- 🔑 API credentials with permissions to create and transition issues on the `BOSS` board.
- The following environment variables set:
  - `SQUIRREL_JIRA_URL`
  - `SQUIRREL_JIRA_APIKEY`
  - `SQUIRREL_JIRA_USERNAME`

## Installation 🛠️

```sh
go build -o squirrelcli
```

## Usage 🚦
```sh
./squirrelcli "Story summary" "Story description"
```

- The first argument is the Story summary (required).
- The second argument is the Story description (optional).

## License 📄
[MIT](https://github.com/OrKarstoft/squirrelcli/tree/LICENSE)
