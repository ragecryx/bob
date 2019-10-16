# bob
A minimalistic CI/CD webhook service in Go. Fully compatible with [WebHawk](https://github.com/dbalaouras/webhawk)'s config and recipes.

<p align="center">
  <img height="250" src="https://raw.githubusercontent.com/ragecryx/bob/dev/assets/bob.png?raw=true" alt="Bob"/>
</p>

## F.A.Q

#### Can I use this to automatically build my project that is hosted in GitHub when changes happen in a branch?
Yes

#### Why it's named Bob?
Because it's a builder.

#### Can we build it?
Yes we can! :hammer: :construction_worker: :wrench: (Just follow the instructions below)

#### Why you made this tool since there are way better tools that do the job out there?
Because I can and because I wanted to get more familiar with Go. Worry not, I'll still :sparkling_heart: you even if you never use it.

## Build instructions
* Make sure you have [latest Go installed and configured](https://golang.org/doc/install)
* Clone the project in your workspace directory (what `$GOPATH` points at)
* `cd` in the project directory
* Use `make deps` to get all dependencies.
* Use `make` to build the executable (named `bob_builder`)

## License
This is under [MIT License](https://en.wikipedia.org/wiki/MIT_License). Check the `LICENSE` file in the project root directory for more info.
