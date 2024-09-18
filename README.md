# player-protocol

**Open VLC and Pot Player directly  from browser.Support ``file://`` ``https://`` ``http://`` protocols.You can't open local file directly from browser in any of these players so player-protocol enables this through file:// protocol.**

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
vlc://file:///path/to/local/file #linux

vlc://file://X:/path/to/local/file # windows

vlc://https://ia804605.us.archive.org/22/items/big-buck-bunny-4k/Big_Buck_Bunny_4K.mp4

potplayer://https://www.youtube.com/watch?v=dQw4w9WgXcQ

```
**For PotPlayer use `potplayer://` instead of `vlc://`**
