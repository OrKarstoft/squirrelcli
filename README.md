# squirrelcli 🐿️

**squirrelcli** is a command-line tool written in Go that automates the creation of JIRA Story issues. It assigns the new Story to the authenticated user, adds it to the active sprint, and immediately transitions it to "Done". 🚀

## Features ✨

- 📝 Creates a new JIRA Story with a summary and description.
- 👤 Assigns the Story to the authenticated user.
- 📋 Finds the project board and the active sprint automatically.
- 🏃 Adds the Story to the current active sprint.
- ✅ Transitions the Story directly to "Done".

## Prerequisites ⚙️

- 🌐 Access to a JIRA instance via API.
- 🔑 API credentials with permissions to create and transition issues on your JIRA project.
- The application will prompt for configuration on first run and save it to `~/.config/squirrelcli/config.json`

## Installation 🛠️

```sh
go build -o squirrelcli
```

## Setup 🔧

The application will automatically prompt you for your JIRA configuration on first run.

1. **First run**: Simply run the application to start the interactive setup:

   ```sh
   ./squirrelcli "test" "test"
   ```

   You'll be prompted to enter:

   - JIRA URL (e.g., https://your-domain.atlassian.net)
   - JIRA API Key (generate one [here](https://id.atlassian.com/manage-profile/security/api-tokens))
   - JIRA Username
   - Project Key (e.g., BOSS)

2. **Configuration is saved**: Your settings are automatically saved to `~/.config/squirrelcli/config.json`

3. **Ready to use**: Subsequent runs will use your saved configuration

## Usage 🚦

```sh
./squirrelcli "Story summary" "Story description"
```

- The first argument is the Story summary (required).
- The second argument is the Story description (optional).

## License 📄

[MIT](https://github.com/OrKarstoft/squirrelcli/tree/LICENSE)
