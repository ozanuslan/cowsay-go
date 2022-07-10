package cow

import (
	"math"
	"regexp"
)

const balloonTopChar = '_'
const balloonBottomChar = '-'
const balloonSingleLineLeftSideChar = '<'
const balloonSingleLineRightSideChar = '>'
const balloonLeftSideChar = '|'
const balloonRightSideChar = '|'
const balloonTopLeftChar = '/'
const balloonTopRightChar = '\\'
const balloonBottomLeftChar = '\\'
const balloonBottomRightChar = '/'

func Say(say string) string {
	say = reduceWhiteSpace(say)
	return speechBalloon(say) + cowBody()
}

func speechBalloon(say string) string {
	balloonTopBottomLen := len(say) + 2
	if balloonTopBottomLen > 40 + 2 {
		balloonTopBottomLen = 40 + 2
	}
	/*
	* Top and bottom of the speech balloon need extra 2
	* characters because of the 2 whitespaces on both
	* sides of the say string.
	*
	* Example:
	*  ______________
	* < Hello, World >
	*  --------------
	*/

	balloon := speechBalloonTop(balloonTopBottomLen)
	balloon += speechBalloonMiddle(say, 40)
	balloon += speechBalloonBottom(balloonTopBottomLen)
	return balloon
}

func speechBalloonTop(topLen int) string {
	balloonTop := " "
	for i := 0; i < topLen; i++ {
		balloonTop += string(balloonTopChar)
	}
	balloonTop += "\n"
	return balloonTop
}

func speechBalloonMiddle(say string, lineWidth int) string {
	balloonMiddle := ""
	lineCount := lineCount(len(say), lineWidth)
	if lineCount > 1 {
		for i := 0; i < lineCount; i++ {
			if i == 0 {
				balloonMiddle += string(balloonTopLeftChar)
				balloonMiddle += " "
				balloonMiddle += say[0:lineWidth]
				balloonMiddle += " "
				balloonMiddle += string(balloonTopRightChar)
			} else if i == lineCount - 1 {
				lastLineText := say[i*lineWidth:]
				spacesToAppend := lineWidth - len(lastLineText)

				balloonMiddle += string(balloonBottomLeftChar)
				balloonMiddle += " "
				balloonMiddle += lastLineText
				for i := 0; i < spacesToAppend; i++ {
					balloonMiddle += " "
				}
				balloonMiddle += " "
				balloonMiddle += string(balloonBottomRightChar)
			} else {
				balloonMiddle += string(balloonLeftSideChar)
				balloonMiddle += " "
				balloonMiddle += say[lineWidth*i:lineWidth*(i+1)]
				balloonMiddle += " "
				balloonMiddle += string(balloonRightSideChar)
			}
			balloonMiddle += "\n"
		}
	} else {
		balloonMiddle += string(balloonSingleLineLeftSideChar)
		balloonMiddle += " "
		balloonMiddle += say
		balloonMiddle += " "
		balloonMiddle += string(balloonSingleLineRightSideChar)
		balloonMiddle += " "
		balloonMiddle += "\n"
	}
	return balloonMiddle
}

func speechBalloonBottom(bottomLen int) string {
	balloonBottom := " "
	for i := 0; i < bottomLen; i++ {
		balloonBottom += string(balloonBottomChar)
	}
	balloonBottom += "\n"
	return balloonBottom
}

func reduceWhiteSpace(str string) string {
	r := regexp.MustCompile("\\s+")
	return r.ReplaceAllString(str, " ")
}

func lineCount(strLen int, lineWidth int) int {
	fLen := float64(strLen)
	fWidth := float64(lineWidth)
	return int(math.Ceil(fLen / fWidth))
}

func cowBody() string {
	cow :=
			"        \\   ^__^\n" +
			"         \\  (oo)\\_______\n" +
			"            (__)\\       )\\/\\\n" +
			"                ||----w |\n" +
			"                ||     ||"
	return cow
}