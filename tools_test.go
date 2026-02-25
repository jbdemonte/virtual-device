package virtual_device

import (
	"errors"
	"testing"
)

func TestConcatErrors_AllNil(t *testing.T) {
	if err := concatErrors(nil, nil); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestConcatErrors_Empty(t *testing.T) {
	if err := concatErrors(); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestConcatErrors_SingleError(t *testing.T) {
	err := concatErrors(errors.New("boom"))
	if err == nil || err.Error() != "boom" {
		t.Errorf("expected 'boom', got %v", err)
	}
}

func TestConcatErrors_MultipleErrors(t *testing.T) {
	err := concatErrors(errors.New("a"), nil, errors.New("b"))
	if err == nil || err.Error() != "a; b" {
		t.Errorf("expected 'a; b', got %v", err)
	}
}

func TestConcatErrors_SingleNil(t *testing.T) {
	if err := concatErrors(nil); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
