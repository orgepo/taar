Hello, thank if you are read this before send PRs (:

Before send a PR, make sure to follow this rules:

# PR title
Your PR title should follow [`conventional commit`](https://www.conventionalcommits.org/en/v1.0.0/) rule.

# Linting
Make sure you used `make check` command before sending PR.


# Some good make command

you can use this commands for ease of develope:

### make devtools
install all tools you need to develope taar.

### make fmt
formats your code.

### make check
run the `golangci-linter` on your code.

### make build
build the code.

### make cp
this command needs **SUDO** permission, then it's build the code and move it to /usr/bin.
