#### release command (local system)

```bash
GH_TOKEN="$(gh auth token)" eg compute local --hotswap -e GH_TOKEN
```

#### install flatpak daemon

```bash
mkdir derp; cd derp
flatpak remote-add --if-not-exists --user flathub https://dl.flathub.org/repo/flathub.flatpakrepo
flatpak-builder --user --install-deps-from=flathub --install --ccache --force-clean derp .eg.cache/flatpak.daemon.yml
flatpak run --user space.retrovibe.Daemon
```

#### install flatpak gui

```bash
mkdir derp; cd derp
flatpak remote-add --if-not-exists --user flathub https://dl.flathub.org/repo/flathub.flatpakrepo
flatpak-builder --user --install-deps-from=flathub --install --ccache --force-clean derp .eg.cache/flatpak.client.yml
flatpak run --user space.retrovibe.Client
```
