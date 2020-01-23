package goalseek

import(
	"math"
)

type Result struct{
	TargetValue float64
	AccuracyLevel float64
	Iterations int
	IsGoalReached bool
	ClosetValue float64
}

const delta float64 = 0.0001

func Run(funcInside func(x float64)float64, setters ...Option) *Result{
	
	//Default Options
	args :=&Options{
		TargetValue:0,
		AccuracyLevel:0.0000001,
		Iterations:25,
		ClosetValue:0,
	}

	for _,setter:=range setters{
		setter(args)
	}

	iterations:=0
	var guess float64 = 0

	result1:=(funcInside(guess)-args.TargetValue)

	for math.Abs(result1)>args.AccuracyLevel && iterations<args.Iterations{

		newGuess:=(guess+delta)
		result2:=(funcInside(newGuess)-args.TargetValue)

		if (result2-result1) !=0{
			guess=guess - result1 * (newGuess - guess) / (result2 - result1)
		}else{
			break
		}

		result1 = (funcInside(guess)-args.TargetValue)

		iterations++
	}

	if iterations>args.Iterations{
		iterations=args.Iterations
	}

	isGoalReache:=false
	if math.Abs(result1) <= args.AccuracyLevel{isGoalReache=true}

	return &Result{
		TargetValue:args.TargetValue,
		AccuracyLevel:args.AccuracyLevel,
		Iterations:iterations,
		IsGoalReached:isGoalReache,
		ClosetValue:guess,
	}
}

