## Steam verify age bypass
Bypasses the annoying age prompt on the Steam Store page.

Due to legal requirements, steam is forced to always request your age when browsing games with a PEGI rating of 16 or higher. This disables that nonsense once and for all.

> *I hate repeating the same utterly useless action over and over again, such as entering my birthdate a godzillion times in the Steam store...*


## Donate
Psst... if u liked this, please send me some BTC so I can buy kebab: 
- `bc1qzaen9cr8q754zah86fkzrk7m7ura8t9r89rsgm`

## Usage


### Tip
you may use the `--file` parameter to point to a custom location of where your steam Cookie file is. Example:
```bash
./bypass.elf --file  /home/urusername/.steam/steam/config/htmlcache/Cookies
```

### Expected output
The expected output will look like this depending on os, and install location

```
implementing bypass...
Found Cookies file:
> /home/urusername/.steam/steam/config/htmlcache/Cookies
Bypass completed successfully!
You can now open Steam and access the store without age verification.
You may close the program.
```


### Linux

You have 3 options:

- Build the source code from [`source (RTFM)`](https://go.dev/doc/tutorial/compile-install)
- Use a one-liner command to bypass automatically (most easy and preferred way):
  
  ```bash
  /bin/bash -c "$(curl -sSL 'https://raw.githubusercontent.com/aamaanaa/steam-verify-age-bypass/refs/heads/main/bypass.sh')"
  ```
  
- Or, use the precompiled (*- Linux btw*) binary from the [`releases page->`](https://github.com/aamaanaa/steam-verify-age-bypass/releases/download/v1.1.0/bypass.elf)
  
  ```bash
  # Make executable
  chmod +x ./bypass.elf
  # Run it
  ./bypass.elf
  ```

### Windows

You have 2 options:

- Build the source code from [`source (RTFM)`](https://go.dev/doc/tutorial/compile-install)
- Or, use the precompiled binary from the [`releases page->`](https://github.com/aamaanaa/steam-verify-age-bypass/releases/download/v1.0.0/bypass.exe)
  
  ```bash
  # Run it
  ./bypass.exe
  ```
 - Or, simply double click and run it
