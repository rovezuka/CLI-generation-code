package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "gencli.exe"}

	var createFileCmd = &cobra.Command{
		Use:   "createfile [file-name]",
		Short: "Create a new file in the current directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]

			// Получаем путь к директории, где находится исполняемый файл
			exePath, err := os.Executable()
			if err != nil {
				fmt.Printf("Failed to get executable path: %s\n", err)
				return
			}

			// Извлекаем директорию из пути к исполняемому файлу
			projectDir := filepath.Dir(exePath)

			// Создаем новый файл в директории проекта
			filePath := filepath.Join(projectDir, fileName)
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Printf("Failed to create file: %s\n", err)
				return
			}
			defer file.Close()

			fmt.Printf("Created file: %s\n", filePath)
		},
	}

	rootCmd.AddCommand(createFileCmd)

	var createDirCmd = &cobra.Command{
		Use:   "createdir [dir-name]",
		Short: "Create a new directory in the current directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dirName := args[0]

			exePath, err := os.Executable()
			if err != nil {
				fmt.Printf("Failed to get executable path: %s\n", err)
				return
			}

			projectDir := filepath.Dir(exePath)

			dirPath := filepath.Join(projectDir, dirName)
			if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
				fmt.Printf("Failed to create directory: %s\n", err)
				return
			}

			fmt.Printf("Created directory: %s\n", dirPath)
		},
	}

	rootCmd.AddCommand(createDirCmd)

	var deleteFileCmd = &cobra.Command{
		Use:   "deletefile [file-name]",
		Short: "Delete a file in the current directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]

			exePath, err := os.Executable()
			if err != nil {
				fmt.Printf("Failed to get executable path: %s\n", err)
				return
			}

			projectDir := filepath.Dir(exePath)

			filePath := filepath.Join(projectDir, fileName)
			if err := os.Remove(filePath); err != nil {
				fmt.Printf("Failed to delete file: %s\n", err)
				return
			}

			fmt.Printf("Deleted file: %s\n", filePath)
		},
	}

	rootCmd.AddCommand(deleteFileCmd)

	var deleteDirCmd = &cobra.Command{
		Use:   "deletedir [dir-name]",
		Short: "Delete a directory in the current directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dirName := args[0]

			exePath, err := os.Executable()
			if err != nil {
				fmt.Printf("Failed to get executable path: %s\n", err)
				return
			}

			projectDir := filepath.Dir(exePath)

			dirPath := filepath.Join(projectDir, dirName)
			if err := os.RemoveAll(dirPath); err != nil {
				fmt.Printf("Failed to delete directory: %s\n", err)
				return
			}

			fmt.Printf("Deleted directory: %s\n", dirPath)
		},
	}
	rootCmd.AddCommand(deleteDirCmd)

	var generateCmd = &cobra.Command{
		Use:   "generate [template] [output] --name [name]",
		Short: "Generate code based on template",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			templateName := args[0]
			outputFileName := args[1]
			name, _ := cmd.Flags().GetString("name")

			exePath, err := os.Executable()
			if err != nil {
				fmt.Printf("Failed to get executable path: %s\n", err)
				return
			}
			executableDir := filepath.Dir(exePath)

			templatePath := filepath.Join(executableDir, "templates", templateName)
			outputPath := filepath.Join(executableDir, outputFileName)

			// Загрузить содержимое шаблона из файла
			templateContent, err := ioutil.ReadFile(templatePath)
			if err != nil {
				fmt.Printf("Failed to read template file: %s\n", err)
				return
			}

			// Создать шаблон и вставить значение в него
			tmpl, err := template.New("codeTemplate").Parse(string(templateContent))
			if err != nil {
				fmt.Printf("Failed to parse template: %s\n", err)
				return
			}

			data := struct {
				Name string
			}{
				Name: name,
			}

			// Создать выходной файл и записать в него результат
			outputFile, err := os.Create(outputPath)
			if err != nil {
				fmt.Printf("Failed to create output file: %s\n", err)
				return
			}
			defer outputFile.Close()

			if err := tmpl.Execute(outputFile, data); err != nil {
				fmt.Printf("Failed to generate code: %s\n", err)
				return
			}

			fmt.Printf("Generated code from template %s to %s\n", templatePath, outputPath)
		},
	}

	generateCmd.Flags().String("name", "", "Name to insert into the template")

	// Добавьте эту команду к вашему rootCmd
	rootCmd.AddCommand(generateCmd)

	var currentDirCmd = &cobra.Command{
		Use:   "currentdir",
		Short: "Show the current working directory",
		Run: func(cmd *cobra.Command, args []string) {
			currentDir, err := os.Getwd()
			if err != nil {
				fmt.Printf("Failed to get current directory: %s\n", err)
				return
			}
			fmt.Printf("Current working directory: %s\n", currentDir)
		},
	}

	rootCmd.AddCommand(currentDirCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
