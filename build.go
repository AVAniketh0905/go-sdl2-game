package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

const (
	SrcPath   = "C:/mingw64/bin/" // Path where your SDL2 DLLs are located
	BuildsDir = "builds/"         // Directory to store builds
	ExeName   = "possessed"       // Name of your executable
)

func Build() {
	buildNumber := getNextBuildNumber() // Get the next build number
	buildDir := fmt.Sprintf("%sbuild_v%d/", BuildsDir, buildNumber)
	os.MkdirAll(buildDir, os.ModePerm) // Create new build directory

	// Build the project
	fmt.Println("Building the project...")
	exePath := fmt.Sprintf("%s%s.exe", buildDir, ExeName)
	if err := buildExecutable(exePath); err != nil {
		log.Fatalf("Error building executable: %v", err)
	}

	// Copy SDL2 DLLs
	fmt.Println("Copying SDL2 DLLs...")
	if err := copyFile("SDL2.dll", buildDir); err != nil {
		log.Fatalf("Error copying SDL2.dll: %v", err)
	}
	if err := copyFile("SDL2_Image.dll", buildDir); err != nil {
		log.Fatalf("Error copying SDL2_Image.dll: %v", err)
	}
	if err := copyFile("SDL2_Mixer.dll", buildDir); err != nil {
		log.Fatalf("Error copying SDL2_Image.dll: %v", err)
	}

	// Create ZIP archive
	zipPath := fmt.Sprintf("%sbuild_v%d.zip", BuildsDir, buildNumber)
	fmt.Println("Creating ZIP archive...")
	if err := createZIPArchive(buildDir, zipPath); err != nil {
		log.Fatalf("Error creating ZIP archive: %v", err)
	}

	fmt.Println("Build completed successfully!")
}

// buildExecutable compiles the Go code into an executable
func buildExecutable(outputPath string) error {
	cmd := exec.Command("go", "build", "-ldflags", "-H=windowsgui", "-o", outputPath, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// copyFile copies a file from sourcePath to destinationDir
func copyFile(filename, destinationDir string) error {
	sourcePath := filepath.Join(SrcPath, filename)
	destinationPath := filepath.Join(destinationDir, filename)

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

// createZIPArchive creates a ZIP archive of the specified directory and saves it to outputPath
func createZIPArchive(sourceDir, outputPath string) error {
	zipFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, err := filepath.Rel(sourceDir, path)
			if err != nil {
				return err
			}

			zipEntry, err := zipWriter.Create(relPath)
			if err != nil {
				return err
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(zipEntry, file)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

// getNextBuildNumber gets the next available build number
func getNextBuildNumber() int {
	latestBuild := 0
	filepath.Walk(BuildsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			dirName := filepath.Base(path)
			if dirName != "builds" && len(dirName) > 7 && dirName[:7] == "build_v" {
				buildNumber, err := strconv.Atoi(dirName[7:])
				if err == nil && buildNumber > latestBuild {
					latestBuild = buildNumber
				}
			}
		}
		return nil
	})
	return latestBuild + 1
}
