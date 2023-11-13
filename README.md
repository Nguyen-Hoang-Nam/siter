# Siter - Simple Terminal Emulator

Base on [here](https://ishuah.com/2021/03/10/build-a-terminal-emulator-in-100-lines-of-go/)

![Ubuntu](https://raw.githubusercontent.com/namngh/assets/main/siter/linux-ls.jpg)
*Ubuntu*

![Windows 10](https://raw.githubusercontent.com/namngh/assets/main/siter/windows-echo.jpg)
*Windows 10*

## Configuration

If the `SITER_CONFIG_DIRECTORY` environment variable is not defined, the configuration directory will 
be set to `path.Join(os.UserConfigDir, "siter")`. Siter reads configuration from the `config.toml` file.

```toml
open_url_with = "brave"
foreground_color = "#FFFFFF"
background_color = "#000000"
```

## Run

```
go run main.go
```

## Build

Requirements:

- fyne
- fyne-cross

Check how to build and install cross-compile [here](https://developer.fyne.io/started/cross-compiling)

```
fyne-cross linux -arch=* -icon "icon/icon-32x32.png"
fyne-cross windows -arch=amd64,386 -icon "icon/icon-32x32.png"
```

## Platform

- [ ] Shell - WIP
- [ ] Bash - WIP
- [ ] ZSH - WIP
- [ ] CMD - WIP
- [ ] Powershell - Not supported yet
- [ ] WSL - Not supported yet

## Contributing

Pull requests are welcome. For major changes,
please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit)