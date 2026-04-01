#!/usr/bin/env bash

set -xe

rm -rf go.mod go.sum
go mod init github.com/upbound/provider-opentofu
go get -tool golang.org/x/tools/cmd/goimports@latest
go get \
	github.com/MakeNowJust/heredoc@latest \
	github.com/alecthomas/kingpin/v2@latest \
	github.com/angelokurtis/go-otel/starter@latest \
	github.com/crossplane/crossplane-runtime/v2@latest \
	github.com/crossplane/crossplane-tools@latest \
	github.com/google/go-cmp@latest \
	github.com/google/uuid@latest \
	github.com/hashicorp/go-getter@latest \
	github.com/pkg/errors@latest \
	github.com/spf13/afero@latest \
	go.uber.org/zap@latest \
	k8s.io/api@latest \
	k8s.io/apiextensions-apiserver@latest \
	k8s.io/apimachinery@latest \
	k8s.io/client-go@latest \
	sigs.k8s.io/controller-runtime@latest \
	sigs.k8s.io/controller-tools@latest \
  && \
  go mod tidy
