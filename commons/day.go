package commons

type DayType struct {
	number      string
	readInputs  func(s string) []string
	SolveFirst  func(e bool) error
	SolveSecond func(e bool) error
}

var AllDays map[string]DayType = make(map[string]DayType)

func NewDay(d string, solveFirst, solveSecond func([]string) error) *DayType {
	var day DayType

	day.number = d

	day.SolveFirst = func(exampleMode bool) error {
		inputs, err := readInputs(exampleMode, day.number)

		if err != nil {
			return err
		}

		return solveFirst(inputs)
	}

	day.SolveSecond = func(exampleMode bool) error {
		inputs, err := readInputs(exampleMode, day.number)

		if err != nil {
			return err
		}
		
		return solveSecond(inputs)
	}

	AllDays[day.number] = day

	return &day
}

func readInputs(exampleMode bool, day string) ([]string, error) {
	if exampleMode {
		return ReadExample(day)
	} else {
		return ReadInput(day)
	}
}
