package solver

import (
	"testing"
)

var invalidRecords = [][]string{
	{"id", "first_name", "last_name", "email", "gender", "ip_address"},
	{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"},
	{"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"},
	{"3", "Gerri", "Choffin", "gchoffin2@ning.com", "", "9.254.198.50"},
	{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"},
	{"5", "Benoite", "Jaffray", "bjaffray4@github.com", "Female", ""},
	{"6", "Clint", "Oliphard", "coliphard5@ft.com", "Genderfluid", "39.69.123.72"},
	{"7", "Else", "Mc Dermid", "emcdermid6@plala.or.jp", "Female", "72.200.10.99"},
	{"8", "Andrea", "", "amckeran7@example.com", "", ""},
	{"", "Marnia", "De Roberto", "mderoberto8@mac.com", "Female", "99.28.96.140"},
	{"10", "Marnia", "De Roberto", "mderoberto8@mac.com"}} //avem mai putine records, trebuie eliminat

func TestDeleteInvalidLines(t *testing.T) {
	actualString := DeleteInvalidLines(invalidRecords)
	expectedResult := 6
	if len(actualString) != expectedResult {
		t.Errorf("The formatation have not been executed correctly")
	}
}

func TestNumberOfFilesNeeded(t *testing.T) {
	actualNumber := NumberOfFilesNeeded(invalidRecords, 2)
	expectedNumber := 6
	if actualNumber != expectedNumber {
		t.Errorf("The number of files to be created is invalid")
	}
}

func TestOutputValidation(t *testing.T) {
	var expectedOutput = [][]string{
		{"id", "first_name", "last_name", "email", "gender", "ip_address"},
		{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"},
		{"2", "Fan", "Gilvear", "fgilvear1@people.com.cn", "Female", "125.219.253.132"},
		{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"},
		{"6", "Clint", "Oliphard", "coliphard5@ft.com", "Genderfluid", "39.69.123.72"},
		{"7", "Else", "Mc Dermid", "emcdermid6@plala.or.jp", "Female", "72.200.10.99"}}

	actualResult := DeleteInvalidLines(invalidRecords)
	if len(expectedOutput) == len(actualResult) {
		for i := 0; i < len(expectedOutput); i++ {
			for j := 0; j < len(expectedOutput); j++ {
				if expectedOutput[i][j] != actualResult[i][j] {
					t.Errorf("The result is incorrect")
				}
			}
		}
	}
}
