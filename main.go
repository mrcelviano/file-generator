package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "file-generator"
	app.Usage = "File Generation Application"
	app.Commands = []cli.Command{
		{
			Name: "generate", ShortName: "g",
			Usage: "Generate files",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "fileName, fn",
					Usage: "Specifies the name of the files. The default name is 'test'",
					Value: "test",
				},
				cli.StringFlag{
					Name:  "filePath, fp",
					Usage: "Specifies the path to files. The default path is 'files'",
					Value: "files",
				},
				cli.Int64Flag{
					Name:  "fileSize, fs",
					Usage: "Specifies the size of files. Default size 10kb",
					Value: 1e4,
				},
				cli.Int64Flag{
					Name:  "numberFiles , nf",
					Usage: "Sets the number of files to generate. Default amount 10",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				var (
					fileName    = c.String("fileName")
					filePath    = c.String("filePath")
					fileSize    = c.Int("fileSize")
					numberFiles = c.Int("numberFiles")
				)

				err := os.Mkdir(filePath, os.ModeDir)
				if err != nil {
					return err
				}

				for i := 1; i < numberFiles+1; i++ {
					file, err := os.Create(fmt.Sprintf("%v/%v%v", filePath, fileName, i))
					if err != nil {
						return err
					}
					defer file.Close()
					data := make([]byte, fileSize, fileSize)
					file.Write(data)
				}

				return nil
			},
		},
		{
			Name: "delete", ShortName: "d",
			Usage: "Delete files",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "deleteDirectory, dir",
					Usage: "Delete all files in a directory. The default directory is 'files'",
					Value: "files",
				},
			},
			Action: func(c *cli.Context) error {
				var (
					fileDirectory = c.String("deleteDirectory")
				)

				err := os.RemoveAll(fileDirectory)
				if err != nil {
					return err
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
