package scriptfilter

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestResponseWriteOmitsAbsentFields(t *testing.T) {
	resp := Response{Items: []Item{{Title: "Hello"}}}

	var buf bytes.Buffer
	if err := resp.Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	items, ok := decoded["items"].([]any)
	if !ok || len(items) != 1 {
		t.Fatalf("items = %v, want 1 item", decoded["items"])
	}
	item, ok := items[0].(map[string]any)
	if !ok {
		t.Fatalf("item = %v, want object", items[0])
	}
	if item["title"] != "Hello" {
		t.Errorf("title = %v, want %q", item["title"], "Hello")
	}
	for _, absent := range []string{"uid", "subtitle", "arg", "valid", "variables"} {
		if _, present := item[absent]; present {
			t.Errorf("field %q should be omitted when unset, got %v", absent, item[absent])
		}
	}
}

func TestResponseWriteIncludesExplicitFalse(t *testing.T) {
	resp := Response{Items: []Item{{Title: "Info", Valid: BoolPtr(false)}}}

	var buf bytes.Buffer
	if err := resp.Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}

	var decoded struct {
		Items []struct {
			Valid *bool `json:"valid"`
		} `json:"items"`
	}
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(decoded.Items) != 1 || decoded.Items[0].Valid == nil || *decoded.Items[0].Valid {
		t.Errorf("decoded valid = %v, want explicit false", decoded.Items[0].Valid)
	}
}

func TestResponseWriteIncludesVariables(t *testing.T) {
	resp := Response{Items: []Item{{Title: "X", Variables: map[string]string{"k": "v"}}}}

	var buf bytes.Buffer
	if err := resp.Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}

	var decoded struct {
		Items []struct {
			Variables map[string]string `json:"variables"`
		} `json:"items"`
	}
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got := decoded.Items[0].Variables["k"]; got != "v" {
		t.Errorf("variables[k] = %q, want %q", got, "v")
	}
}
