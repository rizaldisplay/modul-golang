	for i := 0; i < data.Len(); i++ {
		data.Value = "value " + strconv.Itoa(i)
		data = data.Next()
	}