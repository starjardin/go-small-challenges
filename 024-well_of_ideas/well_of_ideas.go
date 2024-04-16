package kata

func Well(x []string) string {
	var good int

	for _, v := range x {
		if v == "good" {
			good += 1
		}
	}

	if good > 2 {
		return "I smell a series!"
	} else if good > 0 && good <= 2 {
		return "Publish!"
	} else {
		return "Fail!"
	}
}
