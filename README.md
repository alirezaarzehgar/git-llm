# git-llm
git-llm is just a simple git subcommand to integrate and use large language models during daily git usages.
features include fixing grammar of your commits or generating a complete commit from `git diff --cached` output!

Currently, we only support the Groq API.

# Installation
```
go install github.com/alirezaarzehgar/git-llm
```

# Getting help
```text
$ git llm -h
utilizing LLM for managing git projects

Usage:
  git-llm [command]

Available Commands:
  commitfix   fix grammar and structure of commit
  commitgen   generate commit message and open editor
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  setup       Set configuration

Flags:
  -h, --help     help for git-llm
  -t, --toggle   Help message for toggle

Use "git-llm [command] --help" for more information about a command.

$ git llm commitgen -h
generate commit message using configured LLM and
open configured EDITOR to change it. Then commit that message.
For example:

git llm commitgen

Usage:
  git-llm commitgen [flags]

Flags:
  -h, --help   help for commitgen

$ git llm commitfix -h
Error: required flag(s) "message" not set
Usage:
  git-llm commitfix [flags]

Flags:
  -h, --help             help for commitfix
  -m, --message string   Use the given message as the commit message
```

# Configuration
```yaml
# ~/.config/git-llm.yaml
editor: vim
grok_api_key: TOKEN
llm_model: llama-3.3-70b-versatile
```