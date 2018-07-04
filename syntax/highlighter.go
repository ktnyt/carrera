package syntax

import "github.com/ktnyt/pars"

type Highlighter struct {
	Black   pars.Parser
	Red     pars.Parser
	Green   pars.Parser
	Yellow  pars.Parser
	Blue    pars.Parser
	Magenta pars.Parser
	Cyan    pars.Parser
	White   pars.Parser

	BrightBlack   pars.Parser
	BrightRed     pars.Parser
	BrightGreen   pars.Parser
	BrightYellow  pars.Parser
	BrightBlue    pars.Parser
	BrightMagenta pars.Parser
	BrightCyan    pars.Parser
	BrightWhite   pars.Parser
}
