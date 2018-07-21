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

Some preferences are needed.

```
firebase_url = "https://tmp.firebaseio.com"
student_id = "s1230004"
wifi_interface = "wlp2s0"
duration = "1s"
```

The search paths for it:
- ~/.whroom.toml
- ~/.config/whroom/config.toml

The `firebase_url` for UoA students: ref. [http://web-int.u-aizu.ac.jp/~s1230004/whroom.toml](http://web-int.u-aizu.ac.jp/~s1230004/whroom.toml)

The `--config` option can be used to specify additional path.

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
$ launchctl load /Library/LaunchDaemons/whroom.plist
```

#### Load Systemd Unit file

For Ubuntu or other Linux distributions using Systemd user, the `.service` file is provided.

```
$ vi whroom.service  # Edit it to fit your environment.
$ sudo systemctl link ./whroom.service
$ sudo systemctl enable whroom
$ sudo systemctl start whroom
```

## Authors

- [denougorilla](https://github.com/dennougorilla)
- [acomagu](https://github.com/acomagu)
- [Ryo3939](https://github.com/Ryo3939)
