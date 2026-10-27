# Install

langcheck is written in Go. So you only have to install a binary in your `PATH`.

There are some ways to install langcheck.

1. [Homebrew](#homebrew)
1. [Scoop](#scoop)
1. [aqua](#aqua)
1. [mise-en-place](#mise-en-place)
1. [GitHub Releases](#github-releases)
1. [Build an executable binary from source code yourself using Go](#build-an-executable-binary-from-source-code-yourself-using-go)

## Homebrew

You can install langcheck using [Homebrew](https://brew.sh/).

```sh
brew install suzuki-shunsuke/langcheck/langcheck
```

## Scoop

You can install langcheck using [Scoop](https://scoop.sh/).

```sh
scoop bucket add suzuki-shunsuke https://github.com/suzuki-shunsuke/scoop-bucket
scoop install langcheck
```

## aqua

You can install langcheck using [aqua](https://aquaproj.github.io/).

```sh
aqua g -i suzuki-shunsuke/langcheck
```

## Build an executable binary from source code yourself using Go

```sh
go install github.com/suzuki-shunsuke/langcheck/cmd/langcheck@latest
```

## GitHub Releases

You can download an asset from [GitHub Releases](https://github.com/suzuki-shunsuke/langcheck/releases).
Please unarchive it and install a pre built binary into `$PATH`. 

### Verify downloaded assets from GitHub Releases

You can verify downloaded assets using some tools.

1. [GitHub CLI](https://cli.github.com/)
1. [Cosign](https://github.com/sigstore/cosign)
1. [Packslip](https://packslip.dev/) `langcheck >= v0.0.2`
1. [slsa-verifier](https://github.com/slsa-framework/slsa-verifier) `langcheck <= v0.0.1`

### 1. GitHub CLI

You can install GitHub CLI by aqua.

```sh
aqua g -i cli/cli
```

```sh
version=v0.0.2
asset=langcheck_darwin_arm64.tar.gz
gh release download -R suzuki-shunsuke/langcheck "$version" -p "$asset"
gh attestation verify "$asset" \
  -R suzuki-shunsuke/langcheck \
  --signer-workflow suzuki-shunsuke/go-release-workflow/.github/workflows/release.yaml
```

### 2. Cosign

You can install Cosign by aqua.

```sh
aqua g -i sigstore/cosign
```

Cosign v2.4.2 or later is required.

```sh
version=v0.0.2
checksum_file="langcheck_checksums.txt"
asset=langcheck_darwin_arm64.tar.gz
gh release download "$version" \
  -R suzuki-shunsuke/langcheck \
  -p "$asset" \
  -p "$checksum_file" \
  -p "${checksum_file}.sigstore.json"
cosign verify-blob \
  "$checksum_file" \
  --bundle "${checksum_file}.sigstore.json" \
  --certificate-identity-regexp 'https://github\.com/suzuki-shunsuke/go-release-workflow/\.github/workflows/release\.yaml@.*' \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com"
cat "$checksum_file" | sha256sum -c --ignore-missing -
```

### 3. Packslip

`langcheck >= v0.0.2`

You can install packslip by aqua.

```sh
aqua g -i jdx/packslip
```

```sh
version=v0.0.2
asset=langcheck_darwin_arm64.tar.gz
gh release download -R suzuki-shunsuke/langcheck "$version" -p "$asset" -p packslip.sigstore.json
packslip verify packslip.sigstore.json \
    --pin ps1_fboao3n7berji7j5v2e3fllywq \
    --identity-prefix https://github.com/suzuki-shunsuke/go-release-workflow/.github/workflows/release.yaml@ \
    --artifact "$asset"
```

- `--pin` requires the release to be signed from the repository `suzuki-shunsuke/langcheck`, identified by its repository ID rather than its name
- `--identity-prefix` requires the release to be signed by the release workflow of [suzuki-shunsuke/go-release-workflow](https://github.com/suzuki-shunsuke/go-release-workflow)

On success, packslip prints a line starting with `ok:`.
Please make sure it shows the project and version you meant, e.g. `github.com/suzuki-shunsuke/langcheck 0.0.2`.

### 4. slsa-verifier

> [!WARNING]
> As of langcheck v0.0.2, `multiple.intoto.jsonl` hasn't been published anymore.

You can install slsa-verifier by aqua.

```sh
aqua g -i slsa-framework/slsa-verifier
```

```sh
version=v0.0.1
asset=langcheck_darwin_arm64.tar.gz
gh release download -R suzuki-shunsuke/langcheck "$version" -p "$asset" -p multiple.intoto.jsonl
slsa-verifier verify-artifact "$asset" \
  --provenance-path multiple.intoto.jsonl \
  --source-uri github.com/suzuki-shunsuke/langcheck \
  --source-tag "$version"
```
