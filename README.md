[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/ebrahimahmadi/ar-cli">
    <img src="https://www.arvancloud.com/images/other/arvan-api-docs-logo.svg" alt="Logo" width="180" height="180">
  </a>

  <h3 align="center">Arvan Cloud CLI</h3>

  <p align="center">
    An awesome CLI to intract with dns-sec, domains, cloud-sec and the CDN of Arvan cloud
    <br />
    <a href="https://www.arvancloud.com/docs/api/cdn/4.0"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template">View Demo</a>
    ·
    <a href="https://github.com/ebrahimahmadi/ar-cli/issues/new">Report Bug</a>
    ·
    <a href="https://github.com/ebrahimahmadi/ar-cli/issues/new">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
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
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#cmd">Commands</a>
      <ul>
        <li><a href="#DNS">DNS</a></li>  
        <li><a href="#CDN">CDN</a></li>  
        <li><a href="#domain">Domain</a></li>  
        <li><a href="#auth">Auth</a></li>  
        <li><a href="#CS">Cloud Security</a></li>
      </ul>
    </li>        
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://drive.google.com/file/d/1uS6r6udhcvMOLe_28Gsk1RHRP9slOq3J/view?usp=sharing)

This package provides a unified command line interface to Arvan CDN Services.


### Built With

This section should list any major frameworks that you built your project using. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.
* [Go](https://golang.org/)
* [Cobra](https://github.com/spf13/cobra)

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Installation

1. Get a free API Key at [Arvan Cloud](https://npanel.arvancloud.com/profile/api-keys)
2. Clone the repo
   ```sh
   git clone https://github.com/ebrahimahmadi/ar-cli.git
   ```
3. Install Dependencies
   ```sh
   go install
   ```
4. Set Your API key using the CLI
   ```sh
   go run main.go auth --key 'Apikey xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx';
   ```
   If you are succssufully authenticated, you will see `Configuration saved successfully` message.


<!-- Domain -->
## Playing with the Domain
The domain command will enable you to have interaction with domains tied with your account.

To start interacting with domains run 
   ```sh
    ar-cli domain [command-to-execute] [relavent-flags]
   ```

| Command  | Desc | Example 
|---|---|---
|  search | serach for a domain by given key-word. if the key-word is not passed it will list all domains  |   ```ar-cli domain search --key-word arvan```
|  info | will retreive data for an specific record and parse the data in a table  |   ```ar-cli domain info --name arvancloud.com```
| remove  | will remove given domain | ```ar-cli domain remove --name arvancloud.com --id 3541b0ce-e8a6-42f0-b65a-f03a7c387486```
|ns-records|fetch and show ns-keys and ns-domains of the given domain| ```ar-cli domain ns-records --name arvancloud.com```
|create| create given domain name| ```ar-cli domain create --name hello.com```
|check| send a request to Arvan to recheck your NS. Also will report the status of ns|```ar-cli domain check --name arvancloud.com```


<!-- USAGE EXAMPLES -->
## Usage Examples

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/ebrahimahmadi/ar-cli/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


<!-- CONTACT -->
## Contact

Ebi - eahmadi641@gmail.com

Project Link: [Arvan Cli](https://github.com/ebrahimahmadi/ar-cli)

[forks-shield]: https://img.shields.io/github/forks/ebrahimahmadi/ar-cli
[forks-url]: https://github.com/ebrahimahmadi/ar-cli/network/members
[stars-shield]: https://img.shields.io/github/stars/ebrahimahmadi/ar-cli
[stars-url]: https://github.com/ebrahimahmadi/ar-cli/stargazers
[issues-shield]: https://img.shields.io/bitbucket/issues-raw/ebrahimahmadi/ar-cli
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[product-screenshot]: images/screenshot.png
