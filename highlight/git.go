package highlight

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)

func GetGitDiff(filePath string) (map[int]bool, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err // .git repository does not exist.
	}

	worktree, err := repo.Worktree() // Get a working tree that represents the current state of files in the project.
	if err != nil {
		return nil, err
	}

	status, err := worktree.Status() // Get the status of the repository tree. The status shows which files have been changed.
	if err != nil {
		return nil, err
	}

	fileStatus, ok := status[filePath] // Checking the status of a specific file.
	if !ok || fileStatus.Worktree == git.Unmodified {
		return nil, nil // No changes and no error.
	}

	currentContent, err := os.ReadFile(filePath) // Read the file contents, normalize the output lines.
	if err != nil {
		return nil, err
	}
	currentLines := strings.Split(normalizeLineEndings(string(currentContent)), "\n")

	if fileStatus.Worktree == git.Untracked {
		changedLines := make(map[int]bool)
		for i := range currentLines {
			changedLines[i+1] = true
		}
		return changedLines, nil
	}

	headRef, err := repo.Head() // Getting a reference to the current branch.
	if err != nil {
		return nil, err
	}

	headCommit, err := repo.CommitObject(headRef.Hash()) // Getting the latest commit.
	if err != nil {
		return nil, err
	}

	headFile, err := headCommit.File(filePath) // Get the state of a file in the latest commit.
	if err != nil {
		return map[int]bool{}, nil
	}

	headContent, err := headFile.Contents() // Get contents of file from latest commit.
	if err != nil {
		return nil, err
	}
	headLines := strings.Split(normalizeLineEndings(headContent), "\n")

	return calculateRealChanges(headLines, currentLines), nil
}

func normalizeLineEndings(s string) string { // We replace the line ending style. This is necessary for correct comparison.
	return strings.ReplaceAll(s, "\r\n", "\n")
}

func calculateRealChanges(before, after []string) map[int]bool {
	changes := make(map[int]bool)

	beforeIdx, afterIdx := 0, 0
	for beforeIdx < len(before) && afterIdx < len(after) {
		if before[beforeIdx] == after[afterIdx] {
			beforeIdx++
			afterIdx++
			continue
		}

		changes[afterIdx+1] = true

		lookAhead := 1
		for {
			if beforeIdx+lookAhead < len(before) && before[beforeIdx+lookAhead] == after[afterIdx] {
				beforeIdx += lookAhead
				break
			}
			if afterIdx+lookAhead < len(after) && before[beforeIdx] == after[afterIdx+lookAhead] {
				afterIdx += lookAhead
				break
			}
			if beforeIdx+lookAhead >= len(before) || afterIdx+lookAhead >= len(after) {
				beforeIdx++
				afterIdx++
				break
			}
			lookAhead++
		}
	}

	for i := afterIdx; i < len(after); i++ {
		changes[i+1] = true
	}

	return changes
}

func PrintWithGitHighlighting(content string, changedLines map[int]bool) {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		lineNum := i + 1
		if changedLines[lineNum] {
			fmt.Printf("\x1b[33m%s\x1b[0m\n", line)
		} else {
			fmt.Println(line)
		}
	}
}
