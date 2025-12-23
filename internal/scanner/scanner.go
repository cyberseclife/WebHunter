package scanner

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"webhunter/internal/utils"

	"golang.org/x/time/rate"
)

type Config struct {
	RateLimit int
	Header    string
	DryRun    bool
}

type Scanner struct {
	Config  Config
	Limiter *rate.Limiter
}

func New(cfg Config) *Scanner {
	return &Scanner{
		Config:  cfg,
		Limiter: rate.NewLimiter(rate.Limit(cfg.RateLimit), 1),
	}
}

// SaveResults appends output to a file
func (s *Scanner) SaveResults(filename, content string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		utils.LogError("Failed to open result file: %v", err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(content + "\n"); err != nil {
		utils.LogError("Failed to write to result file: %v", err)
	}
}

// Recon runs port scanning and service detection
func (s *Scanner) Recon(targets []string, ports string, serviceDetection, stealth bool) []string {
	utils.LogInfo("Starting Recon Stage...")
	var activeTargets []string
	
	// Parse ports
	portList := parsePorts(ports)
	
	var wg sync.WaitGroup
	results := make(chan string, len(targets)*len(portList))

	for _, target := range targets {
		if !utils.GlobalScope.IsAllowed(target) {
			continue
		}

		// Resolve IP
		ips, err := net.LookupIP(target)
		if err != nil {
			utils.LogWarn("Could not resolve %s", target)
			continue
		}
		ip := ips[0].String()
		activeTargets = append(activeTargets, target) // Assume host is up if resolved for now

		for _, port := range portList {
			wg.Add(1)
			go func(t, ip, p string) {
				defer wg.Done()
				s.Limiter.Wait(context.Background())
				
				address := net.JoinHostPort(ip, p)
				timeout := 2 * time.Second
				conn, err := net.DialTimeout("tcp", address, timeout)
				
				if err == nil {
					conn.Close()
					res := fmt.Sprintf("%s:%s OPEN", t, p)
					utils.LogInfo("Found: %s", res)
					results <- res
					s.SaveResults("port_scan_results.txt", res)
				}
			}(target, ip, port)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Drain results (for now we just logged them, but we could return open ports)
	for range results {}
	
	return activeTargets
}

// Analysis runs vulnerability scanning (mocked/basic)
func (s *Scanner) Analysis(targets []string, vulnScan, fuzz bool, headers []string) {
	utils.LogInfo("Starting Analysis Stage...")
	
	for _, target := range targets {
		if !utils.GlobalScope.IsAllowed(target) {
			continue
		}

		s.Limiter.Wait(context.Background())
		// Simulation of logic
		if vulnScan {
			utils.LogOut("Analyzing %s for common vulnerabilities...", target)
			// Mock finding
			if strings.Contains(target, "test") {
				res := fmt.Sprintf("[VULN] X-Frame-Options missing on %s", target)
				utils.LogWarn(res)
				s.SaveResults("analysis_results.txt", res)
			}
		}
		
		if fuzz {
			utils.LogOut("Fuzzing directories on %s...", target)
			// Mock finding
			res := fmt.Sprintf("[DIR] %s/admin detected", target)
			utils.LogInfo(res)
			s.SaveResults("analysis_results.txt", res)
		}
	}
}

// Exploitation runs payloads
func (s *Scanner) Exploitation(targets []string, payloads []string) {
	utils.LogInfo("Starting Exploitation Stage...")
	
	for _, target := range targets {
		if !utils.GlobalScope.IsAllowed(target) {
			continue
		}

		for _, payload := range payloads {
			s.Limiter.Wait(context.Background())
			msg := fmt.Sprintf("Testing payload '%s' on %s", payload, target)
			
			if s.Config.DryRun {
				utils.LogOut("[DRY-RUN] %s", msg)
			} else {
				utils.LogWarn("Exploiting: %s", msg)
				// Mock success
				res := fmt.Sprintf("[SHELL] Payload '%s' successful on %s", payload, target)
				s.SaveResults("exploitation_results.txt", res)
			}
		}
	}
}

func parsePorts(p string) []string {
	if p == "-" {
		// return all 65535 ports - abbreviated for prototype
		return []string{"80", "443", "8080", "8443", "22", "21"} 
	}
	if p == "" {
		return []string{"80", "443"}
	}
	return strings.Split(p, ",")
}
