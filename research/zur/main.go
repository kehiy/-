/*
* Based on: https://github.com/hazbo/the-super-tiny-compiler/blob/master/compiler.go
 */
package main

import (
	"fmt"
	"os"
	"strings"
)

type token struct {
	kind  string
	value string
}

type node struct {
	kind       string
	value      string
	name       string
	callee     *node
	expression *node
	body       []node
	params     []node
	arguments  *[]node
	context    *[]node
}

type ast node

var pc int
var pt []token

func isNumber(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= '0' && n <= '9' {
		return true
	}
	return false
}

func isLetter(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= 'a' && n <= 'z' {
		return true
	}
	return false
}

func tokenize(input string) []token {
	input += "\n"
	pos := 0
	tokens := make([]token, 0)

	for pos < len([]rune(input)) {
		char := string([]rune(input)[pos])

		switch {
		case char == "(", char == ")":
			tokens = append(tokens, token{
				kind:  "paren",
				value: char,
			})
			pos++
			continue

		case char == " ", char == "\n":
			pos++
			continue

		case char == "@":
			pos++
			char = string([]rune(input)[pos])

			for char != "@" {
				pos++
				char = string([]rune(input)[pos])
			}

			pos++

			continue

		case isNumber(char):
			value := ""

			for isNumber(char) {
				value += char
				pos++
				char = string([]rune(input)[pos])
			}

			tokens = append(tokens, token{
				kind:  "number",
				value: value,
			})

			continue

		case isLetter(char):
			value := ""

			for isLetter(char) {
				value += char
				pos++
				char = string([]rune(input)[pos])
			}

			tokens = append(tokens, token{
				kind:  "name",
				value: value,
			})

			continue

		default:
			panic(fmt.Sprintf("Invalid token: %s", char))
		}
	}

	return tokens
}

func parse(tokens []token) ast {
	pc = 0
	pt = tokens

	ast := ast{
		kind: "Program",
		body: []node{},
	}

	for pc < len(pt) {
		ast.body = append(ast.body, walk())
	}

	return ast
}

func walk() node {
	token := pt[pc]

	switch {
	case token.kind == "number":
		pc++
		return node{
			kind:  "NumberLiteral",
			value: token.value,
		}

	case token.kind == "paren" && token.value == "(":
		pc++
		token = pt[pc]

		n := node{
			kind:   "CallExpression",
			name:   token.value,
			params: []node{},
		}

		pc++
		token = pt[pc]

		for token.kind != "paren" || (token.kind == "paren" && token.value != ")") {
			n.params = append(n.params, walk())
			token = pt[pc]
		}

		pc++
		return n
	}

	panic(fmt.Sprintf("invalid tokens: %v", pt))
}

type visitor map[string]func(n *node, p node)

func traverse(a ast, v visitor) {
	traverseNode(node(a), node{}, v)
}

func traverseArray(a []node, p node, v visitor) {
	for _, child := range a {
		traverseNode(child, p, v)
	}
}

func traverseNode(n, p node, v visitor) {
	for k, va := range v {
		if k == n.kind {
			va(&n, p)
		}
	}

	switch n.kind {
	case "Program":
		traverseArray(n.body, n, v)
		break

	case "CallExpression":
		traverseArray(n.params, n, v)
		break

	case "NumberLiteral":
		break

	default:
		panic("can't traverse")
	}
}

func transform(a ast) ast {
	nast := ast{
		kind: "Program",
		body: []node{},
	}

	a.context = &nast.body

	traverse(a, map[string]func(n *node, p node){

		"NumberLiteral": func(n *node, p node) {
			*p.context = append(*p.context, node{
				kind:  "NumberLiteral",
				value: n.value,
			})
		},

		"CallExpression": func(n *node, p node) {

			e := node{
				kind: "CallExpression",
				callee: &node{
					kind: "Identifier",
					name: n.name,
				},
				arguments: new([]node),
			}

			n.context = e.arguments

			if p.kind != "CallExpression" {
				es := node{
					kind:       "ExpressionStatement",
					expression: &e,
				}
				*p.context = append(*p.context, es)
			} else {
				*p.context = append(*p.context, e)
			}

		},
	})

	return nast
}


func codeGenerate(n node) string {

	switch n.kind {
	case "Program":
		var r []string
		for _, no := range n.body {
			r = append(r, codeGenerate(no))
		}
		return strings.Join(r, "\n")
	case "ExpressionStatement":
		return codeGenerate(*n.expression) + ";"

	case "CallExpression":
		var ra []string
		c := codeGenerate(*n.callee)

		for _, no := range *n.arguments {
			ra = append(ra, codeGenerate(no))
		}

		r := strings.Join(ra, ", ")
		return c + "(" + r + ")"

	case "Identifier":
		return n.name

	case "NumberLiteral":
		return n.value

	default:
		panic("err")
	}
}

func compile(input string) string {
	tokens := tokenize(input)
	ast := parse(tokens)
	nast := transform(ast)
	out := codeGenerate(node(nast))

	return out
}

func main() {
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0o600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content := make([]byte, 1024)
	n, err := f.Read(content)
	if err != nil {
		panic(err)
	}

	f, err = os.Create("out.c")
    if err != nil {
	    panic(err)
    }
    defer f.Close()

	_, err = fmt.Fprintln(f, compile(string(content[:n])))
    if err != nil {
	    panic(err)
    }
}
