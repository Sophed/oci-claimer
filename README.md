# OCI Claimer

> [!WARNING]
> A \*nix dev environment such as Linux, MacOS or WSL for Windows is HIGHLY recommended to avoid issues

## Building from source
- Install Go from [their site](https://go.dev/dl/)
```
git clone https://github.com/Sophed/oci-claimer
cd oci-claimer
go build
```

## Don't want to build?
- You don't have to!
- Figure out if you're running ARM or x86 with `uname -m`
    - If you're running ARM, `wget https://github.com/Sophed/oci-claimer/releases/download/1.0/oci-claimer-ARM64.tar.gz`
    - If you're running x86, `wget https://github.com/Sophed/oci-claimer/releases/download/1.0/oci-claimer-x86_64.tar.gz`
- Extract the tar contents with `tar -xzf <tar-name>.tar.gz`

## Setup
- Log into your OCI panel and head to user settings

![image](/assets/user-settings.png)

- Head to the API keys section and add a new API key

![image](/assets/api-key.png)

- Download the key and make a note of the path

![image](/assets/download.png)

- Then add the contents of the config to `~/.oci/config`

![image](/assets/config.png)

- You should then edit the `key_file` value of that same config to the path of your previously downloaded API key

## Usage
- Start the program in the background, I personally use a `tmux` session
```bash
./ociclaimer
```