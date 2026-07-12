# langcheck

[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/suzuki-shunsuke/langcheck) | [INSTALL](INSTALL.md) | [USAGE](USAGE.md)

`langcheck` is a CLI for checking if disallowed characters are included in files and texts.
This is useful to prevent coding agents from using disallowed characters in code, commit messages, issues, pull requests, and so on.
In the context of OSS, we usually use English, but sometimes coding agents use non-English languages because many non-English speakers communities with coding agents in their native language.
Even if suggesting using English in documents like `AGENTS.md`, coding agents sometimes ignore the instruction.
By running `langcheck` in Git hooks, you can prevent disallowed characters from being committed.

## Usage

```sh
langcheck check [--text "<text>"]... [<file>...]
```

## Examples

### hk

https://hk.jdx.dev/

hk.pkl

```pkl
amends "package://github.com/jdx/hk/releases/download/v1.50.0/hk@1.50.0#/Config.pkl"
import "package://github.com/jdx/hk/releases/download/v1.50.0/hk@1.50.0#/Builtins.pkl"

local linters = new Mapping<String, Step> {
  ["langcheck"] {
    // https://github.com/suzuki-shunsuke/langcheck
    glob = List("**/*")
    check = "langcheck check {{files}}"
    fix = "langcheck check {{files}}"
    batch = true
  }
}

hooks {
  ["pre-commit"] {
    fix = true
    stash = "git"
    steps = linters
  }
  ["commit-msg"] {
    steps {
      ["langcheck"] {
        // https://github.com/suzuki-shunsuke/langcheck
        check = "langcheck check {{commit_msg_file}}"
      }
    }
  }
  ["fix"] {
    fix = true
    steps = linters
  }
  ["check"] {
    steps = linters
  }
}
```
