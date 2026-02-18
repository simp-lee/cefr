package cefr

import (
"fmt"
"testing"
)

func TestLevelString(t *testing.T) {
tests := []struct {
level Level
want  string
}{
{A1, "A1"},
{A2, "A2"},
{B1, "B1"},
{B2, "B2"},
{C1, "C1"},
{C2, "C2"},
{Level(0), "Unknown"},
{Level(7), "Unknown"},
{Level(-1), "Unknown"},
}
for _, tt := range tests {
t.Run(fmt.Sprintf("Level(%d)", int(tt.level)), func(t *testing.T) {
got := tt.level.String()
if got != tt.want {
t.Errorf("Level(%d).String() = %q; want %q", int(tt.level), got, tt.want)
}
})
}
}

func TestLevelConstants(t *testing.T) {
if int(A1) != 1 {
t.Errorf("A1 = %d; want 1", int(A1))
}
if int(A2) != 2 {
t.Errorf("A2 = %d; want 2", int(A2))
}
if int(B1) != 3 {
t.Errorf("B1 = %d; want 3", int(B1))
}
if int(B2) != 4 {
t.Errorf("B2 = %d; want 4", int(B2))
}
if int(C1) != 5 {
t.Errorf("C1 = %d; want 5", int(C1))
}
if int(C2) != 6 {
t.Errorf("C2 = %d; want 6", int(C2))
}
}

func TestScoreToLevel(t *testing.T) {
tests := []struct {
score float64
want  string
}{
// A1: [1.0, 1.5)
{1.0, "A1"},
{1.49, "A1"},
// A2: [1.5, 2.5)
{1.5, "A2"},
{2.49, "A2"},
// B1: [2.5, 3.5)
{2.5, "B1"},
{3.49, "B1"},
// B2: [3.5, 4.5)
{3.5, "B2"},
{4.49, "B2"},
// C1: [4.5, 5.5)
{4.5, "C1"},
{5.49, "C1"},
// C2: [5.5, 6.0]
{5.5, "C2"},
{6.0, "C2"},
// Edge: clamped values
{0.5, "A1"},
{7.0, "C2"},
}
for _, tt := range tests {
t.Run(fmt.Sprintf("score_%.2f", tt.score), func(t *testing.T) {
got := scoreToLevel(tt.score)
if got != tt.want {
t.Errorf("scoreToLevel(%.2f) = %q; want %q", tt.score, got, tt.want)
}
})
}
}

func TestLevelToBaseScore(t *testing.T) {
tests := []struct {
level int
want  float64
}{
{1, 1.0},
{2, 2.0},
{3, 3.0},
{4, 4.0},
{5, 5.0},
{6, 6.0},
{0, 1.0},
{7, 6.0},
{-1, 1.0},
}
for _, tt := range tests {
t.Run(fmt.Sprintf("level_%d", tt.level), func(t *testing.T) {
got := levelToBaseScore(tt.level)
if got != tt.want {
t.Errorf("levelToBaseScore(%d) = %.1f; want %.1f", tt.level, got, tt.want)
}
})
}
}

func TestClampScore(t *testing.T) {
tests := []struct {
input float64
want  float64
}{
{3.5, 3.5},
{1.0, 1.0},
{6.0, 6.0},
{0.0, 1.0},
{-1.0, 1.0},
{7.0, 6.0},
{100.0, 6.0},
}
for _, tt := range tests {
t.Run(fmt.Sprintf("clamp_%.1f", tt.input), func(t *testing.T) {
got := clampScore(tt.input)
if got != tt.want {
t.Errorf("clampScore(%.1f) = %.1f; want %.1f", tt.input, got, tt.want)
}
})
}
}
