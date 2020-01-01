# Cyaml

Split and combine multiple yaml files into one.

## Motivation

Think of a helm chart combined into one yaml file.

## Usage

Combine:

```sh
./cyaml combine example/nfs -s example/nfs/
```

Split:

```sh
./cyaml split example/nfs.yaml nfs
```