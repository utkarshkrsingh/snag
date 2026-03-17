package config

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func (c *Config) Validate() error {
	if len(c.Watch) == 0 {
		return errors.New("no watch rules defined: at least one rule is required in the 'watch' section")
	}

	if c.Global.Debounce != "" {
		if _, err := time.ParseDuration(c.Global.Debounce); err != nil {
			return fmt.Errorf("invalid global debounce duration: %s : %w", c.Global.Debounce, err)
		}
	}

	for i, rule := range c.Watch {
		if strings.TrimSpace(rule.Name) == "" {
			return fmt.Errorf("rule at index %d is missing a 'name'", i)
		}

		if len(rule.Patterns) == 0 {
			return fmt.Errorf("rule '%s' must have at least one pattern (e.g., '**/*.go')", rule.Name)
		}

		if len(rule.Run) == 0 {
			return fmt.Errorf("rule '%s' is missing a 'run' command", rule.Name)
		}
	}

	return nil
}
