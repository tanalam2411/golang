package f


func Sprintf(format string, values ...interface{}) (r string) {
	offset := 0

	for i := 0; i < len(format); i++ {
		k := format[i]

		switch k {
		case '%':
			switch format[i+1] {
			case 's':
				r = r + values[offset].(string)
			case 'd':
				r = r + itoa(values[offset].(int))
			case '.':

			default:
				r = r + "(unhandled)"
			}

			i ++
			offset ++
		default:
			r = r + string(k)
		}
	}
	return
}