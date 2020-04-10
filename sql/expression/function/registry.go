package function

import (
	"math"

	"github.com/src-d/go-mysql-server/sql"
	"github.com/src-d/go-mysql-server/sql/expression/function/aggregation"
)

// Defaults is the function map with all the default functions.
var Defaults = []sql.Function{
	sql.Function0{Name: "connection_id", Fn: NewConnectionID},
	sql.Function0{Name: "user", Fn: NewUser},
	sql.Function0{Name: "current_user", Fn: NewUser},
	sql.Function0{Name: "now", Fn: NewNow},

	sql.Function1{Name: "abs", Fn: NewAbsVal},
	sql.Function1{Name: "array_length", Fn: NewArrayLength},
	sql.Function1{Name: "avg", Fn: func(e sql.Expression) sql.Expression { return aggregation.NewAvg(e) }},
	sql.Function1{Name: "ceil", Fn: NewCeil},
	sql.Function1{Name: "ceiling", Fn: NewCeil},
	sql.Function1{Name: "char_length", Fn: NewCharLength},
	sql.Function1{Name: "character_length", Fn: NewCharLength},
	sql.Function1{Name: "count", Fn: func(e sql.Expression) sql.Expression { return aggregation.NewCount(e) }},
	sql.Function1{Name: "date", Fn: NewDate},
	sql.Function1{Name: "day", Fn: NewDay},
	sql.Function1{Name: "dayofmonth", Fn: NewDay},
	sql.Function1{Name: "dayofweek", Fn: NewDayOfWeek},
	sql.Function1{Name: "dayofyear", Fn: NewDayOfYear},
	sql.Function1{Name: "explode", Fn: NewExplode},
	sql.Function1{Name: "first", Fn: func(e sql.Expression) sql.Expression { return aggregation.NewFirst(e) }},
	sql.Function1{Name: "floor", Fn: NewFloor},
	sql.Function1{Name: "from_base64", Fn: NewFromBase64},
	sql.Function1{Name: "hour", Fn: NewHour},
	sql.Function1{Name: "is_binary", Fn: NewIsBinary},
	sql.Function1{Name: "json_unquote", Fn: NewJSONUnquote},
	sql.Function1{Name: "last", Fn: func(e sql.Expression) sql.Expression { return aggregation.NewLast(e) }},
	sql.Function1{Name: "length", Fn: NewLength},
	sql.Function1{Name: "ln", Fn: NewLogBaseFunc(float64(math.E))},
	sql.Function1{Name: "log10", Fn: NewLogBaseFunc(float64(10))},
	sql.Function1{Name: "log2", Fn: NewLogBaseFunc(float64(2))},
	sql.Function1{Name: "lower", Fn: NewLower},
	sql.Function1{Name: "ltrim", Fn: NewTrimFunc(lTrimType)},
	sql.Function1{Name: "max", Fn: func(e sql.Expression) sql.Expression { return aggregation.NewMax(e) }},
	sql.Function1{Name: "min", Fn: func(e sql.Expression) sql.Expression { return aggregation.NewMin(e) }},
	sql.Function1{Name: "minute", Fn: NewMinute},
	sql.Function1{Name: "month", Fn: NewMonth},
	sql.Function1{Name: "reverse", Fn: NewReverse},
	sql.Function1{Name: "rtrim", Fn: NewTrimFunc(rTrimType)},
	sql.Function1{Name: "second", Fn: NewSecond},
	sql.Function1{Name: "sleep", Fn: NewSleep},
	sql.Function1{Name: "soundex", Fn: NewSoundex},
	sql.Function1{Name: "sqrt", Fn: NewSqrt},
	sql.Function1{Name: "sum", Fn: func(e sql.Expression) sql.Expression { return aggregation.NewSum(e) }},
	sql.Function1{Name: "to_base64", Fn: NewToBase64},
	sql.Function1{Name: "trim", Fn: NewTrimFunc(bTrimType)},
	sql.Function1{Name: "upper", Fn: NewUpper},
	sql.Function1{Name: "weekday", Fn: NewWeekday},
	sql.Function1{Name: "year", Fn: NewYear},

	sql.Function3{Name: "if", Fn: NewIf},
	sql.Function2{Name: "ifnull", Fn: NewIfNull},
	sql.Function2{Name: "nullif", Fn: NewNullIf},
	sql.Function2{Name: "pow", Fn: NewPower},
	sql.Function2{Name: "power", Fn: NewPower},
	sql.Function2{Name: "repeat", Fn: NewRepeat},
	sql.Function2{Name: "split", Fn: NewSplit},

	sql.Function3{Name: "replace", Fn: NewReplace},
	sql.Function3{Name: "substring_index", Fn: NewSubstringIndex},

	sql.FunctionN{Name: "coalesce", Fn: NewCoalesce},
	sql.FunctionN{Name: "concat", Fn: NewConcat},
	sql.FunctionN{Name: "concat_ws", Fn: NewConcatWithSeparator},
	sql.FunctionN{Name: "date_add", Fn: NewDateAdd},
	sql.FunctionN{Name: "date_sub", Fn: NewDateSub},
	sql.FunctionN{Name: "greatest", Fn: NewGreatest},
	sql.FunctionN{Name: "json_extract", Fn: NewJSONExtract},
	sql.Function2{Name: "instr", Fn: NewInstr},
	sql.FunctionN{Name: "least", Fn: NewLeast},
	sql.Function2{Name: "left", Fn: NewLeft},
	sql.FunctionN{Name: "log", Fn: NewLog},
	sql.FunctionN{Name: "lpad", Fn: NewPadFunc(lPadType)},
	sql.FunctionN{Name: "mid", Fn: NewSubstring},
	sql.FunctionN{Name: "regexp_matches", Fn: NewRegexpMatches},
	sql.FunctionN{Name: "round", Fn: NewRound},
	sql.FunctionN{Name: "rpad", Fn: NewPadFunc(rPadType)},
	sql.FunctionN{Name: "substr", Fn: NewSubstring},
	sql.FunctionN{Name: "substring", Fn: NewSubstring},
	sql.FunctionN{Name: "unix_timestamp", Fn: NewUnixTimestamp},
	sql.FunctionN{Name: "timestamp", Fn: NewTimestamp},
	sql.FunctionN{Name: "datetime", Fn: NewDatetime},
	sql.FunctionN{Name: "yearweek", Fn: NewYearWeek},
}
