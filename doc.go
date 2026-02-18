// Package cefr provides CEFR (Common European Framework of Reference) difficulty
// level assessment for English text.
//
// It uses a multi-feature fusion algorithm combining vocabulary analysis (50%),
// syntactic complexity (30%), and readability formulas (20%) to produce a CEFR level
// (A1–C2) and a continuous score (1.0–6.0) for any given English text.
//
// All word frequency data (Oxford 5000, NGSL, AWL, irregular forms, etc.) is embedded
// at compile time via embed.FS, so the library has zero external dependencies and works
// fully offline.
//
// Basic usage:
//
//	result, err := cefr.Assess("The cat sat on the mat.")
//	fmt.Println(result.Level, result.Score)
package cefr
