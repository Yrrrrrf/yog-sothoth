package health

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

type CheckResult struct {
	Name    string
	Status  string // "pass", "warn", "fail"
	Message string
}

func RunAllChecks() []CheckResult {
	results := []CheckResult{}

	results = append(results, checkTools()...)
	results = append(results, checkEnv())
	results = append(results, checkGit())
	results = append(results, checkFile("README.md", "README"))
	results = append(results, checkFile("LICENSE", "License"))

	return results
}

func checkTools() []CheckResult {
	results := []CheckResult{}
	
	// Deno
	if _, err := exec.LookPath("deno"); err != nil {
		results = append(results, CheckResult{"Deno Runtime", "warn", "Deno not found in PATH"})
	} else {
		out, _ := exec.Command("deno", "--version").Output()
		version := strings.Split(string(out), "\n")[0]
		results = append(results, CheckResult{"Deno Runtime", "pass", version})
	}

	// Bun
	if _, err := exec.LookPath("bun"); err != nil {
		results = append(results, CheckResult{"Bun Runtime", "warn", "Bun not found in PATH"})
	} else {
		out, _ := exec.Command("bun", "--version").Output()
		version := "bun " + strings.TrimSpace(string(out))
		results = append(results, CheckResult{"Bun Runtime", "pass", version})
	}
	
	return results
}

func checkEnv() CheckResult {
	envVars, err := godotenv.Read(".env")
	if err != nil {
		if os.IsNotExist(err) {
			return CheckResult{".env Completeness", "warn", ".env file is missing"}
		}
		return CheckResult{".env Completeness", "fail", "Failed to read .env file"}
	}

	exampleVars, err := godotenv.Read(".env.example")
	if err != nil {
		if os.IsNotExist(err) {
			return CheckResult{".env Completeness", "pass", ".env found (no .env.example to check against)"}
		}
		return CheckResult{".env Completeness", "warn", "Failed to read .env.example"}
	}

	missing := []string{}
	for key := range exampleVars {
		if _, ok := envVars[key]; !ok {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		return CheckResult{".env Completeness", "fail", fmt.Sprintf("Missing %d variables defined in .env.example", len(missing))}
	}

	return CheckResult{".env Completeness", "pass", "All variables present"}
}

func checkGit() CheckResult {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		return CheckResult{"Git Repository", "warn", "Not initialized"}
	}
	return CheckResult{"Git Repository", "pass", "Initialized"}
}

func checkFile(filename, contextName string) CheckResult {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return CheckResult{contextName, "warn", "Missing " + filename}
	}
	return CheckResult{contextName, "pass", "Found " + filename}
}
