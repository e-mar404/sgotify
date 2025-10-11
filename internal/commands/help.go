package command

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	constants "github.com/e-mar404/sgotify/internal/const"
)

func Help() error {
	var output string
	title := `                               $$\     $$\  $$$$$$\            
                               $$ |    \__|$$  __$$\           
 $$$$$$$\  $$$$$$\   $$$$$$\ $$$$$$\   $$\ $$ /  \__|$$\   $$\ 
$$  _____|$$  __$$\ $$  __$$\\_$$  _|  $$ |$$$$\     $$ |  $$ |
\$$$$$$\  $$ /  $$ |$$ /  $$ | $$ |    $$ |$$  _|    $$ |  $$ |
 \____$$\ $$ |  $$ |$$ |  $$ | $$ |$$\ $$ |$$ |      $$ |  $$ |
$$$$$$$  |\$$$$$$$ |\$$$$$$  | \$$$$  |$$ |$$ |      \$$$$$$$ |
\_______/  \____$$ | \______/   \____/ \__|\__|       \____$$ |
          $$\   $$ |                                 $$\   $$ |
          \$$$$$$  |                                 \$$$$$$  |
           \______/                                   \______/`

	
	title = constants.SecondaryTextCLI. 
		Border(lipgloss.ASCIIBorder()).
		BorderForeground(constants.Purple).
		Render(title)

	output += title + "\n\n"
	
	content := make([][]string, len(List))
	for i := range content {
		content[i] = make([]string, 2)
	}
	
	i := 0
	for _, cmd := range List {
		content[i] = []string{cmd.Name, cmd.Description}
		i++
	}

	constants.HelpTableCLI. 
		Headers("command", "description"). 
		Rows(content...)

	output += constants.HelpTableCLI.Render()
	
	width, _, _ := term.GetSize(os.Stdout.Fd()) 
	output = lipgloss.PlaceHorizontal(width, lipgloss.Center, output)

	fmt.Println(output)
	return nil
}

