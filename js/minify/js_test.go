package minify

import (
	"os"
	"strings"
	"testing"
)

func TestMinify_ValidJS(t *testing.T) {
	input := `
		function add(a, b) {
			return a + b;
		}
	`
	want := `function add(e,t){return e+t}`
	got, err := Minify(input)
	if err != nil {
		t.Fatalf("Minify returned unexpected error: %v", err)
	}
	if strings.TrimSpace(got) != want {
		t.Errorf("Minify output = %q; want %q", got, want)
	}
}

func TestMinify_ValidJS_SimpleExpression(t *testing.T) {
	input := `var x = 1 + 2;`
	want := `var x=1+2`
	got, err := Minify(input)
	if err != nil {
		t.Fatalf("Minify returned unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("Minify output = %q; want %q", got, want)
	}
}

func TestMinify_ValidJS_WithComments(t *testing.T) {
	input := `
// This is a comment
function greet() {
    console.log("Hello, world!"); // inline comment
}
`
	want := `function greet(){console.log("Hello, world!")}`
	got, err := Minify(input)
	if err != nil {
		t.Fatalf("Minify returned unexpected error: %v", err)
	}
	if strings.TrimSpace(got) != want {
		t.Errorf("Minify output = %q; want %q", got, want)
	}
}

func TestMinify_ValidJS_RegexLiteral(t *testing.T) {
	input := `
var re = /ab+c/i;
`
	want := `var re=/ab+c/i`
	got, err := Minify(input)
	if err != nil {
		t.Fatalf("Minify returned unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("Minify output = %q; want %q", got, want)
	}
}

func TestMinify_ValidJS_TemplateLiteral(t *testing.T) {
	input := `
const msg = ` + "`Hello ${name}!`" + `;
`
	want := "const msg=`Hello ${name}!`"
	got, err := Minify(input)
	if err != nil {
		t.Fatalf("Minify returned unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("Minify output = %q; want %q", got, want)
	}
}

func TestMinify_ValidJS_MultipleStatements(t *testing.T) {
	input := `
let a = 5; let b = 10;
function sum() { return a + b; }
`
	want := `let a=5,b=10;function sum(){return a+b}`
	got, err := Minify(input)
	if err != nil {
		t.Fatalf("Minify returned unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("Minify output = %q; want %q", got, want)
	}
}

func TestMinify_EmptyString(t *testing.T) {
	input := ``
	want := ``
	got, err := Minify(input)
	if err != nil {
		t.Fatalf("Minify returned unexpected error for empty input: %v", err)
	}
	if got != want {
		t.Errorf("Minify output for empty string = %q; want %q", got, want)
	}
}

func TestMinify_SyntaxError(t *testing.T) {
	// Missing closing brace should trigger an error
	d, err := os.ReadFile("test.js")
	if err != nil {
		t.Fatalf("ReadFile returned unexpected error: %v", err)
	}
	r, err := Minify(string(d))
	if err != nil {
		t.Fatal("Minify return error for invalid JS")
	}

	c, err := os.ReadFile("test_minified.js")

	if r != string(c) {
		t.Fatal("Minify return not equal result")
	}
}

func TestMinify_big_code(t *testing.T) {
	// Missing closing brace should trigger an error
	input := `function bad(a, b { return a + b;`
	_, err := Minify(input)
	if err == nil {
		t.Fatal("Minify did not return error for invalid JS")
	}
}
