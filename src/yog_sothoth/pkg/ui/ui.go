package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	PrimaryColor   = lipgloss.Color("63") // Purple
	SecondaryColor = lipgloss.Color("39") // Blue
	SuccessColor   = lipgloss.Color("42") // Green
	WarningColor   = lipgloss.Color("208") // Orange
	ErrorColor     = lipgloss.Color("196") // Red
	MutedColor     = lipgloss.Color("240") // Gray

	// Base Styles
	TitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(PrimaryColor).
		PaddingLeft(1).
		MarginBottom(1)

	SuccessStyle = lipgloss.NewStyle().
		Foreground(SuccessColor).
		Bold(true)

	ErrorStyle = lipgloss.NewStyle().
		Foreground(ErrorColor).
		Bold(true)

	WarnStyle = lipgloss.NewStyle().
		Foreground(WarningColor).
		Bold(true)

	InfoStyle = lipgloss.NewStyle().
		Foreground(SecondaryColor)

	ItemStyle = lipgloss.NewStyle().
		PaddingLeft(2)

	BorderedBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(PrimaryColor).
		Padding(1, 2)
)

// Helpers to quickly style text
func RenderTitle(text string) string {
	return TitleStyle.Render(text)
}

func RenderSuccess(text string) string {
	return SuccessStyle.Render("✓ " + text)
}

func RenderError(text string) string {
	return ErrorStyle.Render("✗ " + text)
}

func RenderWarn(text string) string {
	return WarnStyle.Render("! " + text)
}

func RenderInfo(text string) string {
	return InfoStyle.Render("ℹ " + text)
}

func RenderItem(text string) string {
	return ItemStyle.Render("• " + text)
}

func RenderBox(text string) string {
	return BorderedBoxStyle.Render(text)
}
