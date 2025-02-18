package railfence

func Encode(message string, rails int) string {
	len := len(message)

	if len == 0 {
		return ""
	}

	// declare a 2D slice of strings
	// with the number of rows equal to the number of rails
	// and the number of columns equal to the length of the message
	// initialize all elements to empty strings
	rail := make([][]string, rails)

	for i := 0; i < rails; i++ {
		rail[i] = make([]string, len)
		for j := 0; j < len; j++ {
			rail[i][j] = "."
		}
	}

	// declare a variable to keep track of the current row
	row := 0

	// declare a variable to keep track of the direction of the rail
	// 1 means the rail is going down
	// -1 means the rail is going up
	direction := -1

	// iterate through the message
	for i := 0; i < len; i++ {
		// assign the current character to the current row
		rail[row][i] = string(message[i])

		// if the current row is the first row
		// or the current row is the last row
		// change the direction of the rail
		if row == 0 || row == rails-1 {
			direction = -direction
		}

		// move to the next row
		row += direction
	}

	result := ""
	for i := 0; i < rails; i++ {
		for j := 0; j < len; j++ {
			if rail[i][j] != "." {
				result += rail[i][j]
			}
		}
	}

	return result
}

func Decode(message string, rails int) string {
	len := len(message)

	if len == 0 {
		return ""
	}

	rail := make([][]string, rails)
	for i := 0; i < rails; i++ {
		rail[i] = make([]string, len)
	}

	row, col, direction := 0, 0, -1

	for i := 0; i < len; i++ {
		rail[row][col] = "?"

		if row == 0 || row == rails-1 {
			direction = -direction
		}

		row += direction
		col++
	}

	n := 0
	for i := 0; i < rails; i++ {
		for j := 0; j < len; j++ {
			if rail[i][j] == "?" {
				rail[i][j] = string(message[n])
				n++
			}
		}
	}

	result := ""
	row, direction = 0, -1
	for i := 0; i < len; i++ {
		result += rail[row][i]

		if row == 0 || row == rails-1 {
			direction = -direction
		}

		row += direction
	}

	return result
}
