package attendances

/* func mapAttendances(dbResult []DbAttendanceResult) []Attendance {
	res := make([]Attendance, len(dbResult))
	for i, v := range dbResult {
		res[i] = *v.Into()
	}

	return res
}

func allHook(_ *dbx.Query, sliceA any, op func(sliceB any) error) error {
	switch v := sliceA.(type) {
	case *[]Attendance:
		dbAttendancesRes := new([]DbAttendanceResult)
		if err := op(dbAttendancesRes); err != nil {
			return err
		}

		attendances := mapAttendances(*dbAttendancesRes)
		*v = attendances
		return nil
	default:
		return op(sliceA)
	}
} */
