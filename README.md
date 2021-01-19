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
        <li><a href="#auth">Auth</a></li>  
        <li><a href="#domain">Domain</a></li>  
        <li><a href="#cloud-security">Cloud Security</a></li>
        <li><a href="#dns">DNS</a></li>  
      </ul>
    </li>        
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

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
## Domain
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


<!-- Cloud Security -->
## Cloud security
The domain command will enable you to get an overview of your cloud-security app and update the subscription plan.

To start interacting with domains run 
   ```sh
    ar-cli cloud-security [command-to-execute] [relavent-flags]
   ```

| Command  | Desc | Example 
|---|---|---
|  info | Shows the current status of cloud-security service  |   ```ar-cli cloud-security info --name arvancloud.com```
|  update | Updates your cloud-security service plan. Available subscriptions plans are  bronze, silver, gold, platinum  |   ```ar-cli cloud-security update --name arvancloud.com --plan gold```

## DNS

By using dns command, you can have complete management over DNS services.

To start interacting with domains run 
   ```sh
    ar-cli dns [command-to-execute] [relavent-flags]
   ```
| Command  | Desc | Example 
|---|---|---
|  list | Lists all records that are tied with the given domain with their details  |   ```ar-cli dns list --name arvancloud.com```
|  info | Show single record details. details consists of: id,type,name,host,cloud status,ttl, upstream http,protection status,IP filter count, geo filter, ip filter order |   ```ar-cli dns info --name arvancloud.com --record-id 65b03b20-3598-4a1b-a467-c6e3a4ec652a```
|  remove | remove a record tied to the domain  |   ```ar-cli dns remove --name arvancloud.com --record-id 65b03b20-3598-4a1b-a467-c6e3a4ec652a```
|  toggle | toggles cloud service. if `--cloud` flag is passed to the command the cloud service will be enlabled for the record. in case you want to disable cloud service remove `--cloud` at the end of the command  |   ```ar-cli dns toggle --name arvancloud.com --record-id 65b03b20-3598-4a1b-a467-c6e3a4ec652a --cloud```

### dns create record command
In order to create a record the syntax will be a little bit different.

To start interacting with domains run 
   ```sh
    ar-cli dns [record-type] [command-to-execute] [relavent-flags]
   ```

  available record types:
   - a-record
   - aaaa-record
   - mx-record
   - ns-record
   - srv-record
   - txt-record
   - spf-record
   - dkim-record
   - aname-record
   - cname-record
   - ptr-record

| Command  | Desc | Example 
|---|---|---
| craete | created new record with  | ```ar-cli dns create a-record --name arcancloud.com --ttl 120 --ip 192.168.0.1 --ip-filter-order rr --ip-geo-filter country --upstream-https https --ip-filter-count single```

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/ebrahimahmadi/ar-cli/issues) for a list of proposed features (and known issues).

But general feature under develops are :
CDN-accelration
CDN-firewall
CDN-cache
CDN-https
CDN-apps

Also in case you have an idea please <a href="https://github.com/ebrahimahmadi/ar-cli/issues/new">Request Feature</a>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Commit conventions
We follow [conventional-commit](https://www.conventionalcommits.org/) to make our track easier and 
better to review.

acceptable scopes for commits are:

1. core
2. validator
3. api
4. cdn-segment
5. dns-segment
6. cs-segment
7. domain-segment
8. auth-segment
9. configuaration

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
[product-screenshot]: demo.gif
