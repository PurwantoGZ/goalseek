package goalseek

type Options struct{
	TargetValue float64
	AccuracyLevel float64
	Iterations int
	ClosetValue float64
}

type Option func(*Options)

func TargetValue(targetValue float64)Option{
	return func(args *Options){
		args.TargetValue=targetValue
	}
}

func AccuracyLevel(accuracyLevel float64)Option{
	return func(args *Options){
		args.AccuracyLevel=accuracyLevel
	}
}

func Iterations(iterations int)Option{
	return func(args *Options){
		args.Iterations=iterations
	}
}

func ClosetValue(closetValue float64)Option{
	return func(args *Options){
		args.ClosetValue=closetValue
	}
}