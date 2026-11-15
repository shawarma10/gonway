# gonway
Conway's Game of Life written in Go using [Ebitengine](https://ebitengine.org).

The board is a 160×120 grid shown in a 640×480 window. It starts empty, so draw some cells and watch them evolve.

## Cloning & Building
Requires Go 1.27 or newer. On Linux, Ebitengine also needs a C compiler and a few X11/OpenGL development packages; see [Ebitengine's install guide](https://ebitengine.org/en/documents/install.html) for the list for your distro.

```bash
git clone https://github.com/shawarma10/gonway
cd gonway
go build -ldflags="-s -w" -o gonway cmd/gonway/main.go
./gonway
```

## Usage
| Input | Action |
|---|---|
| Left mouse (hold and drag) | Bring cells to life |
| Right mouse (hold and drag) | Kill cells |
| `Space` | Pause / resume (the window title shows `[Paused]`) |
| `R` | Clear the board |
| `1`–`9`, `0` | Set the speed: `1` is fastest (60 generations/s), `0` is slowest (6 generations/s). Default is `6` (10 generations/s). |

You can draw while paused or while the simulation is running. Pausing first makes it easier to set up a pattern.

## Rules
Each cell has 8 neighbors, including diagonals. On every generation:

- A live cell with 2 or 3 live neighbors survives; otherwise it dies.
- A dead cell with exactly 3 live neighbors comes to life.

Cells beyond the edge of the board count as dead, so the board doesn't wrap around.

To try it out, pause, draw a glider, and resume:

```
.#.
..#
###
```

## License
MIT. See [LICENSE](LICENSE).
