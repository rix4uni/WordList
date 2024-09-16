package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const maxTokenSize = 10 * 1024 * 1024 // 10 MB, adjust as necessary

func main() {
	// Root directory for nuclei templates
	rootDir := os.ExpandEnv("$HOME/cent-nuclei-templates")

	// Walk through the directory recursively and process .yaml files
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only .yaml files
		if filepath.Ext(path) == ".yaml" {
			processYAMLFile(path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking through directory: %v", err)
	}

	// Execute the command to concatenate and process the output files
	fmt.Println("Running unew to create *-all.txt file for every directory...")

	// List all directories
	findCmd := exec.Command("find", ".", "-type", "d")
	output, err := findCmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute find command: %v", err)
	}

	// Split the output into lines (directories)
	directories := strings.Split(string(output), "\n")

	for _, directory := range directories {
		// Skip empty lines and the current directory (.)
		if directory == "" || directory == "." {
			continue
		}

		// Construct the command to process files in the directory
		cmdStr := fmt.Sprintf("cat %s/%s-*.txt | unew -t -el -q %s/%s-all.txt", directory, directory, directory, directory)

		// Execute the command
		cmd := exec.Command("sh", "-c", cmdStr)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to execute command for directory %s: %v", directory, err)
		}
	}

	fmt.Println("Finished...")
}

// processYAMLFile reads and processes each YAML file
func processYAMLFile(templateFile string) {
	// Open the file for reading
	file, err := os.Open(templateFile)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Variables to store extracted values
	var tags, severity string
	var baseURLs []string

	// Regular expressions to match fields in the YAML file
	tagsRegex := regexp.MustCompile(`^ *tags: *(.+)$`)
	severityRegex := regexp.MustCompile(`^ *severity: *(.+)$`)
	baseURLRegex := regexp.MustCompile(`- ['"]{{BaseURL}}([^'"]*)['"]`)

	// Create a scanner with a large buffer
	scanner := bufio.NewScanner(file)
	buf := make([]byte, maxTokenSize)
	scanner.Buffer(buf, maxTokenSize)

	for scanner.Scan() {
		line := scanner.Text()

		// Match tags
		if matches := tagsRegex.FindStringSubmatch(line); len(matches) > 0 {
			tags = matches[1]
		}

		// Match severity
		if matches := severityRegex.FindStringSubmatch(line); len(matches) > 0 {
			severity = matches[1]
		}

		// Match all BaseURL paths
		if matches := baseURLRegex.FindStringSubmatch(line); len(matches) > 0 {
			baseURLs = append(baseURLs, matches[1])
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Split tags into individual values
	tagList := strings.Split(tags, ",")

	// Regular expression to validate tag names (alphanumeric, dashes, underscores only)
	validTagRegex := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)

	// Define a map to correct severity values
    severityMapping := map[string]string{
        "none":       "unknown",
        "informative": "info",
        "meduim": 	"medium",
        "hight":      "high",
        "highx":      "high",
        "criticall":  "critical",
        "ciritical":  "critical",
        "cretical":   "critical",
    }

    // List of valid severities
    validSeverities := map[string]bool{
        "unknown":  true,
        "info":     true,
        "low":      true,
        "medium":   true,
        "high":     true,
        "critical": true,
    }

    // Get the correct severity value if there's a match in the mapping
    correctedSeverity, exists := severityMapping[strings.ToLower(severity)]
    if exists {
        severity = correctedSeverity
    }

    // Skip if severity is not valid
    if !validSeverities[strings.ToLower(severity)] {
        fmt.Printf("Skipping invalid severity: '%s' in file: %s\n", severity, templateFile)
        return
    }

    // Check if no BaseURLs were extracted, skip file creation if so
    if len(baseURLs) == 0 {
        fmt.Printf("Skipping file creation for: %s because no BaseURLs were found.\n", templateFile)
        return
    }

	// Extract the base filename from the template file path
    baseFilename := filepath.Base(templateFile)

	for _, tag := range tagList {
		tag = strings.TrimSpace(tag)

		// Skip empty or invalid tags
		if tag == "" || !validTagRegex.MatchString(tag) {
			fmt.Printf("Skipping invalid tag: '%s'\n", tag)
			continue
		}

		// Create directory for the tag
		err := os.MkdirAll(tag, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory '%s': %v", tag, err)
		}

		// Convert severity to lowercase for the filename
        lowercaseSeverity := strings.ToLower(severity)

		// Define the output file name
		outputFile := filepath.Join(tag, fmt.Sprintf("%s-%s.txt", tag, lowercaseSeverity))

		// Prepare a command to run `unew` with the BaseURLs through stdin
		cmd := exec.Command("unew", "-t", "-el", "-q", outputFile)
		cmd.Stdin = strings.NewReader(strings.Join(baseURLs, "\n")) // Provide baseURLs as stdin

		// Run the command
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to execute unew command: %v", err)
		}

		// Display the message in the format: SOURCE: <YAML File> => DESTINATION: <Output File>
        fmt.Printf("SOURCE:%s => DESTINATION:%s\n", baseFilename, outputFile)
	}
}
