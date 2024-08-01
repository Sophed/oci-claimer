# OCI Claimer
This script is made with the intention of helping new developers get access to useful resources for their projects, that being the free tier of Oracle Cloud Infrastructure, which offers a cloud VPS with 24GB of memory and 4 OCPUs.

If this script helps you at all or you would like to support my work, consider donating to my ko-fi 🩷

[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/sophed)

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
- Download the latest release for your system with `wget <tar-name>.tar.gz`
- Extract the tar contents with `tar -xzf <tar-name>.tar.gz`

## Configuration
- `discord_id` is the user ID of your Discord account, used to notify you
- `webhook_url` is the URL for the Discord webhook you want to use
- `ssh_public_Key` is the path to the SSH public key you will use to connect to the instance
- `notify_out_of_capacity` determines whether or not to send webhook messages for every "out of capacity" error, will cause spam
- `retry_delay` is how long to wait before trying all domains again (recommended 120)
- `availability_domain_switch_delay` is how long to wait between each domain attempt (recommended 60)
- `availability_domains` is a list of all the availability domains to try, [find them here](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm#ad-names)
- `display_name` is the name of the instance to create
- `shape` determines the shape of the instance, can be found on OCI panel
- `ocpus` is the number of OCPUs to allocate (max 4 for free tier)
- `memory_gbs` is the amount of memory to allocate in GBs (max 24 for free tier)
- `boot_volume_gbs` is the size of the boot volume (storage) in GBs, free tier has a global limit of 200, I recommend 100 to save space for other instances
- `compartment_id` is the ID of your OCI compartment, [find it here](https://docs.oracle.com/en-us/iaas/Content/GSG/Tasks/contactingsupport_topic-Locating_Oracle_Cloud_Infrastructure_IDs.htm#Finding_the_OCID_of_a_Compartment)
- `image_id` is the ID of the OS image to use for your instace, I actually don't know how you're supposed to find these, I just used dev tools in my browser to grab it from the HTTP request
- `subnet_id` is the ID of the subnet for your instance, can be found in the subnet details

## Pre-launch setup
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