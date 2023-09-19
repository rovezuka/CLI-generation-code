package main

import (
	"fmt"
	"os"
	"path/filepath"

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

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
