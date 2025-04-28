package logger

type FormatterTokenType int

const (
	TEXT FormatterTokenType = iota
	DATETIME
	LEVEL
	MESSAGE
)

var DefaultFormatString string = "[%datetime%] [%level%]: %message%"

type FormatterToken struct {
	ttype FormatterTokenType
	data  string
}

type Formatter struct {
	orderedTokens *[]FormatterToken
}

var DefaultFormatter *Formatter = NewFormatter(DefaultFormatString)

func _NewFormatter(formatString string) (*Formatter, error) {
	tokens, err := tokenizeFormatString(formatString)

	if err != nil {
		return nil, err
	}

	formatter := &Formatter{
		orderedTokens: tokens,
	}

	return formatter, nil
}

func NewFormatter(formatString string) *Formatter {
	formatter, err := _NewFormatter(formatString)
	if err != nil {
		printAndExit(err.Error(), 1)
	}
	return formatter
}
