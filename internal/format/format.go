package format

import (
	"fmt"
	"time"
)

//nolint:mnd // All those numbers are fine.
func Time(t time.Time) string {
	ago := time.Until(t).Abs()

	if ago < 24*time.Hour {
		return "today"
	}

	if ago < 7*24*time.Hour {
		return "this week"
	}

	if ago < 30*24*time.Hour {
		return "this month"
	}

	if ago < 365*24*time.Hour {
		return "this year"
	}

	return fmt.Sprintf("%d years ago", int(ago.Hours()/24/365))
}

//nolint:mnd // All those numbers are fine.
func Count(i int) string {
	if i < 1000 {
		return fmt.Sprintf("%d", i)
	}

	if i < 1_000_000 {
		return fmt.Sprintf("%dK", i/1000)
	}

	if i < 1_000_000_000 {
		return fmt.Sprintf("%dM", i/1_000_000)
	}

	return "a lot"
}
