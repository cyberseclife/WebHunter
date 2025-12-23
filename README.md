# WebHunter

All-In-One Bug Bounty Tool Written In GoLang

cyberseclife@debian-vm ~/Projects/webhunter $ ./webhunter
``` _       __     __    __  __            __
| |     / /__  / /_  / / / /_  ______  / /____  _____ 
| | /| / / _ \/ __ \/ /_/ / / / / __ \/ __/ _ \/ ___/
| |/ |/ /  __/ /_/ / __  / /_/ / / / / /_/  __/ /
|__/|__/_\___/_.___/_/ /_/\__,_/_/ /_/\__/ _\___/_/
```

    High-Performance Security Assessment CLI

A modular CLI application for Recon, Analysis, and Exploitation.
Automates the security assessment pipeline.

Usage:
  webhunter [command]

Available Commands:
  analysis     Analyze targets for vulnerabilities
  completion   Generate the autocompletion script for the specified shell
  config       Manage configuration and scope files
  exploitation Run exploitation modules
  help         Help about any command
  recon        Perform reconnaissance and port scanning
  start        Start the master assessment workflow

Flags:
      --exclude string   File containing excluded targets (Out of Scope)
      --header string    Custom header to include in requests
  -h, --help             help for webhunter
      --include string   File containing allowed targets (Scope)
      --rate int         Rate limit in requests per second (default 5)

Use "webhunter [command] --help" for more information about a command.
