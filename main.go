package main

import (
    "fmt"
    "os"
    "os/exec"
)

// ANSI escape codes for colors
const (
    reset = "\033[0m"
    blue  = "\033[1;34m"
    cyan  = "\033[1;36m"
    red   = "\033[1;31m"
    green = "\033[1;32m"
    bold  = "\033[1m"
)

func main() {
    args := os.Args[1:]

    if len(args) == 0 || isHelp(args[0]) {
        showHelp()
        return
    }

    switch args[0] {
    case "-v", "--version", "version":
        showVersion()
    case "alias":
        editFile("$PREFIX/etc/profile.d/alias.sh")
    case "defaults":
        editFile("$PREFIX/etc/profile.d/defaults.sh")
    case "functions":
        editFile("$PREFIX/etc/profile.d/functions.sh")
    case "paths":
        editFile("$PREFIX/etc/profile.d/paths.sh")
    case "profile":
        editFile("$PREFIX/etc/profile")
    default:
        fmt.Printf("%sUnknown command: %s%s\n", red, args[0], reset)
        showHelp()
    }
}

func isHelp(arg string) bool {
    helpOptions := []string{"-?", "-h", "--help", "help"}
    for _, opt := range helpOptions {
        if arg == opt {
            return true
        }
    }
    return false
}

func showHelp() {
    editor := os.Getenv("EDITOR")
    if editor == "" {
        editor = "nano"
    }

    prefix := os.Getenv("PREFIX")
    if prefix == "" {
        prefix = "/data/data/com.termux/files/usr"
    }

	fmt.Println("")
    fmt.Printf("%susage%s: setup <%soption%s>:\n", bold, reset, red, reset)
    fmt.Println("")
    fmt.Printf("       %salias%s     - Modify %s$PREFIX%s/etc/profile.d/alias.sh\n", red, reset, blue, reset)
    fmt.Printf("       %sdefaults%s  - Modify %s$PREFIX%s/etc/profile.d/defaults.sh\n", red, reset, blue, reset)
    fmt.Printf("       %sfunctions%s - Modify %s$PREFIX%s/etc/profile.d/functions.sh\n", red, reset, blue, reset)
    fmt.Printf("       %spaths%s     - Modify %s$PREFIX%s/etc/profile.d/paths.sh\n", red, reset, blue, reset)
    fmt.Printf("       %sprofile%s   - Modify %s$PREFIX%s/etc/profile\n", red, reset, blue, reset)
    fmt.Println("")
    fmt.Printf("       %s-?, -h, --help, help%s   - Shows help\n", red, reset)
    fmt.Printf("       %s-v, --version, version%s - Shows version\n", red, reset)
    fmt.Println("")
    fmt.Printf("%sCurrent PREFIX%s: %s%s%s\n", bold, reset, blue, prefix, reset)
    fmt.Printf("%sCurrent EDITOR%s: %s%s%s\n", bold, reset, cyan, editor, reset)
    fmt.Println("")
    fmt.Println("")
}

func showVersion() {
    fmt.Println("setup v1.0\nby PhateValleyman\nJonas.Ned@outlook.com")
}

func editFile(path string) {
    expandedPath := os.ExpandEnv(path)
    //fmt.Printf("Editing:\n%s%s%s\n", bold, expandedPath, reset)

    editor := os.Getenv("EDITOR")
    if editor == "" {
        editor = "nano"
    }

    cmd := exec.Command(editor, expandedPath)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Printf("%sError opening file: %v%s\n", red, err, reset)
        return
    }

    sourceFile(expandedPath)
}

func sourceFile(path string) {
    cmd := exec.Command("bash", "-c", fmt.Sprintf("source %s", path))
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Printf("%sError sourcing file: %v%s\n", red, err, reset)
    } else {
        fmt.Printf("%s%s\nedited and sourced%s\n", path, green, reset)
    }
}
