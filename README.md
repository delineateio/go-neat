[![PRs Welcome][pr-welcome-shield]][pr-welcome-url]
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <img alt="delineate.io" src="https://github.com/delineateio/.github/blob/master/assets/logo.png?raw=true" height="75" />
  <h2 align="center">delineate.io</h2>
  <p align="center">portray or describe (something) precisely.</p>

  <h3 align="center">neat</h3>

  <p align="center">
    Small CLI for opinionated developer workflow
    <br />
    <a href="https://github.com/delineateio/oss-template"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/delineateio/oss-template/issues">Report Bug</a>
    ·
    <a href="https://github.com/delineateio/oss-template/issues">Request Feature</a>
  </p>
</p>

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [About The Project](#about-the-project)
- [Built With](#built-with)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [CLI Usage](#cli-usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

<!-- ABOUT THE PROJECT -->
## About The Project

An opinionated CLI that enable repeatable sets of commands and common tasks be automated locally while developing software.

## Built With

Further logos can be inserted to highlight the specific technologies used to create the solution from [here](https://github.com/Ileriayo/markdown-badges).

| Syntax | Description |
| --- | ----------- |
| ![pre-commit](https://img.shields.io/badge/precommit-%235835CC.svg?style=for-the-badge&logo=precommit&logoColor=white) | Pre-commit `git` hooks that perform checks before pushes|
| ![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white) | Source control management platform  |
| ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) | Containerise applications and provide local environment |

A series of Go packages have been used to build up `neat`:

| Syntax | Description |
| --- | ----------- |
|[cobra](github.com/spf13/cobra)| Provides the CLI framework for `neat` |
|[survey](github.com/AlecAivazis/survey)| Provides ability to capture input |
|[color](github.com/fatih/colo)| Enables console coloured output |
|[go-git](https://github.com/go-git/go-git)| Fully functioning package for `git`|
|[viper](https://github.com/spf13/viper)|  Configuration package for managing config |
|[zerolog](https://github.com/rs/zerolog)| Simple high performance logging framework |
|[lumberjack](https://github.com/natefinch/lumberjack)| Rolling log file management compatible with `zerolog`|
|[orderedmap](https://github.com/elliotchance/orderedmap)| Provides missing data structure for ordered maps|

<!-- GETTING STARTED -->
## Getting Started

This repo follows the principle of minimal manual setup of the local development environment, and utilises [devcontainer](https://containers.dev/) in [vscode](https://code.visualstudio.com/).

In addition a [taskfile](https://taskfile.dev) provides commonly required commands, these can be listed using `task --list`.

## Configuration

The global config file is expected to be located at `~/.config/.neat.yaml`

```yml
automation: false
log:
  level: debug # info | warn | error
  sinks: # file | console
    - file
  file:
    dir: .logs/.neat
    filename: neat.log
    size_mb: 1
    backups: 3
    age_days: 10
```

For each `git` repo where you want to use `neat` the following file should be present.

```yaml
git:
  branches:
    default: main
    prune: auto # none | select | auto
```

<!-- USAGE EXAMPLES -->
## CLI Usage

```shell
# displays help for neat
neat -h

# inits the global config
neat init

# inits a repo
neat init repo

# creates a new branch in the existing repo
neat new feature -n 'new-client-design'

# creates a new branch in a different repo
neat new feature -n 'another-feature' -p /tmp/go-neat-test

# refreshes the sub directories
neat refresh repos
```

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/delineateio/oss-template/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

If you would like to contribute to any delineate.io OSS projects please read:

* [Code of Conduct](https://github.com/delineateio/.github/blob/master/CODE_OF_CONDUCT.md)
* [Contributing Guidelines](https://github.com/delineateio/.github/blob/master/CONTRIBUTING.md)

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [Best README Template](https://github.com/othneildrew/Best-README-Template)
* [Markdown Badges](https://github.com/Ileriayo/markdown-badges)
* [DocToc](https://github.com/thlorenz/doctoc)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[pr-welcome-shield]: https://img.shields.io/badge/PRs-welcome-ff69b4.svg?style=for-the-badge&logo=github
[pr-welcome-url]: https://github.com/delineateio/oss-template/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue
[contributors-shield]: https://img.shields.io/github/contributors/delineateio/oss-template.svg?style=for-the-badge&logo=github
[contributors-url]: https://github.com/delineateio/oss-template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/delineateio/oss-template.svg?style=for-the-badge&logo=github
[forks-url]: https://github.com/delineateio/oss-template/network/members
[stars-shield]: https://img.shields.io/github/stars/delineateio/oss-template.svg?style=for-the-badge&logo=github
[stars-url]: https://github.com/delineateio/oss-template/stargazers
[issues-shield]: https://img.shields.io/github/issues/delineateio/oss-template.svg?style=for-the-badge&logo=github
[issues-url]: https://github.com/delineateio/oss-template/issues
[license-shield]: https://img.shields.io/github/license/delineateio/oss-template.svg?style=for-the-badge&logo=github
[license-url]: https://github.com/delineateio/oss-template/blob/master/LICENSE
