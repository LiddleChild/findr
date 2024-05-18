# findr

A cli application for searching file name / content with best-effort speed purely written in Go with the minimal (of one) dependency. The search algorithm simply uses KMP pattern searching algorithm.

## Basic command

```bash
findr <query> [options]
```

> **Note:** `finder --help` to see full descriptions

| Option               | Description                                                              |
| -------------------- | ------------------------------------------------------------------------ |
| ` -c \| --content`   | Search with file content mode<br>**\*Defaults to filename mode\***       |
| `-i \| --ignore`     | Ignore paths / directories to be walked                                  |
| `-mx \| --max-depth` | Max directory depth to be walked<br>**\*Level 0 is working directory\*** |
| `-d \| --dir`        | Set working directory                                                    |
| `-C \| --case`       | Set case sensitive search                                                |

## Journey

The main focus is pruely on content mode only since it is the slowest part of searching.

> **Note:** Every performance numbers seen afterwards were run on Windows 10 with AMD Ryzen 5 3600 and 16 GB of RAM and were averaged by 5 consecutive runs after a fresh `go build`.

### **v1**

In the first version, searching is done by walking down file tree, grabbing file content by `os.ReadFile()`, and plugging that directly into KMP algorithm which is copied and pasted from the Internet. A SvelteKit project is used as a benchmark by searching throughout the project (including `node_modules`). This version runs in approximately 33 seconds.

### **v2**

The performance was nearly doubled in this version as the search algorithm is rewritten by doing file content reading and pattern searching simultaneously. This reduces overhead of moving loooooooong string around the memory. Anyways, runtime reduces to about 17 seconds.

### **v3**

This latest version is, of course, the fastest with a whopping 6 seconds runtime. Inevitably, the speed up was achieved using the power of parallelization. The core algorithm is splitted into two phases, file walking and content searching which have their own goroutines. As the walker routine walks through the file tree, paths are sent through a channel and fed into a group of workers which runs the same algorithm as before, v2 algorithm. The number of workers are set to be the same with the number of logical cores.

## Discussion

Realization: Go has built-in file tree walker lol :clown_emoji:.
Moreover, cli parsing library would does a better job than mine.
Both would have further reduced run time.
<br><br>
Anyways, there was soooooooo much fun doing the project, writing cli parser, writing tailor-made pattern matching algorithm, and of course go concurrency. Seeing runtime went from half a minute to few seconds was pretty satisfying. Also, seeing every cores spiked up simultaneouly was a pure joy.

> **Ps.** Concurrency could not happen without this great article [https://go.dev/blog/pipelines](https://go.dev/blog/pipelines)
