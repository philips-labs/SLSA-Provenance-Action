package options

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/philips-labs/slsa-provenance-action/lib/github"
	"github.com/philips-labs/slsa-provenance-action/lib/intoto"
)

const defaultGenerateOutputPath = "provenance.json"

// GenerateOptions Commandline flags used for the generate command.
type GenerateOptions struct {
	GitHubContext  string
	RunnerContext  string
	OutputPath     string
	ExtraMaterials []string
}

// GetGitHubContext The '${github}' context value, retrieved in a GitHub workflow.
func (o *GenerateOptions) GetGitHubContext() (*github.Context, error) {
	if o.GitHubContext == "" {
		return nil, RequiredFlagError("github-context")
	}
	var gh github.Context
	if err := json.Unmarshal([]byte(o.GitHubContext), &gh); err != nil {
		return nil, fmt.Errorf("failed to unmarshal github context json: %w", err)
	}
	return &gh, nil
}

// GetRunnerContext The '${runner}' context value, retrieved in a GitHub workflow.
func (o *GenerateOptions) GetRunnerContext() (*github.RunnerContext, error) {
	if o.RunnerContext == "" {
		return nil, RequiredFlagError("runner-context")
	}
	var runner github.RunnerContext
	if err := json.Unmarshal([]byte(o.RunnerContext), &runner); err != nil {
		return nil, fmt.Errorf("failed to unmarshal runner context json: %w", err)
	}
	return &runner, nil
}

// GetOutputPath The location to write the provenance file.
func (o *GenerateOptions) GetOutputPath() (string, error) {
	if o.OutputPath == "" {
		return "", RequiredFlagError("output-path")
	}
	return o.OutputPath, nil
}

// GetExtraMaterials Additional material files to be used when generating provenance.
func (o *GenerateOptions) GetExtraMaterials() ([]intoto.Item, error) {
	var materials []intoto.Item

	for _, extra := range o.ExtraMaterials {
		file, err := os.Open(extra)
		if err != nil {
			return nil, fmt.Errorf("failed retrieving extra materials: %w", err)
		}
		defer file.Close()

		m, err := intoto.ReadMaterials(file)
		if err != nil {
			return nil, fmt.Errorf("failed retrieving extra materials for %s: %w", extra, err)
		}
		materials = append(materials, m...)
	}

	return materials, nil
}

// AddFlags Registers the flags with the cobra.Command.
func (o *GenerateOptions) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&o.GitHubContext, "github-context", "", "The '${github}' context value.")
	cmd.PersistentFlags().StringVar(&o.RunnerContext, "runner-context", "", "The '${runner}' context value.")
	cmd.PersistentFlags().StringVar(&o.OutputPath, "output-path", defaultGenerateOutputPath, "The path to which the generated provenance should be written.")
	cmd.PersistentFlags().StringSliceVarP(&o.ExtraMaterials, "extra-materials", "m", nil, "The '${runner}' context value.")
}
