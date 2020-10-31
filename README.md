# httcli
Simple and fast CLI to make HTTP request in a convenient and straightforward way

- [Description](#doing-requests-like-never-before)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [Support](#support)
- [License](#license)

## Doing requests like never before
Tired of curl? Postman is too weighty? Or maybe you just want an easy tool to test your APIs in a fast way. I think that you will find in httcli a response to all of your requests!

## Features
- Very fast and pretty light (damn Go, you create some big binaries at times)
- Easy to use
- Totally written in Go

## Installation
**LINUX**

Download the zip in the releases ([this](https://github.com/GianlucaTarantino/httcli/releases/download/v1.0.0/installer.zip), but always check if there is a newer zip).
Extract it, and in the folder where you extracted the zip, run this command
```bash 
sudo sh ./install.sh
```
*If this doesn't work, follow the instruction in info.txt, or just contact me*

**WINDOWS**

On Windows you need to compile the project, so you will need the Go compiler (or wait until the developer releases an exe for this. Will happen pretty soon, don't worry). Download the repository, go in the repository directory and, with the terminal, give this command
```batch
go build -o ./httcli.exe ./src/main.go
```
Then just move the just created exe to a folder included in the PATH environmental variable.

## Usage
Simply write `httcli` in terminal to start the CLI.
It is pretty straightforward. Select the type of request, write the host of the request, a path for the file that will be used for the body (I used this option because I thought that letting the user use the editor he wants would be better, than to force him using a specific editor), add some header if you need and click on the "Send Request" button. A request will be made and you will see all the informations of the response in the center section of the screen. There is also an history of the requests, so you can make a request you did earlier just by clicking on the request you want from the history!

[![asciicast](https://asciinema.org/a/RPOETCynv8aemv1vAlHBIzpzf.svg)](https://asciinema.org/a/RPOETCynv8aemv1vAlHBIzpzf)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. There is already a template for pull requests and issues.

## Support
For any problem regarding httcli, you can always open an issue! If you want to contact me, feel free to write me at gianlutara@gmail.com

# License

[MIT](https://github.com/GianlucaTarantino/httcli/blob/main/LICENSE)
