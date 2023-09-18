package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "filecli"}

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

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
