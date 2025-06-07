# v0.9.0 (pre-release)
## Added:
* **Syntax highlighting**
* **Detecting programming/markup language for correct syntax highlighting**
## Modified:
* **The code has been rewritten in go** (the original version from [sharkdp](https://github.com/sharkdp) was written in the Rust programming language)
# v0.9.1 (pre-release)
## Added:
* **Support for various flags:** \
    `-n`: - Line numbering in output. In this case, syntax highlighting will be disabled. \
    `--git`: Shows code changes compared to the repository. Highlights changed lines. In this case, syntax highlighting is disabled.\
    `--version`: Displays the version.
## Modified:
* The project structure has been changed | main.go does not contain any functions, they call from the highlight package.