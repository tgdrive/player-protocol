# player-protocol

**Open VLC and Pot Player from browser.Support ``file://`` ``https://`` ``http://`` file protocols.**

## Setup

### Install on Windows

- Download the latest release from [here](https://github.com/tgdrive/player-protocol/releases/latest).
- Extract the zip file.
- Run ``protocol-register.bat`` as administrator.

**Uninstall:**
- Run ``protocol-deregister.bat`` as administrator.

### Install on Linux

```bash
curl -s https://sh-install.vercel.app/tgdrive/player-protocol?move=1 | bash
curl -LO https://raw.githubusercontent.com/tgdrive/player-protocol/main/vlc-protocol.desktop
xdg-desktop-menu install vlc-protocol.desktop
rm vlc-protocol.desktop
```
**Uninstall:**
```bash
xdg-desktop-menu uninstall vlc-protocol.desktop
sudo rm /usr/local/bin/player-protocol
```
## Usage

```
vlc://file:///path/to/file

vlc://https://example.com/path/to/file

vlc://http://example.com/path/to/file

```
**For PotPlayer use `potplayer://` instead of `vlc://`**
