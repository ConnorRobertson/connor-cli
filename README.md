<a id="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/ConnorRobertson/connor-cli">
    <img src="images/connor_cli_logo.png" width="300" height="300">
  </a>

<h3 align="center">Connor CLI</h3>

  <p align="center">
    A CLI project written in Golang for learning and experimentation purposes!
    <br>
    <a href="https://github.com/ConnorRobertson/connor-cli/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    &middot;
    <a href="https://github.com/ConnorRobertson/connor-cli/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

A custom CLI for Connor by Connor. Supports list (ls), lsall (recursive list basically), greet (says hello to an input argument), cre (touch), del (remove), rename (mv filename filename), help, and running without a command.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

As this is a project to learn Golang this is going to be almost entirely written in Go!

- [![Go][Go.dev]][go-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

This is just a template I will add more details to this section when this is in a more complete state.

### Prerequisites

Install Go

- [Here is a Great Guide](https://go.dev/doc/install)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/ConnorRobertson/connor-cli.git
   ```
2. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin ConnorRobertson/connor-cli
   git remote -v # confirm the changes
   ```
3. Compile code into a binary (Go should automatically create a binary based on your OS installation)
   ```sh
   go build -o connorcli cmd/main.go
   ```
4. You can now run `./connorcli`
   ```sh
   ./connorcli -h
   ```
5. When you are done running Connor CLI you can remove the binaries with `go clean`

   Alternatively, if you are on Ubuntu (haven't tested any other distros) you can run the shell script it will need sudo privileges.
   It adds the output binary to the /usr/local/bin so you can just run `connorcli`

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the unlicense. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

**Connor Robertson** - robertsonc4@gmail.com

**Project Link:** [https://github.com/ConnorRobertson/connor-cli](https://github.com/ConnorRobertson/connor-cli)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

- [README Template](https://github.com/othneildrew/Best-README-Template/tree/main)
- [Cobra CLI Package](https://github.com/spf13/cobra-cli/tree/main)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/ConnorRobertson/connor-cli.svg?style=for-the-badge
[contributors-url]: https://github.com/ConnorRobertson/connor-cli/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/ConnorRobertson/connor-cli.svg?style=for-the-badge
[forks-url]: https://github.com/ConnorRobertson/connor-cli/network/members
[stars-shield]: https://img.shields.io/github/stars/ConnorRobertson/connor-cli.svg?style=for-the-badge
[stars-url]: https://github.com/ConnorRobertson/connor-cli/stargazers
[issues-shield]: https://img.shields.io/github/issues/ConnorRobertson/connor-cli.svg?style=for-the-badge
[issues-url]: https://github.com/ConnorRobertson/connor-cli/issues
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/connor-robertson-software-engineer
[go.dev]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[go-url]: https://go.dev/
