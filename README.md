# hymn-sheet
Generates Mass hymn sheets with Liturgy of the Word overview (developed and tested on Linux. If you're on Windows or Mac, you're on your own!)

## Dependencies
This requires LaTeX to run (the programme generates a LaTeX file then shell exec's `lualatex`, providing it with the file in question). For Debian-based Linux distros you can install LaTeX with:
```bash
sudo apt install texlive-full
```
For RHEL/Fedora based it's
```bash
sudo dnf install texlive-scheme-full
```

For Windows or Mac, please consult the relevant documentation for installing LaTeX.

It also requires Arial font, which come as standard on Windows, but you will probably need to install them if you are using Linux. For Ubuntu it's
```bash
sudo apt install ttf-mscorefonts-installer
```

For compilation, I've included a makefile. Make can be installed with
```bash
sudo apt install make
```

## Installation
Requires Go to build from source. Install go using your relevant package manager, e.g.:
```bash
sudo apt install go
```
or see [installation instructions](https://go.dev/doc/install#install)

The programme can be built by running:
```bash
make
```

## Running
There are some example `json` files which are uses as input in the `examples` directory. They are passed to the programme like so:
```bash
./hymn-book-generator -file examples/2025-09-07.json
```

The `date` field in the `json` file is used as the lookup when the programme reads `resource/ordo.json` to work out the proper for that date (i.e. 22nd Sunday in Ordinary Time Year C). It then uses that information to look up the Liturgy of the Word Scripture references and antiphons from `resources/calendar.json`. Following that, it looks up the hymns from the `hymns` array in the config file, and fetches the lyrics (if they exist) in `resources/LAU or resources/LHOAN`.  Finally it generates a `.tex` field, filling it with info from the config file and the things it's look up (metioned above). After generating the `.tex` file, it shell exec's `lualatex` on the `.tex` file to generate a PDF hymn sheet that includes LOTW Scripture references, etc.

This programme doesn't guarantee to generate a perfectly formatted sheet, but it aims to take away a lot of the manual effort. The generated `.tex` file allows for manual tweaks and for the user to rerun `lualatex` on it to fine-tune the formatting.

Ideally, once there is a mature repository of hymn lyrics and a complete `calendar.json`, the only file that needs to be maintained is the `ordo.json`, as well as the user creating a new config json file for the given Sunday.