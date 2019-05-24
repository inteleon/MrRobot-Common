package slack

import (
	"testing"
)

func TestIsCommandSuccess(t *testing.T) {
	userID := "ABCD"

	commands := []string{
		"<@ABCD> yoyo",
		"<@ABCD>: blabla yoyo <@ABCD>",
		"<@ABCD>yoyo",
		"<@ABCD>:blabla yoyo <@ABCD>",
	}

	for _, cmd := range commands {
		if !IsCommand(userID, cmd) {
			t.Error(cmd, "should be a command")
		}
	}
}

func TestIsCommandFailure(t *testing.T) {
	userID := "ABCD"

	commands := []string{
		"<ABCD> yoyo",
		"<ABCD>: yoyo",
		"blabla @ABCD yoyo",
		"blabla yoyo <@ABCD",
		"<@ABDC > blabla yoyo @ABCD>",
	}

	for _, cmd := range commands {
		if IsCommand(userID, cmd) {
			t.Error(cmd, "shouldn't be a command")
		}
	}
}

func TestExtractCommandSuccess(t *testing.T) {
	userID := "ABCD"

	inputs := []string{
		"<@ABCD> list",
		"<@ABCD> List",
		"<@ABCD> list hacker johnson",
		"<@ABCD> list hacker johnson ",
		"<@ABCD> list  hacker johnson ",
	}

	expectedCommand := "list"
	expectedArguments := []string{
		"",
		"",
		"hacker johnson",
		"hacker johnson",
		"hacker johnson",
	}

	for i, input := range inputs {
		extractedCmd, extractedArgs := ExtractCommand(userID, input)
		if extractedCmd != expectedCommand {
			t.Error("expected", "list", "got", extractedCmd)
		}
		if extractedArgs != expectedArguments[i] {
			t.Error("expected", expectedArguments[i], "got", extractedArgs)
		}
	}
}

func TestFormatCommand(t *testing.T) {
	commands := []string{
		"list",
		"LIST",
		" lIsT ",
		" lIst 		",
	}

	for _, cmd := range commands {
		formattedCmd := FormatCommand(cmd)
		if formattedCmd != "list" {
			t.Error("expected", "list", "got", formattedCmd)
		}
	}
}
