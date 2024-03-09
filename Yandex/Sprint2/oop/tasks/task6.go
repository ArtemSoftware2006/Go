package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	ReportID int
	Date     string
	User
}

func CreateReport(user User, reportDate string) Report {
	return Report{
		ReportID: 1,
		Date:     reportDate,
		User:     user,
	}
}

func PrintReport(report Report) {
	fmt.Printf("Report ID: %d\n Date: %s\n User ID: %d", report.ReportID, report.Date, report.ID)
}

func GenerateUserReports(users []User, reportDate string) []Report {
	var reports []Report
	for _, user := range users {
		reports = append(reports, CreateReport(user, reportDate))
	}
	return reports
}
