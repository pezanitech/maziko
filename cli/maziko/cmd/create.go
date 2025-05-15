package cmd

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var templateRepo = "https://github.com/pezanitech/maziko"

var availableTemplates = map[string]string{
	"react": "templates/react",
}

// Represents the create command
var createCmd = &cobra.Command{
	Use:   "create [template]",
	Short: "Create a new Maziko project",
	Long: fmt.Sprintf(
		"%s\n\n%s\n%s\n",
		"Create a new Maziko project based on a template.",
		"Currently supported templates:",
		"  - react: A React-based frontend with Inertia.js",
	),
	Args: cobra.ExactArgs(1),
	Run:  runCreate,
}

// Initialize the create command
func init() {
	RootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP(
		"name", "n", "",
		"Name of the project (defaults to the template name)",
	)

	createCmd.Flags().StringP(
		"output", "o", "",
		"Directory to create the project in (defaults to current directory)",
	)
}

// Handler that outputs logs without timestamps and keys
type CustomHandler struct {
	w io.Writer
}

// Implements slog.Handler
func (h *CustomHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

// Implements slog.Handler
func (h *CustomHandler) Handle(_ context.Context, r slog.Record) error {
	level := strings.ToUpper(r.Level.String())
	msg := r.Message

	output := fmt.Sprintf("%s %s", level, msg)

	// Process attributes if any
	r.Attrs(func(a slog.Attr) bool {
		if a.Key != "time" && a.Key != "level" && a.Key != "msg" {
			output += fmt.Sprintf(" %s=%v", a.Key, a.Value.Any())
		}
		return true
	})

	_, err := fmt.Fprintln(h.w, output)
	return err
}

// Implements slog.Handler.
func (h *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

// Implements slog.Handler.
func (h *CustomHandler) WithGroup(name string) slog.Handler {
	return h
}

// Runs the create command
func runCreate(cmd *cobra.Command, args []string) {
	// initialize custom logger without timestamps and extra keys
	logger := slog.New(&CustomHandler{w: os.Stderr})

	templateName := args[0]

	// validate template
	templatePath, ok := availableTemplates[templateName]
	if !ok {
		logger.Error(
			"Unknown template",
			"template", templateName,
			"available", getAvailableTemplates(),
		)
		os.Exit(1)
	}

	// get flags
	projectName, _ := cmd.Flags().GetString("name")
	outputDir, _ := cmd.Flags().GetString("output")

	// set defaults if not provided
	if projectName == "" {
		projectName = templateName
	}

	if outputDir == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			logger.Error(
				"Failed to get current directory",
				"error", err,
			)
			os.Exit(1)
		}

		outputDir = filepath.Join(
			currentDir,
			projectName,
		)
	} else {
		outputDir = filepath.Join(
			outputDir,
			projectName,
		)
	}

	// create output directory if it doesn't exist
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		logger.Info(
			"Creating project directory",
			"path", outputDir,
		)

		if err := os.MkdirAll(outputDir, 0755); err != nil {
			logger.Error(
				"Failed to create project directory",
				"error", err,
			)
			os.Exit(1)
		}
	} else {
		// check if directory is empty
		empty, err := isDirEmpty(outputDir)
		if err != nil {
			logger.Error(
				"Failed to check if directory is empty",
				"error", err,
			)
			os.Exit(1)
		}
		if !empty {
			logger.Error(
				"Directory is not empty",
				"path", outputDir,
			)
			os.Exit(1)
		}
	}

	// clone template
	logger.Info(
		"Creating new Maziko project",
		"template", templateName,
		"path", outputDir,
	)

	// clone a specific template and not the whole repo
	if err := cloneTemplate(outputDir, templatePath); err != nil {
		logger.Error(
			"Failed to clone template",
			"error", err,
		)
		os.Exit(1)
	}

	// Customize the template
	if err := customizeTemplate(outputDir, projectName); err != nil {
		logger.Error("Failed to customize template", "error", err)
		os.Exit(1)
	}

	// Success message
	logger.Info("✅ Project created successfully!", "path", outputDir)
	logger.Info("Next steps:")
	logger.Info("Make sure you have go & pnpm installed")
	logger.Info("cd " + projectName)
	logger.Info("maziko install")
	logger.Info("maziko dev")
}

// Helper function to check if directory is empty
func isDirEmpty(dir string) (bool, error) {
	f, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

// Helper function to clone a template
func cloneTemplate(outputDir string, templatePath string) error {
	// create a temp directory for the full repo
	tmpDir, err := os.MkdirTemp("", "maziko-clone-")
	if err != nil {
		return fmt.Errorf(
			"failed to create temporary directory: %w", err,
		)
	}
	defer os.RemoveAll(tmpDir)

	// clone the repository
	cmd := exec.Command(
		"git", "clone", "--depth", "1", templateRepo, tmpDir,
	)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"failed to clone repository: %w", err,
		)
	}

	// copy the temp directory to the output directory
	templateSourcePath := filepath.Join(tmpDir, templatePath)
	return copyDir(templateSourcePath, outputDir)
}

// Helper function to recursively copy a directory
func copyDir(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// skip .git directory
			if entry.Name() == ".git" {
				continue
			}

			if err := os.MkdirAll(dstPath, 0755); err != nil {
				return err
			}
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {

			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// Copies a file
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Sync()
}

// Customizes the template with project-specific values
func customizeTemplate(projectDir string, projectName string) error {
	// update package.json with project name
	packageJSONPath := filepath.Join(projectDir, "package.json")

	if _, err := os.Stat(packageJSONPath); err == nil {
		content, err := os.ReadFile(packageJSONPath)
		if err != nil {
			return err
		}

		// TODO: use proper JSON parsing
		updated := strings.Replace(
			string(content),
			`"name": "maziko-react"`,
			fmt.Sprintf(`"name": "%s"`, projectName), 1,
		)

		if err := os.WriteFile(packageJSONPath, []byte(updated), 0644); err != nil {
			return err
		}
	}

	mazikoJSONPath := filepath.Join(projectDir, "maziko.json")

	// update maziko.json with project name
	if _, err := os.Stat(mazikoJSONPath); err == nil {
		content, err := os.ReadFile(mazikoJSONPath)
		if err != nil {
			return err
		}

		// TODO: use proper JSON parsing
		updated := strings.Replace(
			string(content),
			`"name": "Maziko App"`,
			fmt.Sprintf(`"name": "%s"`, projectName), 1,
		)

		if err := os.WriteFile(mazikoJSONPath, []byte(updated), 0644); err != nil {
			return err
		}
	}

	envExamplePath := filepath.Join(projectDir, ".env.example")
	envPath := filepath.Join(projectDir, ".env")

	// create a .env file from .env.example
	if _, err := os.Stat(envExamplePath); err == nil {
		if err := copyFile(envExamplePath, envPath); err != nil {
			return err
		}
	}

	return nil
}

// Get a list of available templates as a string
func getAvailableTemplates() string {
	templates := make([]string, 0, len(availableTemplates))

	for name := range availableTemplates {
		templates = append(templates, name)
	}

	return strings.Join(templates, ", ")
}
