# usenet-drive (WIP)

## Description

This is a simple script that allows you to mount a usenet server as a webdav drive.

***This is not a tool to mount any nzb files, nzb files that are supported on the tool needs to be created by this tool.***

***Use at your own risk***

## Usage with rclone

### Install rclone

```bash
curl https://rclone.org/install.sh | sudo bash
```

### Configure rclone

```bash
rclone config
```

Add a new webdav remote with the following parameters:

- **Name**: `usenet`
- **URL**: `http://localhost:8080`

### Mount the remote

```bash
rclone mount --allow-other --async-read=true --dir-cache-time=1000h --buffer-size=32M --poll-interval=15s --rc --rc-no-auth --rc-addr=localhost:5572 --use-mmap --vfs-read-ahead=128M --vfs-read-chunk-size=32M --vfs-read-chunk-size-limit=2G --vfs-cache-max-age=504h --vfs-cache-mode=full --vfs-cache-poll-interval=30s --vfs-cache-max-size=50G --timeout=10m usenet: ${PATH_TO_MOUNT} --umask=002
```

## Features

- Allow mount nzb files as the original file
- Allow streaming of video files
- Allow upload new files full obfuscated to prevent DMCA takedowns
- Filesystem watch for new files and upload them to usenet automatically
- Multiple server for upload and download

## Todo

- [ ] Multiple server for download
- [ ] Open split files
- [ ] Open rar files
- [ ] Tests

## Docker usage

### Build

```bash
docker-compose up
```

## Config Struct

The `Config` struct defines the configuration for the Usenet Drive application. See an example at [config.yaml](config.example.toml).

### Fields

- `NzbPath` (string): The path to the NZBs file. This will be the path where all nzbs will be saved making the virtual file system.
- `ServerPort` (string): The port number for the server. Default value is `8080`.
- `Usenet` (Usenet): The Usenet configuration.
- `DBPath` (string): The path where the database will be saved. Default value is `/config/usenet-drive.db`.

## Usenet Struct

The `Usenet` struct defines the Usenet configuration.

### Fields

- `Download` (UsenetProvider): The Usenet provider for downloading.
- `Upload` (Upload): The Usenet options for uploading.

## Upload Struct

The `Upload` struct defines the Usenet provider for uploading.

### Fields

- `Provider` (UsenetProvider): The Usenet provider for uploading.
- `FileWhitelist` ([]string): The list of allowed file extensions. For example, `[".mkv", ".mp4"]`, in this case only files with the extensions `.mkv` and `.mp4` will be uploaded to usenet. Take care not upload files that change frequently, like subtitules or text files, since they will be uploaded every time they change. In usenet you can not edit files.
- `NyuuVersion` (string): The version of [Nyuu](https://github.com/animetosho/Nyuu). Default value is `0.4.1`. Used for upload files to usenet.
- `NyuuPath` (string): The path to [Nyuu](https://github.com/animetosho/Nyuu). Default value is `/config/nyuu`. If nyuu executable is not found, it will be auto downloaded for the given system arch.  Used for upload files to usenet.
- `MaxActiveUploads` (int): The maximum number of active uploads. Be aware that the number of active uploads are related with the number of maxConnections. For example, if your provider allows 20 connections, to avoid problems, if you want 2 active uploads you should put 10 max connections. Default value is `2`.
- `UploadIntervalInSeconds` (float64): The upload interval in seconds. After X seconds the system will check for pending uploads and perform them. Default value is `60`.

## UsenetProvider Struct

The `UsenetProvider` struct defines the Usenet provider configuration.

### Fields

- `Host` (string): The hostname of the Usenet provider. For example, `news.usenetserver.com`.
- `Port` (int): The port number of the Usenet provider. For example, `563`.
- `Username` (string): The username for the Usenet provider. For example, `user`.
- `Password` (string): The password for the Usenet provider. For example, `pass`.
- `Groups` ([]string): The list of Usenet groups. For example, `["alt.binaries.teevee", "alt.binaries.movies"]`.
- `SSL` (bool): Whether to use SSL for the Usenet provider. Default value is `true`.
- `MaxConnections` (int): The maximum number of connections to the Usenet provider.