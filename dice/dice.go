package dice

import (
	"errors"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/maja42/goval"
)

func Roll(rollText string) (string, error) {
	var Result int = 0
	var RegexRoll = regexp.MustCompile(`(?m)[+-]{0,1}\d{0,3}d\d{1,3}|[+-]\d{1,3}`)
	var foundSomething bool = false
	var ret string = "`"

	for _, match := range RegexRoll.FindAllString(rollText, -1) {
		foundSomething = true
		if match[0:1] != "+" && match[0:1] != "-" {
			match = "+" + match
		}

		if strings.Contains(match, "d") {
			//roll dices
			var sign string = match[0:1]
			totalResult, resultList, err := RollDice(match[1:])
			if err != nil {
				log.Println(err.Error())
				return "", err
			}

			ret = ret + "[ "
			for _, element := range resultList {
				ret = ret + strconv.Itoa(element) + " "
			}
			ret = ret + "] "

			if sign == "+" {
				Result = Result + totalResult
			} else {
				Result = Result - totalResult
			}

		} else {
			//add or subtract
			value, err := eval(strconv.Itoa(Result) + match)
			if err != nil {
				log.Println(err.Error())
				return "", err
			} else {
				Result = value
			}
		}
	}

	if !foundSomething {
		return "", errors.New("No match found in string " + rollText)
	}

	ret = ret[0:len(ret)-1] + "` for a total equal to: `" + strconv.Itoa(Result) + "`"

	return ret, nil
}

func RollDice(dice string) (int, []int, error) {
	var ret []int
	var total int = 0
	if dice[0:1] == "d" {
		dice = "1" + dice
	}
	splitted := strings.Split(dice, "d")
	numberOfDices, _ := strconv.Atoi(splitted[0])
	sidesOfDice, _ := strconv.Atoi(splitted[1])
	if numberOfDices > 500 {
		return 0, nil, errors.New("You can't roll more than 500 dices")
	}
	if sidesOfDice > 100 {
		return 0, nil, errors.New("You can't roll dices with more than 100 sides")
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < numberOfDices; i++ {
		ret = append(ret, r1.Intn(sidesOfDice)+1)
		total = total + ret[i]
	}

	return total, ret, nil
}

func eval(expression string) (int, error) {
	eval := goval.NewEvaluator()
	value, err := eval.Evaluate(expression, nil, nil)
	if err != nil {
		return 0, err
	} else {
		return value.(int), nil
	}
}
