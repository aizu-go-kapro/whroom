# whroom

`whroom` is a location sharing tool in our university. If you are in the building, `whroom` try to identify which room you are in and send it to server. Also, you can look up the other's last location log within the university by it.

## Usage

To get the location log of `s1230004`, type:

```
$ whroom get s1230004
```

## Setup

### Get binary

#### Homebrew

```
$ brew install aizu-go-kapro/whroom
```

#### go get

```
$ go get github.com/aizu-go-kapro/whroom
```

### Put config file

Some preferences are needed. ref. 

### Make it daemon

To logging your location periodicaly, an background job must be setup.

#### brew services

If you have installed it with Homebrew, it provides also the way to setup daemon.

```
$ brew services start whroom
```

#### Load plist file

For macOS user, the `.plist` file is provided.

```
$ cp ./whroom.plist /Library/LaunchDaemons/
$ launchd load /Library/LaunchDaemons/whroom.plist
```

#### Load Systemd Unit file

For Ubuntu or other Linux distributions using Systemd user, the `.service` file is provided.

```
$ sudo systemctl link ./whroom.service
$ sudo systemctl enable whroom.service
```

## Authors

- [dennougorilla](github.com/dennougorilla)
- [acomagu](github.com/acomagu)
- [Ryo3939](github.com/Ryo3939)
