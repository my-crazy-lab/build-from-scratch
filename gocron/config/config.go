package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

type Env struct {
	Key   string `validate:"required,uppercase" yaml:"key"`
	Value string `validate:"required" yaml:"value"`
}

type Command struct {
	Command    string `validate:"required" yaml:"command"`
	FileOutput string `validate:"omitempty" yaml:"file_output"`
}

type Job struct {
	Name     string    `validate:"required" yaml:"name"`
	Cron     string    `validate:"required,cron" yaml:"cron"`
	Envs     []Env     `validate:"omitempty,dive,required" yaml:"envs"`
	Commands []Command `validate:"required,dive,required" yaml:"commands"`
}

type Config struct {
	Defaults struct {
		Cron string `yaml:"cron"`
		Envs []Env  `yaml:"envs"`
	} `yaml:"defaults"`
	Jobs []Job `validate:"required,dive,required" yaml:"jobs"`
}

func (c *Config) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}

func mergeEnvs(defaultEnvs, jobEnvs []Env) []Env {
	// Use defaultEnvs as the base
	envMap := make(map[string]int)
	mergedEnvs := make([]Env, len(defaultEnvs))
	copy(mergedEnvs, defaultEnvs)

	// Map keys in defaultEnvs to their positions
	for i, env := range defaultEnvs {
		envMap[env.Key] = i
	}

	// Process jobEnvs
	for _, env := range jobEnvs {
		if idx, exists := envMap[env.Key]; exists {
			// Override value if key already exists
			mergedEnvs[idx].Value = env.Value
		} else {
			// Append new key-value pair
			mergedEnvs = append(mergedEnvs, env)
			envMap[env.Key] = len(mergedEnvs) - 1
		}
	}

	return mergedEnvs
}

func (c *Config) processConfig() {
	for i, job := range c.Jobs {
		if job.Cron == "" {
			c.Jobs[i].Cron = c.Defaults.Cron
		}
		c.Jobs[i].Envs = mergeEnvs(c.Defaults.Envs, job.Envs)
	}
}

func New(filePath string) (*Config, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %s, error: %v", filePath, err)
	}

	var config Config
	if err := yaml.Unmarshal([]byte(fileData), &config); err != nil {
		return nil, fmt.Errorf("failed to parse file: %s, error: %v", filePath, err)
	}

	config.processConfig()
	if err := config.Validate(); err != nil {
		return nil, err
	}
	return &config, nil
}
