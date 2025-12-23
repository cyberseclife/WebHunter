package utils

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"sync"
)

type ScopeManager struct {
	IncludeRules []string
	ExcludeRules []*regexp.Regexp
	mu           sync.RWMutex
}

var GlobalScope *ScopeManager

func InitScope(includeFile, excludeFile string) error {
	GlobalScope = &ScopeManager{
		IncludeRules: []string{},
		ExcludeRules: []*regexp.Regexp{},
	}
	return GlobalScope.Load(includeFile, excludeFile)
}

func (s *ScopeManager) Load(includePath, excludePath string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Load Includes (optional, if empty usually implies everything provided in CLI is allowed)
	if includePath != "" {
		inc, err := readLines(includePath)
		if err == nil {
			s.IncludeRules = inc
		}
	}

	// Load Excludes
	if excludePath != "" {
		exc, err := readLines(excludePath)
		if err == nil {
			for _, pattern := range exc {
				// We treat lines as regex
				re, err := regexp.Compile(pattern)
				if err != nil {
					// fallback to literal string match logic if regex fails? 
					// For now, let's assume valid regex or escape it.
					// Actually, simpler is to treat as substring if regex fails,
					// but standard is usually regex or glob. Let's use Regex.
					LogWarn("Invalid regex in exclude file: %s", pattern)
					continue
				}
				s.ExcludeRules = append(s.ExcludeRules, re)
			}
		}
	}
	return nil
}

func (s *ScopeManager) IsAllowed(target string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Check Exclusions first
	for _, rule := range s.ExcludeRules {
		if rule.MatchString(target) {
			LogError("Skipping: %s - Out of Scope", target)
			return false
		}
	}

	// If IncludeRules are defined, target must match at least one.
	// If empty, we assume the user explicitly provided the target via CLI so it's trusted (unless excluded).
	if len(s.IncludeRules) > 0 {
		matched := false
		for _, rule := range s.IncludeRules {
			// Simple contains or regex? Let's go with substring for includes or Regex?
			// Let's assume Includes are exact domains or regex.
			// For simplicity: If include rule is part of the target.
			if strings.Contains(target, rule) {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}

	return true
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}
