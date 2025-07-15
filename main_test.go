package main

import (
	"testing"
)

func _TestBump(t *testing.T, given, level, expected string) {
	res, _ := bump(level, given)
	if res != expected {
		t.Errorf("%s is not %s", res, expected)
	}
}

func TestPatchUpdate(t *testing.T) {
	_TestBump(t, "1.0.0", "Patch", "1.0.1")
	_TestBump(t, "1.1.0", "Patch", "1.1.1")
	_TestBump(t, "2.1.0", "Patch", "2.1.1")
}

func TestMinorUpdate(t *testing.T) {
	_TestBump(t, "1.0.0", "Minor", "1.1.0")
	_TestBump(t, "1.1.0", "Minor", "1.2.0")
	_TestBump(t, "2.1.0", "Minor", "2.2.0")
}

func TestMajorUpdate(t *testing.T) {
	_TestBump(t, "1.0.0", "Major", "2.0.0")
	_TestBump(t, "1.1.0", "Major", "2.0.0")
	_TestBump(t, "2.1.0", "Major", "3.0.0")
}

func TestInvalidPatch(t *testing.T) {
	_TestBump(t, "f1.0.0", "Patch", "f1.0.0")
	_TestBump(t, "1.f0.0", "Patch", "1.f0.0")
	_TestBump(t, "1.0.f0", "Patch", "1.0.f0")
}

func TestInvalidMinor(t *testing.T) {
	_TestBump(t, "f1.0.0", "Minor", "f1.0.0")
	_TestBump(t, "1.f0.0", "Minor", "1.f0.0")
	_TestBump(t, "1.0.f0", "Minor", "1.0.f0")
}

func TestInvalidMajor(t *testing.T) {
	_TestBump(t, "f1.0.0", "Major", "f1.0.0")
	_TestBump(t, "1.f0.0", "Major", "1.f0.0")
	_TestBump(t, "1.0.f0", "Major", "1.0.f0")
}

func TestBranchNamesPatch(t *testing.T) {
	_TestBump(t, "1.0.0-feature", "Patch", "1.0.0-feature")
	_TestBump(t, "1.0.0-feature/FOO-1", "Patch", "1.0.0-feature/FOO-1")
	_TestBump(t, "1.0.0-fix/foo", "Patch", "1.0.0-fix/foo")
}

func TestBranchNamesMinor(t *testing.T) {
	_TestBump(t, "1.0.0-feature", "Minor", "1.0.0-feature")
	_TestBump(t, "1.0.0-feature/FOO-1", "Minor", "1.0.0-feature/FOO-1")
	_TestBump(t, "1.0.0-fix/foo", "Minor", "1.0.0-fix/foo")
}

func TestBranchNamesMajor(t *testing.T) {
	_TestBump(t, "1.0.0-feature", "Major", "1.0.0-feature")
	_TestBump(t, "1.0.0-feature/FOO-1", "Major", "1.0.0-feature/FOO-1")
	_TestBump(t, "1.0.0-fix/foo", "Major", "1.0.0-fix/foo")
}
