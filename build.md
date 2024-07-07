# How to Build

## Instructions

### Easier Build

Just run `build.go` file!

### Manual Build

1 Build in `go`

```bash
go build -ldflags "-H=windowsgui" -o builds/gome.exe .
```

2 Copy Dependencies

- SDL.dll
- SDL_image.dll

3 Run the executable

```bash
./builds/gome.exe
```

4 Zip the build

```bash
zip -r builds/gome.zip builds
```
