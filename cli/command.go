package main

type Command struct {
	name string
	args []string
}

func pdfCommand(filePath string) Command {
	args := []string{"latexmk", "-pdflatex", "-bibtex", filePath}
	return Command{name: "run", args: args}
}