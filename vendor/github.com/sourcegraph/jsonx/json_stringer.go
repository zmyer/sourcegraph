// Code generated by "stringer -type=ParseErrorCode,ScanErrorCode,SyntaxKind,NodeType -output=json_stringer.go"; DO NOT EDIT.

package jsonx

import "strconv"

const _ParseErrorCode_name = "InvalidSymbolInvalidNumberFormatPropertyNameExpectedValueExpectedColonExpectedCommaExpectedCloseBraceExpectedCloseBracketExpectedEndOfFileExpected"

var _ParseErrorCode_index = [...]uint8{0, 13, 32, 52, 65, 78, 91, 109, 129, 146}

func (i ParseErrorCode) String() string {
	if i < 0 || i >= ParseErrorCode(len(_ParseErrorCode_index)-1) {
		return "ParseErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ParseErrorCode_name[_ParseErrorCode_index[i]:_ParseErrorCode_index[i+1]]
}

const _ScanErrorCode_name = "NoneUnexpectedEndOfCommentUnexpectedEndOfStringUnexpectedEndOfNumberInvalidUnicodeInvalidEscapeCharacterInvalidCharacter"

var _ScanErrorCode_index = [...]uint8{0, 4, 26, 47, 68, 82, 104, 120}

func (i ScanErrorCode) String() string {
	if i < 0 || i >= ScanErrorCode(len(_ScanErrorCode_index)-1) {
		return "ScanErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ScanErrorCode_name[_ScanErrorCode_index[i]:_ScanErrorCode_index[i+1]]
}

const _SyntaxKind_name = "UnknownOpenBraceTokenCloseBraceTokenOpenBracketTokenCloseBracketTokenCommaTokenColonTokenNullKeywordTrueKeywordFalseKeywordStringLiteralNumericLiteralLineCommentTriviaBlockCommentTriviaLineBreakTriviaTriviaEOF"

var _SyntaxKind_index = [...]uint8{0, 7, 21, 36, 52, 69, 79, 89, 100, 111, 123, 136, 150, 167, 185, 200, 206, 209}

func (i SyntaxKind) String() string {
	if i < 0 || i >= SyntaxKind(len(_SyntaxKind_index)-1) {
		return "SyntaxKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SyntaxKind_name[_SyntaxKind_index[i]:_SyntaxKind_index[i+1]]
}

const _NodeType_name = "ObjectArrayPropertyStringNumberBooleanNull"

var _NodeType_index = [...]uint8{0, 6, 11, 19, 25, 31, 38, 42}

func (i NodeType) String() string {
	if i < 0 || i >= NodeType(len(_NodeType_index)-1) {
		return "NodeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NodeType_name[_NodeType_index[i]:_NodeType_index[i+1]]
}
