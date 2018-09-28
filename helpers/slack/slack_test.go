package slack

import (
	"strings"
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

	commands := []string{
		"<@ABCD> list",
		"<@ABCD> List",
		"<@ABCD>: List",
		"<@ABCD>List",
		"<@ABCD>:List",
		"<@ABCD> List <@ABCD>",
		"<@ABCD> List <@ABCD> ",
		"<@ABCD> List ",
		"<@ABCD> List 	",
	}

	for _, cmd := range commands {
		extractedCmd := ExtractCommand(userID, cmd)
		if extractedCmd != "list" {
			t.Error("expected", "list", "got", extractedCmd)
		}
	}
}

func TestExtractCommandFailure(t *testing.T) {
	userID := "ABCD"

	commands := []string{
		"<@HAXX> list",
		"<@HAXX> List",
		"<@HAXX>: List",
		"<@HAXX>List",
		"<@HAXX>:List",
		"<@HAXX> List <@HAXX>",
		"<@HAXX> List <@HAXX> ",
		"<@HAXX> List ",
		"<@HAXX> List 	",
	}

	for _, cmd := range commands {
		extractedCmd := ExtractCommand(userID, cmd)
		expectedCmd := strings.TrimSpace(strings.ToLower(cmd))
		if extractedCmd != expectedCmd {
			t.Error("expected", expectedCmd, "got", extractedCmd)
		}
	}
}

func TestExtractCommandFailurePartRemoved(t *testing.T) {
	userID := "ABCD"

	commands := []string{
		"<@HAXX> List <@ABCD>",
		"<@HAXX> List <@ABCD> ",
	}

	expectedCmd := "<@haxx> list"
	for _, cmd := range commands {
		extractedCmd := ExtractCommand(userID, cmd)
		if extractedCmd != expectedCmd {
			t.Error("expected", expectedCmd, "got", extractedCmd)
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
