package test

import (
    "testing"

    "github.com/gruntwork-io/terratest/modules/docker"
    "github.com/gruntwork-io/terratest/modules/packer"
    "github.com/stretchr/testify/assert"
)

func TestPackerHelloWorldExample(t *testing.T) {
    packerOptions := &packer.Options{
        Template: "./build.json.pkr.hcl",
    }

    packer.BuildArtifact(t, packerOptions)

    opts := &docker.RunOptions{Command: []string{"cat", "/test.txt"}}
    output := docker.Run(t, "gruntwork/packer-hello-world-example", opts)
    assert.Equal(t, "Hello, World!", output)
}

