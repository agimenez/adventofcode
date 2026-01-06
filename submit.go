package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get positional arguments
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: go run submit.go <level> <answer>\n")
		os.Exit(1)
	}

	level := os.Args[1]
	answer := os.Args[2]

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	// Extract year and day from path
	dir := filepath.Base(cwd)
	parentDir := filepath.Base(filepath.Dir(cwd))

	year, err := strconv.Atoi(parentDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not parse year from directory path\n")
		os.Exit(1)
	}

	day, err := strconv.Atoi(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not parse day from directory path\n")
		os.Exit(1)
	}

	// Get session cookie from environment
	sessionCookie := os.Getenv("AOC_SESSION")
	if sessionCookie == "" {
		fmt.Fprintf(os.Stderr, "Error: AOC_SESSION environment variable not set\n")
		os.Exit(1)
	}

	// Submit the answer
	submitAnswer(year, day, level, answer, sessionCookie)
}

func submitAnswer(year, day int, level, answer, sessionCookie string) {
	// Construct the URL
	submitURL := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)

	// Prepare form data
	formData := url.Values{}
	formData.Set("level", level)
	formData.Set("answer", answer)

	// Create request
	req, err := http.NewRequest("POST", submitURL, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating request: %v\n", err)
		os.Exit(1)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionCookie))
	req.Header.Set("User-Agent", "aoc-submit-script")

	// Send request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		os.Exit(1)
	}

	// Parse response
	responseStr := string(body)

	// Check for common responses
	switch {
	case strings.Contains(responseStr, "That's the right answer"):
		fmt.Println("✅ Correct! That's the right answer!")
		if level == "1" {
			fmt.Printf("Part 2 is now unlocked for day %d.\n", day)
		} else {
			fmt.Printf("You've completed day %d!\n", day)
		}
	case strings.Contains(responseStr, "That's not the right answer"):
		fmt.Println("❌ That's not the right answer.")
		if strings.Contains(responseStr, "too high") {
			fmt.Println("Your answer is too high.")
		} else if strings.Contains(responseStr, "too low") {
			fmt.Println("Your answer is too low.")
		}
		if strings.Contains(responseStr, "Please wait") {
			fmt.Println("You have a rate limit. Please wait before trying again.")
		}
	case strings.Contains(responseStr, "You gave an answer too recently"):
		fmt.Println("⏱️  You gave an answer too recently. Please wait before submitting again.")
	case strings.Contains(responseStr, "You don't seem to be solving the right level"):
		fmt.Println("🔒 You don't seem to be solving the right level. Did you already complete it?")
	case strings.Contains(responseStr, "Puzzle inputs differ by user"):
		fmt.Println("🔑 Authentication failed. Please check your session cookie.")
	default:
		fmt.Println("⚠️  Unexpected response. Check if you're logged in properly.")
	}
}
