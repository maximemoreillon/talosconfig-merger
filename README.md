# Talosconfig-merger

Merges discrete talosconfig files into `~/.talos/config`.
By default, this application will look for talosconfig files in the `~/.talos/config.d` directory.

## Usage

```bash
./talosconfig-merger
```

Specifying a source directory other than `~/.talos/config.d`

```bash
./talosconfig-merger -source /path/to/config/dir
```
