package courses

func Remove(courses []string, courseToBeRemoved string) []string {
	uniqueCourses := []string{}

	for _, value := range courses {
		if value != courseToBeRemoved {
			uniqueCourses = append(uniqueCourses, value)
		}
	}

	return uniqueCourses
}
