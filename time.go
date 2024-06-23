package checker

func IsBeforeNow(a any) bool {
	return toTime(a).Before(timeNow())
}

func IsBeforeToday(a any) bool {
	return toDate(a).Before(dateNow())
}

func IsBeforeDate(a, b any) bool {
	return toDate(a).Before(toDate(b))
}

func IsBefore(a, b any) bool {
	return toTime(a).Before(toTime(b))
}

func IsAfterNow(a any) bool {
	return toTime(a).After(timeNow())
}

func IsAfterToday(a any) bool {
	return toDate(a).After(dateNow())
}

func IsAfterDate(a, b any) bool {
	return toDate(a).After(toDate(b))
}

func IsAfter(a, b any) bool {
	return toTime(a).After(toTime(b))
}

func IsToday(a any) bool {
	return toDate(a).Equal(dateNow())
}
