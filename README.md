# TFLint Ruleset CIS AWS Foundations Benchmark

![CIS AWS Foundations Benchmark Version](https://img.shields.io/badge/CIS%20benchmark%20version-1.5.0-green)
![https://gruntwork.io/?ref=repo_cis_compliance_aws"](https://img.shields.io/badge/maintained%20by-gruntwork.io-%235849a6.svg)

Tflint rules for CIS AWS Foundations Benchmark compliance checks. These rules work in addition to the recommendations from [Gruntwork's CIS Service Catalog](https://github.com/gruntwork-io/terraform-aws-cis-service-catalog).

> :warning: **This repository is a WIP. It only contains one single rule so far, to validate Security Groups, that is hard to enforce in any other way ([see Rules section](#rules)). In the future, we may add other CIS AWS Foundations Benchmark rules.**


## Requirements

- TFLint v0.40+
- Go v1.19

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "aws-cis" {
  enabled = true

  version = "<VERSION>"
  source  = "github.com/gruntwork-io/tflint-ruleset-aws-cis"
}
```

## Rules

| Name                                       | Description                                                                        |Severity|Enabled| CIS Recommendation |
|--------------------------------------------|------------------------------------------------------------------------------------| --- | --- |--------------------|
| aws_security_group_rule_invalid_cidr_block | Ensure that SG rules do not allow public access to remote administration ports     |ERROR|✔| 5.2 and 5.3        |

## Terragrunt

An effective way to enforce these rules is to add them to your Terragrunt configuration using [Before Hooks](https://terragrunt.gruntwork.io/docs/features/hooks/#tflint-hook).

```hcl
terraform {
  before_hook "before_hook" {
    commands     = ["apply", "plan"]
    execute      = ["tflint"]
  }
}
```

In the root of the Terragrunt project, add a `.tflint.hcl` file, replacing `<VERSION>` below with the latest version from the [releases page](https://github.com/gruntwork-io/tflint-ruleset-aws-cis/releases):

```hcl
plugin "aws" {
    enabled = true
    version = "<VERSION>"
    source  = "github.com/gruntwork-io/tflint-ruleset-aws-cis"
}
```


## Running locally

### Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```bash
$ cat << EOS > .tflint.hcl
plugin "aws-cis" {
  enabled = true
}
EOS
$ tflint
```

### Manual release

In order to release the binaries, this project uses [goreleaser](https://goreleaser.com/) ([install instructions](https://goreleaser.com/install/)).

Export the variable `GPG_FINGERPRINT` in order to sign the release, and `GITHUB_TOKEN` so the binaries can be uploaded to GitHub. The release should run locally from the tag that will have the release.

```
git checkout <TAG FOR THE RELEASE, e.g. v0.40.0>

export GPG_FINGERPRINT=<FINGERPRINT_ID>
export GITHUB_TOKEN=<TOKEN>

goreleaser release
```
