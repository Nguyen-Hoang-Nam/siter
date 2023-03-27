# Siter - Simple Terminal

Base on [here](https://ishuah.com/2021/03/10/build-a-terminal-emulator-in-100-lines-of-go/)

![Ubuntu](https://raw.githubusercontent.com/namngh/assets/main/siter/linux-ls.jpg)
*Ubuntu*

![Windows 10](https://raw.githubusercontent.com/namngh/assets/main/siter/windows-echo.jpg)
*Windows 10*

## Configuration

Not defined yet.

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

- [ ] Linux - Work in progress
- [ ] Windows CMD - Work in progress
- [ ] Windows Powershell - Not supported yet
- [ ] Windows WSL - Not supported yet
- [ ] MacOS - Not supported yet

## TODO

- [ ] Render color correctly
- [ ] Add graphic protocol
- [ ] Add tab manager

## Contributing

Pull requests are welcome. For major changes,
please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit)