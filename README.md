# usenet-drive

## Description

This is a simple script that allows you to mount a usenet server as a drive on your computer.

## Usage with rclone

```bash
rclone mount --async-read=true --dir-cache-time=1000h --buffer-size=32M --poll-interval=15s --rc --rc-no-auth --rc-addr=localhost:5572 --use-mmap --vfs-read-ahead=128M --vfs-read-chunk-size=32M --vfs-read-chunk-size-limit=2G --vfs-cache-max-age=60m --vfs-cache-mode=full --vfs-cache-poll-interval=30s --vfs-cache-max-size=1G --timeout=10m ${YOUR_REMOTE_NAME}: ${PATH_TO_MOUNT} -vv
```

## Features

- Allow mount nzb files as the original file
- Allow streaming of video files

## Todo

- [ ] Upload new files
- [ ] Open split files
- [ ] Open rar files
