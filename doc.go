//### Known issues:
//
//- A regexp in the parser should check for invalid expressions where two incompatible operators are following each other, such as Calculate("1 + * 4")
// If that's the case, the parser only applies the first operation found.
//
//- MathFunctions only handle integers
//
//- Substraction does not support substracting negative ints, e.g "1 - (-1)" should be 2
//
//- Pow does not support negative powers such as 1^-3, because of only using ints
//
//### Technical debt:
//
//- Extending pkg/grammar:
//	- TokenType type should include more types: whitespace, EOL, EOF, more operations
//	for extended capabilities
//
//	- Token type should have initial and ending position in expression
//	to better handle parsing the expression, doing it with indexes instead of strings/runes.
//	That would get rid of the problems of parsing whitespaces, multidigit numbers, etc
//
//- Extending pkg/parser:
//  - A much better approach would be creating a calculator as a Singleton structure, containing pointers to a lexer, a parser, a logger, etc
//	- Everything should be matched as a regexp: numbers, operations and invalid chars
//	- Lexer should be able to update the token initial and final position, to handle the next operator
//	- Include token associativity, which can be left or right depending on the operator
//	- Priorities should be addressed in its on file.go
//
//- Extending pkg/stack:
//	- The stack code was basically copy and paste, with minimal differences. This means that is the perfect use case for generics if the project supports Go 1.18+
//	- Some time ago I wrote a proof of concept library to handle generic slices (https://github.com/fbac/libslice), and the work would be the same
//	- The general idea would be treating Stacks with the same methods (Pop, Shift, Next, Shuffle, ...) indepently of its type. In the end it's just handling a fancy slice.
//
//- About project tree:
//	- I'm not specifically happy with the proposed tree, because there are pkg which are not supposed to be used externally. And for these pkg, such as grammar or eval, should be under internal/
//	- cmd, stack, parser can be extended in the future to be used by external programs as packages, so make sense for them to live under pkg/
//
//### Other design ideas (from best to worst IMHO)
//
//- Using ANTLR (https://www.antlr.org/). Probably the best idea, since it can create a powerful grammar with precedence, lexer and parsers. Also, I've always wanted to have the opportunity to learn more about it. Definitely going to do it asap.
//
//- Using a composite pattern: https://golangbyexample.com/composite-design-pattern-golang/ or a finite automata design.
//
//- Golang using external eval libraries already implemented, such as https://github.com/Knetic/govaluate. Probably the safest and best time saver approach.
//
//- python or other languages with built-in eval(). Not a big fan, due to being a really insecure method.
package main
